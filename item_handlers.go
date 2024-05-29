package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreateItemData struct {
	Name   string `json:"name"`
	Price  uint   `json:"price"`
	UserID uint   `json:"user_id"`
}
type UpdateItemData struct {
	Completed bool `json:"completed"`
}

func (b *AppBackend) CreateItemHandler(w http.ResponseWriter, r *http.Request, group *Group) {

	itemData := CreateItemData{}
	err := json.NewDecoder(r.Body).Decode(&itemData)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}

	item := ListItem{
		Name:    itemData.Name,
		Price:   itemData.Price,
		UserID:  itemData.UserID,
		GroupID: group.ID,
	}

	err = b.CreateItem(&item)

	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}
	HandleSuccessJson(w, 200, item)
}

func (b *AppBackend) UpdateItemHandler(w http.ResponseWriter, r *http.Request, group *Group) {

	itemIdStr := chi.URLParam(r, "itemID")
	itemId, err := strconv.Atoi(itemIdStr)

	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not retrieve group %v", err))
		return
	}

	itemData := UpdateItemData{}
	err = json.NewDecoder(r.Body).Decode(&itemData)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}

	item := ListItem{}

	err = b.DB.Find(&item, itemId).Error

	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}

	item.Completed = itemData.Completed

	err = b.DB.Save(&item).Error

	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}
	HandleSuccessJson(w, 200, item)
}

func (b *AppBackend) FetchItemsByGroupHandler(w http.ResponseWriter, r *http.Request, group *Group) {
	items := []ListItem{}
	err := b.GetItemsByGroup(group.ID, &items)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}
	HandleSuccessJson(w, 200, items)
}
