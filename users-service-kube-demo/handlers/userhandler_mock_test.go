package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"users-service/handlers"
	"users-service/internal/mocks"
	"users-service/messaging"
	"users-service/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newAppWithMocks(t *testing.T) (*fiber.App, *mocks.MockIUserDB, *messaging.Messaging, func()) {
	t.Helper()
	ctrl := gomock.NewController(t)
	mockDB := mocks.NewMockIUserDB(ctrl)
	msg := &messaging.Messaging{ChMessaging: make(chan []byte, 1)} // buffer to avoid blocking

	h := handlers.NewUserHandler(mockDB) // mock object over here

	app := fiber.New()
	// mount routes
	app.Post("/users", h.CreateUser(msg))

	// app.Get("/users/:id", h.GetUserBy)
	// app.Get("/users/:limit/:offset", h.GetUsersByLimit)
	// app.Post("/orders", h.CreateOrder)

	cleanup := func() {
		close(msg.ChMessaging)
		ctrl.Finish()
		_ = app.Shutdown()
	}
	return app, mockDB, msg, cleanup
}

func doJSON(app *fiber.App, method, path string, body any) (*http.Response, []byte, error) {
	var req *http.Request
	if body != nil {
		b, _ := json.Marshal(body)
		req = httptest.NewRequest(method, path, bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req = httptest.NewRequest(method, path, nil)
	}
	res, err := app.Test(req, int(time.Second)*5)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	return res, b, nil
}

func TestCreateUser_OK(t *testing.T) {
	app, mockDB, msg, cleanup := newAppWithMocks(t)
	defer cleanup()

	in := map[string]any{
		"name":   "Jiten",
		"email":  "JitenP@outlook.com",
		"mobile": "9618558500",
	}

	// Expect Create to be called with *models.User parsed from body. We don't care about exact pointer identity.
	mockDB.EXPECT().
		Create(gomock.AssignableToTypeOf(&models.User{})).
		DoAndReturn(func(u *models.User) (*models.User, error) {
			// verify handler-set fields
			assert.Equal(t, "active", u.Status)
			assert.NotZero(t, u.LastModified)
			// return "saved" copy (ID assigned)
			return &models.User{
				CommonModel: models.CommonModel{ID: 1, Status: u.Status, LastModified: u.LastModified},
				Name:        u.Name,
				Email:       u.Email,
				Mobile:      u.Mobile,
			}, nil
		}).Times(1)

	res, body, err := doJSON(app, http.MethodPost, "/users", in)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	var got models.User
	require.NoError(t, json.Unmarshal(body, &got))
	assert.Equal(t, uint(1), got.ID)
	assert.Equal(t, "Jiten", got.Name)
	assert.Equal(t, "active", got.Status)
	assert.NotZero(t, got.LastModified)

	// channel got the message
	select {
	case b := <-msg.ChMessaging:
		var streamed models.User
		_ = json.Unmarshal(b, &streamed) // if ToBytes isn't JSON, just check len(b) > 0
		assert.Equal(t, uint(1), streamed.ID)
	default:
		t.Fatalf("expected a message on ChMessaging")
	}
}
