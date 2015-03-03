package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/ruprict/loccasions-go/app"
	"log"
	"net/http"
	"strconv"
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

type LocLink struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

func (user User) MarshalJSON() ([]byte, error) {
	var locLinks []LocLink
	for _, element := range user.Loccasions {
		l := LocLink{ID: element.ID, Link: "/loccasions/" + strconv.Itoa(element.ID)}
		locLinks = append(locLinks, l)
	}
	return json.Marshal(map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"locassions": locLinks,
	})
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

	var users []User

	context.Db.Preload("Loccasions").Find(&users)

	js, _ := json.Marshal(users)
	rw.Write(js)
	return nil
}

func UsersUpdateHandler(context *app.Context, rw http.ResponseWriter, req *http.Request) error {
	rw.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("ERROR: %v", err)
		return err
	}
	vars := mux.Vars(req)
	userId := vars["userId"]
	if userId == "" {
		log.Println("No UserID provided")
		return errors.New("No UserID provided")
	}

	user.ID, err = strconv.Atoi(userId)
	if err != nil {
		log.Println("ERROR: %v", err)
		return err
	}
	context.Db.Save(&user)
	rw.WriteHeader(http.StatusAccepted)
	js, err := json.Marshal(user)

	if err != nil {
		log.Println("ERROR: %v", err)
		return err
	}
	rw.Write(js)
	return nil
}
