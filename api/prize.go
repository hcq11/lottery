package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"lottery/configs"
	"lottery/format"
	"lottery/format/Code"
	"lottery/model"
	"lottery/utils/common"
	"net/http"
	"os"
	"strconv"
)

// @Tags 奖品设置
// @Summary 获取所有奖品
// @Description
// @Accept json
// @Produce json
// @Success 200 {object} model.Prize
// @Failure 400 {object} format.Response
// @Router /prizes [get]
func GetPrizes(ctx *gin.Context) {
	prizes, err := model.GetPrizes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, format.Response{Code: Code.Failed, Msg: err.Error()})
	}
	ctx.JSON(http.StatusOK, prizes)

}

// @Tags 奖品设置
// @Summary 增加奖品
// @Description
// @Accept multipart/form-data
// @Produce json
// @Param Name formData string true "奖品名称"
// @Param Count formData integer true "数量" 取值 1 ~ 9999" minimum(1) maximum(9999)
// @Param Unit formData string true "单位 个,件..."
// @Param Rate formData integer true "中奖率" 取值 1 ~ 9999" minimum(1) maximum(9999)
// @Param Remarks formData string false "备注"
// @Param file formData file true "奖品图片"
// @Success 200 {object} format.Response
// @Failure 400 {object} format.Response
// @Router /prize [post]
func AddPrize(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, format.Response{Code: Code.Failed, Msg: err.Error()})
		return
	}
	if b, _ := common.IsFileExist(configs.ImgPath); !b {
		os.MkdirAll(configs.ImgPath, os.ModePerm)
	}
	filePath := configs.ImgPath + "/" + header.Filename
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to open the file for writing")
		ctx.JSON(http.StatusBadRequest, format.Response{Code: Code.Failed, Msg: err.Error()})
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println(err)
	}
	Name := ctx.PostForm("Name")
	Count := ctx.PostForm("Count")
	iCount, _ := strconv.Atoi(Count)
	Unit := ctx.PostForm("Unit")
	Rate := ctx.PostForm("Rate")
	iRate, _ := strconv.Atoi(Rate)
	Remarks := ctx.PostForm("Remarks")
	prize := model.Prize{
		Name:    Name,
		Pic:     configs.ImgPrefix + "/" + header.Filename,
		Count:   iCount,
		Unit:    Unit,
		Rate:    iRate,
		Remarks: Remarks,
	}
	model.AddPrize(prize)
	ctx.JSON(http.StatusOK, format.Response{Code: Code.Success, Msg: "success"})
}

// @Tags 奖品设置
// @Summary 编辑奖品
// @Description 编辑奖品
// @Accept multipart/form-data
// @Produce json
// @Param ID formData string true "奖品ID"
// @Param Name formData string false "奖品名称"
// @Param Count formData integer false "数量" 取值 1 ~ 9999" minimum(1) maximum(9999)
// @Param Unit formData string false "单位 个,件..."
// @Param Rate formData integer false "中奖率" 取值 1 ~ 9999" minimum(1) maximum(9999)
// @Param Remarks formData string false "备注"
// @Success 200 {object} format.Response
// @Failure 400 {object} format.Response
// @Router /prize [put]
func EditPrize(ctx *gin.Context) {
	id := ctx.PostForm("ID")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, format.Response{Code: Code.Failed, Msg: "error:not exist prize ID! "})
		return
	}
	iID, _ := strconv.Atoi(id)
	Name := ctx.PostForm("Name")
	Count := ctx.PostForm("Count")
	iCount, _ := strconv.Atoi(Count)
	Unit := ctx.PostForm("Unit")
	Rate := ctx.PostForm("Rate")
	iRate, _ := strconv.Atoi(Rate)
	Remarks := ctx.PostForm("Remarks")
	prize := model.Prize{
		ID:      uint(iID),
		Name:    Name,
		Count:   iCount,
		Unit:    Unit,
		Rate:    iRate,
		Remarks: Remarks,
	}
	model.EditPrize(prize)
	ctx.JSON(http.StatusOK, format.Response{Code: Code.Success, Msg: "success"})
}

// @Tags 奖品设置
// @Summary 删除奖品
// @Description
// @Accept json
// @Produce json
// @Param id path string true "奖品ID"
// @Success 200 {object} format.Response
// @Failure 400 {object} format.Response
// @Router /prize/{id} [delete]
func DelPrize(ctx *gin.Context) {
	ID := ctx.Param("id")
	iID, err := strconv.Atoi(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, format.Response{Code: Code.Failed, Msg: err.Error()})
		return
	}
	prize := model.Prize{
		ID: uint(iID),
	}
	model.DelPrize(prize)
	ctx.JSON(http.StatusOK, format.Response{Code: Code.Success, Msg: "success"})
}


