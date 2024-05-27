package main

import "gorm.io/gorm"

type ListItem struct {
	gorm.Model
	Name    string
	Price   uint
	GroupID uint
	Group   Group `gorm:"constraint:OnDelete:CASCADE;"`
}

type Group struct {
	gorm.Model
	SecretCode string
	Name       string
}

func (b *AppBackend) CreateGroup(groupData *Group) error {
	err := b.DB.Create(groupData).Error
	return err
}

func (b *AppBackend) GetAllGroups(groups *[]Group) error {
	err := b.DB.Find(groups).Error
	return err
}

func (b *AppBackend) GetItemsByGroup(GroupID uint, items *[]ListItem) error {
	err := b.DB.Find(items, ListItem{GroupID: GroupID}).Error
	return err
}

func (b *AppBackend) CreateItem(item *ListItem) error {
	err := b.DB.Create(item).Error
	return err
}
