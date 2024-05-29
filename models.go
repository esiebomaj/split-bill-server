package main

import "gorm.io/gorm"

type ListItem struct {
	gorm.Model
	Name      string
	Completed bool
	Price     uint
	GroupID   uint
	Group     Group `gorm:"constraint:OnDelete:CASCADE;"`

	UserID uint
	User   *User `gorm:"constraint:OnDelete:CASCADE;"`
}

type Group struct {
	gorm.Model
	SecretCode string
	Name       string
	Users      []*User `gorm:"many2many:user_groups;"`
}

type User struct {
	gorm.Model
	Name   string
	Groups []*Group `gorm:"many2many:user_groups;"`
}

func (b *AppBackend) CreateGroup(groupData *Group) error {
	err := b.DB.Create(groupData).Error
	return err
}

func (b *AppBackend) GetAllGroups(groups *[]Group) error {
	err := b.DB.Find(groups).Error
	return err
}

func (b *AppBackend) GetAllGroupById(GroupID uint, group *Group) error {
	err := b.DB.First(group, GroupID).Error
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

func (b *AppBackend) CreateUser(userData *User) error {
	err := b.DB.Create(userData).Error
	return err
}

func (b *AppBackend) DBGetAllUsers(users *[]User) error {
	err := b.DB.Find(users).Error
	return err
}

func (b *AppBackend) DBAddUserToGroup(group *Group, userID uint) error {
	user := User{}
	b.DB.First(&user, userID)
	err := b.DB.Model(group).Association("Users").Append(&user)
	return err
}

func (b *AppBackend) GetGroupWithUsers(group *Group) error {
	err := b.DB.Preload("Users").First(&group).Error
	return err
}
