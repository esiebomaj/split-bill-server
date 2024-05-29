package main

import (
	"encoding/json"
	"net/http"
)

type UserToGroupData struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

func (b *AppBackend) GetAllUsers(w http.ResponseWriter, r *http.Request, group *Group) {

	users := []User{}
	err := b.DBGetAllUsers(&users)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}
	HandleSuccessJson(w, 200, users)
}

func (b *AppBackend) AddUserToGroup(w http.ResponseWriter, r *http.Request, group *Group) {

	body := UserToGroupData{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}
	if body.Name != "" {
		user := User{
			Name: body.Name,
		}
		err = b.CreateUser(&user)
		if err != nil {
			HandleErrorJson(w, 400, err.Error())
			return
		}
		body.UserID = user.ID
	}
	if body.UserID == 0 {
		HandleErrorJson(w, 400, "User ID required")
		return
	}

	err = b.DBAddUserToGroup(group, body.UserID)

	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}

	HandleSuccessJson(w, 200, group)

}
