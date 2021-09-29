package api

import (
	"github.com/gin-gonic/gin"
	"lottery/model"
	"lottery/utils"
	"net/http"
)
type User struct {
	name string
}
type AWard struct {
	UserName string
	prizeID  uint
	Name     string
	Pic      string
	Count    int
	Unit     string //单位
	Remarks  string
}
func Lottery(ctx *gin.Context)  {
	//user := &User{}
	name := ctx.PostForm("name")
	//if err := ctx.Bind(user); err != nil {
	//	ctx.JSON(http.StatusBadRequest,err.Error())
	//	return
	//}
	prizes,err := model.GetPrizes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest,"Lottery failed")
		return
	}
	prize, err:= utils.Lottery(prizes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	award := AWard{}
	award.UserName = name
	award.prizeID = prize.ID
	award.Name = prize.Name
	award.Pic = prize.Pic
	award.Count = 1
	award.Unit = prize.Unit
	award.Remarks = prize.Remarks
	DispatchData(award)
	ctx.JSON(http.StatusOK,award)
}