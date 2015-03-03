package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/ruprict/loccasions-go/app"
	"log"
	"net/http"
	"strconv"
)

type Loccasion struct {
	ID          int    `json: "id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int    `json:"userId"`
}

func (loc Loccasion) MarshalJSON() ([]byte, error) {
	var userlink bytes.Buffer
	userlink.WriteString("/users/")
	userlink.WriteString(strconv.Itoa(loc.UserID))
	return json.Marshal(map[string]interface{}{
		"id":          loc.ID,
		"name":        loc.Name,
		"description": loc.Description,
		"user":        userlink.String(),
	})
}

func LoccasionsIndexHandler(context *app.Context, rw http.ResponseWriter, req *http.Request) error {
	rw.Header().Set("Content-Type", "application/json")
	// Get Loccsions
	vars := mux.Vars(req)
	userId := vars["userId"]
	if userId == "" {
		log.Println("No UserID provided")
		return errors.New("No UserID provided")
	}

	var user User
	var locs []Loccasion
	context.Db.Find(&user, userId).Related(&locs)

	js, _ := json.Marshal(&locs)
	rw.Write(js)
	return nil
}

func LoccasionsCreateHandler(context *app.Context, rw http.ResponseWriter, req *http.Request) error {
	rw.Header().Set("Content-Type", "application/json")
	// Get loccasion params
	decoder := json.NewDecoder(req.Body)
	var loc Loccasion
	err := decoder.Decode(&loc)
	if err != nil {
		log.Println("ERROR: %v", err)
		return err
	}
	vars := mux.Vars(req)
	userId := vars["userId"]

	if userId == "" {
		log.Println("ERROR: UserID not supplied")
		return errors.New("UserID not supplied")
	}
	loc.UserID, err = strconv.Atoi(userId)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	context.Db.Create(&loc)

	rw.WriteHeader(http.StatusCreated)

	js, _ := json.Marshal(map[string]interface{}{
		"created": true,
	})
	rw.Write(js)
	return nil
}
