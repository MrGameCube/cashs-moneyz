package main

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	Sender   Person
	Receiver Person
	Value    int64
}

type Person struct {
	gorm.Model
	Name   string
	Avatar []byte `gorm:"type:blob"`
	Email  string
}

type Group struct {
	gorm.Model
	Name   string
	Avatar []byte `gorm:"type:blob" json:"-"`
	//Members      []Person
	//Transactions []Transaction
	Currency string
}
type Verisimilitude struct {
	gorm.Model
}
