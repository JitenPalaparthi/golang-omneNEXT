package handlers

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"fiber-gorm-users/internal/models"
)

type UsersHandler struct {
	db  *gorm.DB
	log zerolog.Logger
}

func NewUsersHandler(db *gorm.DB, log zerolog.Logger) *UsersHandler {
	return &UsersHandler{db: db, log: log}
}

type createUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   string `json:"status"`
}

type updateUserReq struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Status   *string `json:"status"`
}

func (h *UsersHandler) Create(c *fiber.Ctx) error {
	var req createUserReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name, email, password are required")
	}
	if len(req.Password) < 6 {
		return fiber.NewError(fiber.StatusBadRequest, "password must be at least 6 chars")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to hash password")
	}

	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Status: func() string {
			if req.Status == "" {
				return "active"
			}
			return req.Status
		}(),
	}
	user.Touch()

	if err := h.db.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(strings.ToLower(err.Error()), "unique") {
			return fiber.NewError(fiber.StatusConflict, "email already exists")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "db error")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UsersHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	skip, _ := strconv.Atoi(c.Query("skip", "0"))
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if skip < 0 {
		skip = 0
	}

	var users []models.User
	var total int64
	h.db.Model(&models.User{}).Count(&total) // select count(*) from users

	if err := h.db.
		Order("id ASC").
		Limit(limit).
		Offset(skip).
		Find(&users).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "db error")
	}

	return c.JSON(fiber.Map{
		"data":  users,
		"limit": limit,
		"skip":  skip,
		"total": total,
	})
}

func (h *UsersHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "db error")
	}
	return c.JSON(user)
}

func (h *UsersHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "db error")
	}

	var req updateUserReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = strings.TrimSpace(strings.ToLower(*req.Email))
	}
	if req.Status != nil {
		user.Status = *req.Status
	}
	if req.Password != nil && *req.Password != "" {
		if len(*req.Password) < 6 {
			return fiber.NewError(fiber.StatusBadRequest, "password must be at least 6 chars")
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to hash password")
		}
		user.PasswordHash = string(hash)
	}
	user.LastModified = time.Now().Unix()

	if err := h.db.Save(&user).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "unique") {
			return fiber.NewError(fiber.StatusConflict, "email already exists")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "db error")
	}
	return c.JSON(user)
}

func (h *UsersHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.db.Delete(&models.User{}, id).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "db error")
	}
	return c.SendStatus(fiber.StatusNoContent)
}
