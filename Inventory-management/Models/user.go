package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Role      string    `json:"role" binding:"required,oneof=admin member"`
	CreatedAt time.Time `json:"created_at"`
}

type Userlogin struct {
	ID       uint   `json:"id"`
	Password string `json:"password"`
}
