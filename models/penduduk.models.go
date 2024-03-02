package models

type Penduduks struct {
	Id   int    `json:"id gorm:"primary_key"`
	Nama string `json:"nama"`
	Rt   string `json:"rt"`
}
