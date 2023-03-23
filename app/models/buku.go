package models

import "time"

type Buku struct {
	Id        int       `gorm:"primary_key;" json:"id"`
	Title     string    `gorm:"size:255;null;" json:"name_book"`
	Author    string    `gorm:"size:255;null;" json:"author"`
	Desc      string    `gorm:"size:255;null;" json:"desc"`
	CreatedAt time.Time `gorm:"autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
