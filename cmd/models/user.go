package models

/*
Contains the models for the User Table
*/
type User struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *User) TableName() string {
	return "user"
}
