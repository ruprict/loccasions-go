package api

import (
	"encoding/json"
	"github.com/ruprict/loccasions-go/app"
	"net/http"
	"time"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `sql:"not null"json:"name"`
	Email      string `sql:"not null;unique"json:"email"`
	Loccasions []Loccasion
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func UsersCreateHandler(context *app.Context, rw http.ResponseWriter, req *http.Request) error {
	rw.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		return err
	}

	context.Db.Create(&user)
	rw.WriteHeader(201)

	js, _ := json.Marshal(map[string]interface{}{
		"created": true,
	})
	rw.Write(js)
	return nil
}

func UsersIndexHandler(context *app.Context, rw http.ResponseWriter, req *http.Request) error {
	rw.Header().Set("Content-Type", "application/json")
	// Get Loccsions

	var users []User

	context.Db.Find(&users)

	js, _ := json.Marshal(users)
	rw.Write(js)
	return nil
}
