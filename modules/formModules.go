package modules

import (
	"github.com/go-playground/validator/v10"
)

type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=8"`
}

type RegisterForm struct {
	Username  string `form:"username" json:"username" binding:"required,min=1,max=150,usernamerule"`
	Password  string `form:"password" json:"password" binding:"required,min=1,max=120"`
	Firstname string `form:"first_name" json:"first_name" binding:"max=150"`
	Lastname  string `form:"last_name" json:"last_name" binding:"max=150"`
	Email     string `form:"email" json:"email" binding:"required,email,max=254"`
}

var UsernameRule validator.Func = func(fl validator.FieldLevel) bool {
	s, _ := fl.Field().Interface().(string)
	for _, ch := range s {
		if !((ch >= 48 && ch <= 57) || (ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122) ||
			ch == 64 || ch == 43 || ch == 45 || ch == 46 || ch == 95) {
			return false
		}
	}
	return true
}
