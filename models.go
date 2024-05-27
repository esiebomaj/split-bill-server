package main

import "gorm.io/gorm"

type ListItem struct {
	gorm.Model
	GroupID uint
	Name    string
	Price   uint
}

type Group struct {
	gorm.Model
	SecretCode string
	Name       string
	ListItems  []ListItem
}

func (b *AppBackend) CreateGroup(groupData *Group) error {
	err := b.DB.Create(groupData).Error
	return err
}

func (b *AppBackend) GetAllGroups(groups *[]Group) error {
	err := b.DB.Find(groups).Error
	return err
}
