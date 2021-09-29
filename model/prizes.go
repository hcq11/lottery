package model

//import (
//	"github.com/jinzhu/gorm"
//)

type Prize struct {
	ID      uint `gorm:"primary_key"`
	Name    string
	Pic		string
	Count   int
	Unit    string   //单位
	Rate    int
	Remarks string `gorm:"size:1000"`
}

func GetPrizes() ([]Prize,error) {
	var prizes []Prize
	res := DB.Find(&prizes)
	return prizes,res.Error
}

func AddPrize(prize Prize)  {
  	 DB.Create(&prize)
}

func DelPrize(prize Prize) {
	DB.Delete(&prize)
}