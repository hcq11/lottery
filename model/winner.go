package model

type Winner struct {
	ID   uint `gorm:"primary_key"`
	UserName string
	PrizeID  uint
	Name     string
	Pic      string
	Count    int
	Unit     string //单位
	Remarks  string
	Avatar   string `gorm:"size:1000"`
}

func GetWinners() ([]Winner, error) {
	var winners []Winner
	res := DB.Find(&winners)
	return winners, res.Error
}
func AddWinner(winner Winner) {
	DB.Create(&winner)
}
