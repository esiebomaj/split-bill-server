package main

import (
	"encoding/json"
	"net/http"
)

type CreateGroupData struct {
	Name       string `json:"name"`
	SecretCode string `json:"secret_code"`
}

func (b *AppBackend) FetchGroupsHandler(w http.ResponseWriter, r *http.Request) {
	groups := []Group{}
	err := b.GetAllGroups(&groups)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}
	HandleSuccessJson(w, 200, groups)
}

func (b *AppBackend) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {

	groupData := CreateGroupData{}
	err := json.NewDecoder(r.Body).Decode(&groupData)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}

	group := Group{
		SecretCode: groupData.SecretCode,
		Name:       groupData.Name,
	}

	err = b.CreateGroup(&group)

	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}
	HandleSuccessJson(w, 200, group)
}

func (b *AppBackend) GetGroupDetailsHandler(w http.ResponseWriter, r *http.Request, group *Group) {
	items := []ListItem{}
	err := b.GetItemsByGroup(group.ID, &items)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}

	err = b.GetGroupWithUsers(group)
	if err != nil {
		HandleErrorJson(w, 400, err.Error())
		return
	}

	// type ItemWithoutGroup struct {
	// 	ID    uint   `json:"id"`
	// 	Name  string `json:"name"`
	// 	Price int    `json:"price"`

	// 	CreatedAt string `json:"created_at"`
	// 	UpdatedAt string `json:"updated_at"`

	// 	Completed bool `json:"completed"`
	// 	GroupID   uint `json:"group_id"`
	// }

	// var itemsWithoutGroup []ItemWithoutGroup

	// for _, item := range items {
	// 	itemsWithoutGroup = append(itemsWithoutGroup, ItemWithoutGroup{
	// 		ID:        item.ID,
	// 		Name:      item.Name,
	// 		Price:     int(item.Price),
	// 		CreatedAt: item.CreatedAt.String(),
	// 		UpdatedAt: item.UpdatedAt.String(),
	// 		Completed: item.Completed,
	// 		GroupID:   item.GroupID,
	// 	})
	// }

	type response struct {
		Items []ListItem
		Group Group
	}

	HandleSuccessJson(w, 200, response{items, *group})
}
