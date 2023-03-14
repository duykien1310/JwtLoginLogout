package models

import "gorm.io/gorm"

type VerifyCode struct {
	gorm.Model
	Code int `json:"Code"`
}
