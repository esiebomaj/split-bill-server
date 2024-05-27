package main

import (
	"encoding/json"
	"net/http"
)

func (b *AppBackend) HealthHandler(w http.ResponseWriter, r *http.Request) {
	HandleSuccessJson(w, 200, "Ok!")
}

func (b *AppBackend) FetchGroupsHandler(w http.ResponseWriter, r *http.Request) {
	groups := []Group{}
	err := b.GetAllGroups(&groups)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
	}
	HandleSuccessJson(w, 200, groups)
}

type CreateGroupData struct {
	Name       string `json:"name"`
	SecretCode string `json:"secret_code"`
}

func (b *AppBackend) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {

	groupData := CreateGroupData{}
	err := json.NewDecoder(r.Body).Decode(&groupData)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
	}

	group := Group{
		SecretCode: groupData.SecretCode,
		Name:       groupData.Name,
	}

	err = b.CreateGroup(&group)

	if err != nil {
		HandleErrorJson(w, 400, err.Error())
	}
	HandleSuccessJson(w, 200, group)
}
