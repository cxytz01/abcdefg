package v1

import (
	"net/http"
	"producer/internal/busi/core"
	"producer/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CreateCampaign godoc
// @Description create campaign
// @Tags Producer-External-API-V1
// @Accept multipart/form-data
// @Produce application/json,json
// @Param CreateCampaignParam query core.CreateCampaignParam true "CreateCampaignParam"
// @Param file formData file true "CSV file"
// @Success 200 {object} nil
// @Failure 400 {object} utils.ResponseWithRequestId
// @Failure 500 {object} utils.ResponseWithRequestId
// @Router /api/v1/campaign [post]
func CreateCampaign(c *gin.Context) {
	app := utils.Gin{C: c}

	var r core.CreateCampaignParam
	if err := c.ShouldBindQuery(&r); err != nil {
		app.HTTPResponse(http.StatusBadRequest, utils.NewResponse(utils.CodeBadRequest, err.Error(), nil))
		return
	}

	formFile, err := c.FormFile("file")
	if err != nil {
		app.HTTPResponse(http.StatusBadRequest, utils.NewResponse(utils.CodeBadRequest, err.Error(), nil))
		return
	}

	var f core.FileUploaded
	csvstore, _ := c.Get(CSVStore)
	f.StoreDir = csvstore.(string)

	if err := f.CheckUploadFile(formFile); err != nil {
		app.HTTPResponse(http.StatusBadRequest, utils.NewResponse(utils.CodeBadRequest, err.Error(), nil))
		return
	}

	result, resp := core.CreateCampaign(c.Request.Context(), &r, &f)
	if resp != nil {
		app.HTTPResponse(resp.HttpCode, resp.Response)
		return
	}

	app.HTTPResponseOK(result)
}

// DispatchToKafka godoc
// @Description dispatch suitable messages to kafka
// @Tags Producer-Internal-API-V1
// @Produce application/json,json
// @Success 204 {object} nil
// @Failure 400 {object} utils.ResponseWithRequestId
// @Failure 500 {object} utils.ResponseWithRequestId
// @Router /api/v1/messages [post]
func DispatchToKafka(c *gin.Context) {
	app := utils.Gin{C: c}

	if resp := core.DispatchToKafka(c.Request.Context()); resp != nil {
		app.HTTPResponse(resp.HttpCode, resp.Response)
		return
	}

	app.HTTPResponse204()
}
