package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"lottery/configs"
	"lottery/model"
	"lottery/utils/common"
	"net/http"
	"os"
	"strconv"
)

// @Tags 抽奖
// @Summary 抽奖
// @Description
// @Accept json
// @Produce json
// @Success 200 {object} []model.Prize
// @Failure 400 string failed
// @Router /prizes [get]
func GetPrizes(ctx *gin.Context)  {
	prizes,err := model.GetPrizes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest,"failed")
	}
	ctx.JSON(http.StatusOK,prizes)

}

// @Tags 抽奖
// @Summary 抽奖
// @Description
// @Accept json
// @Produce json
// @Param Name formData string true "奖品名称"
// @Param Count formData integer true "数量" 取值 1 ~ 9999" minimum(1) maximum(9999)
// @Param Unit formData string true "单位 个,件..."
// @Param Rate formData integer true "中奖率" 取值 1 ~ 9999" minimum(1) maximum(9999)
// @Param Remarks formData string false "备注"
// @Param file formData file true "奖品图片"
// @Success 200 string Success
// @Failure 400 string message
// @Router /prize [post]
func AddPrize(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest,  err.Error())
		return
	}
	if b, _ := common.IsFileExist(configs.ImgPath); !b {
		os.MkdirAll(configs.ImgPath, os.ModePerm)
	}
	filePath := configs.ImgPath + "/" + header.Filename
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to open the file for writing")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println(err)
	}
	Name := ctx.PostForm("Name")
	Count := ctx.PostForm("Count")
	iCount,_ := strconv.Atoi(Count)
	Unit := ctx.PostForm("Unit")
	Rate := ctx.PostForm("Rate")
	iRate,_ := strconv.Atoi(Rate)
	Remarks := ctx.PostForm("Remarks")
	prize := model.Prize{
		Name:    Name,
		Pic:     configs.ImgPrefix+"/"+header.Filename,
		Count:   iCount,
		Unit:    Unit,
		Rate:    iRate,
		Remarks: Remarks,
	}
	model.AddPrize(prize)
	ctx.JSON(http.StatusOK, "Success")
}

// @Tags 抽奖
// @Summary 抽奖
// @Description
// @Accept json
// @Produce json
// @Param id path string true "奖品ID"
// @Success 200 string Success
// @Failure 400 string message
// @Router /prize/{id} [delete]
func DelPrize(ctx *gin.Context) {
	ID := ctx.Param("id")
	iID, err := strconv.Atoi(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid prize id")
		return
	}
	prize := model.Prize{
		ID: uint(iID),
	}
	model.DelPrize(prize)
	ctx.JSON(http.StatusOK, "Success")
}

