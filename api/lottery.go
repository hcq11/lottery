package api

import (
	"github.com/gin-gonic/gin"
	"lottery/format"
	"lottery/format/Code"
	"lottery/model"
	"lottery/utils"
	"net/http"
)

// @Tags 抽奖
// @Summary 用户抽奖
// @Description
// @Accept json
// @Produce json
// @Param luckyDog body format.User true "抽奖者信息"
// @Success 200 {object} model.Winner
// @Failure 400 {object} format.Response
// @Router /lottery [post]
func Lottery(ctx *gin.Context)  {
	//name := ctx.PostForm("name")
	//avatar := ctx.PostForm("avatar")
	var user format.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest,format.Response{Code:Code.Failed,Msg:err.Error()})
		return
	}
	prizes,err := model.GetPrizes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest,format.Response{Code:Code.Failed,Msg:"Lottery failed"})
		return
	}
	prize, err:= utils.Lottery(prizes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,format.Response{Code:Code.Failed,Msg:err.Error()})
		return
	}
	award := model.Winner{}
	award.UserName = user.Name
	award.PrizeID = prize.ID
	award.Name = prize.Name
	award.Pic = prize.Pic
	award.Count = 1
	award.Unit = prize.Unit
	award.Remarks = prize.Remarks
	award.Avatar = user.Avatar
	DispatchData(award)
	model.AddWinner(award)
	ctx.JSON(http.StatusOK,award)
}

// @Tags 抽奖
// @Summary 获取中奖记录
// @Description
// @Accept json
// @Produce json
// @Success 200 {object} model.Winner
// @Failure 400 {object} format.Response
// @Router /winners [get]
func GetWinners(ctx *gin.Context) {
	winners, err := model.GetWinners()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, format.Response{Code: Code.Failed, Msg: err.Error()})
	}
	ctx.JSON(http.StatusOK, winners)
}