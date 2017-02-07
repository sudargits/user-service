package main

import "github.com/jinzhu/gorm"

//Model DB User
type User struct {
	gorm.Model
	Apps string //Apps Code
	Username string
	Email string
	Firstname string
	Lastname string
	Additionls []Additional
	RoleID uint
}

type Role struct {
	gorm.Model
	Role string
	User []User
}
type Additional struct {
	gorm.Model
	UserID uint
	Key string
	Field string
}


func (role* Role)createRole(dba *gorm.DB) *Role{
	defer dba.Close()
	dba.NewRecord(role)
	dba.Create(role)
	return &role
}

func (role* Role) updateRole(dba *gorm.DB) *Role{
	defer dba.Close()
	dba.Save(role)
	return &role
}
func (role* Role) deleteRole(dba *gorm.DB) uint {
	defer dba.Close()
	dba.Delete(role)
}

