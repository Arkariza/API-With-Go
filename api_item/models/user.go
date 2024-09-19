package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	UserName string `gorm:"type:varchar(500)" json:"UserName"`
	Password string `gorm:"type:varchar(500)" json:"Password"`
	Role     int64  `gorm:"integer" json:"Role"`
}

func (u *User) IsAdmin() bool {
	return u.Id == 1
}

func (u *User) IsWorker() bool {
	return u.Id == 2
}
