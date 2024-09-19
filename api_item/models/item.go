package models

type Item struct{
	Id int64 `gorm:"pimaryKey" json:"id"`
	NameItem string `gorm:"type:varchar(500)" json:"nameItem"`
	Price int64 `gorm:"integer" json:"price"`
	Stock int64 `gorm:"integer" json:"stock"`
}