package model
import (
	"github.com/jinzhu/gorm"
)
type User struct {
	gorm.Model
	UserName string
	PhoneNum string
	Nickname string
	Remarks  string `gorm:"size:1000"`
	Avatar   string `gorm:"size:1000"`
}
