package models

type Petugases struct {
	Id       int    `json:"id" gorm:"primary_key`
	Nama     int    `json:"nama" `
	Password string `json:"password" gorm:"not null;"`
	Role     string `json:"role"`
}
