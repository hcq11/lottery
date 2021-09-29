package utils

import (
	"errors"
	"lottery/model"
	"math/rand"
)
//抽奖算法
func Lottery(arr []model.Prize) (model.Prize,error){
	var total = 0
	for _,item := range arr {
		if item.Count <= 0 {
			continue
		}
		total += item.Rate
	}

	for _,item := range arr {
		if item.Count <= 0 {
			continue
		}
		var random = rand.Intn(total)
		if random < item.Rate {
			return item,nil
		} else {
			total -= item.Rate
		}
	}
	return model.Prize{},errors.New("抽奖出错")
}