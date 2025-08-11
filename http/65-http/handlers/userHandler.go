package handlers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"io"
	"math/rand/v2"
	"net/http"
	"time"
)

type UserHandler struct {
	FileName string
}

func NewUserHandler(fileName string) *UserHandler {
	return &UserHandler{fileName}
}

func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user := new(models.User)

	err = json.Unmarshal(bytes, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = user.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user.Id = uint(rand.IntN(1000))
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	bytes, err = json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	_, err = utils.SaveToFile(uh.FileName, []byte(string(bytes)+"\n"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("user successfully saved to the file"))
	w.WriteHeader(http.StatusCreated)

	// Write the logic to save to file

}
