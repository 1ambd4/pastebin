package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Pastebin struct {
	ID      int64  `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Context string `json:"context" form:"context"`
}

func (p *Pastebin) TableName() string {
	return "pastebin"
}

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open("pastebin.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Pastebin{})
}

func GetContext(id int64) (string, error) {
	var pastebin Pastebin
	if err := db.Model(&Pastebin{}).Where("id=?", id).Find(&pastebin).Error; err != nil {
		return err.Error(), err
	}
	return pastebin.Context, nil
}

func InsertContext(context string) (id int64, err error) {
	var pastebin = &Pastebin{Context: context}
	if err = db.Create(&pastebin).Error; err != nil {
		return -1, err
	} else {
		return pastebin.ID, nil
	}
}

func UpdateContext(id int64, context string) (err error) {
	var pastebin = &Pastebin{ID: id}

	if err = db.Model(&pastebin).Update("context", context).Error; err != nil {
		return err
	}

	return nil
}

func DeleteContext(id int64) (context string, err error) {
	var pastebin = &Pastebin{ID: id}
	if err = db.Delete(&pastebin).Error; err != nil {
		return err.Error(), err
	}

	return pastebin.Context, nil
}
