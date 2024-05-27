package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreateGroupData struct {
	Name       string `json:"name"`
	SecretCode string `json:"secret_code"`
}

type CreateItemData struct {
	Name    string `json:"name"`
	GroupID uint   `json:"group_id"`
	Price   uint   `json:"price"`
}

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

func (b *AppBackend) CreateItemHandler(w http.ResponseWriter, r *http.Request) {

	itemData := CreateItemData{}
	err := json.NewDecoder(r.Body).Decode(&itemData)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
	}

	item := ListItem{
		Name:    itemData.Name,
		Price:   itemData.Price,
		GroupID: itemData.GroupID,
	}

	err = b.CreateItem(&item)

	if err != nil {
		HandleErrorJson(w, 400, err.Error())
	}
	HandleSuccessJson(w, 200, item)
}

func (b *AppBackend) FetchItemsByGroupHandler(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
	}
	items := []ListItem{}
	err = b.GetItemsByGroup(uint(groupID), &items)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
	}
	HandleSuccessJson(w, 200, items)
}
