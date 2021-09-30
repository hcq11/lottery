package api

import (
	"github.com/gin-gonic/gin"
	"lottery/format"
	"lottery/format/Code"
	"lottery/model"
	"lottery/utils"
	"net/http"
)

func Lottery(ctx *gin.Context)  {
	name := ctx.PostForm("name")
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
	award.UserName = name
	award.PrizeID = prize.ID
	award.Name = prize.Name
	award.Pic = prize.Pic
	award.Count = 1
	award.Unit = prize.Unit
	award.Remarks = prize.Remarks
	DispatchData(award)
	model.AddWinner(award)
	ctx.JSON(http.StatusOK,award)
}