package ginrouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/services/models"
)

// InfoHandler  represent the httphandler for Info
type InfoHandler struct {
	CUsecase models.InfoUsecase
}

// NewInfoHandler will initialize the service/api/v1/ endpoint
func NewInfoHandler(base *gin.RouterGroup, us models.InfoUsecase) {
	handler := &InfoHandler{
		CUsecase: us,
	}

	base.GET("/info", handler.GetInfo) //e.g.: http://127.0.0.1:8888/service/api/v1/info
	base.PATCH("/info", handler.UpdateInfo)

}

// @Summary Get info API
// @Schemes http xx
// @Description Get info API
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /info [get]
func (a *InfoHandler) GetInfo(c *gin.Context) {
	//call usecase method
	info, _ := a.CUsecase.GetInfo()

	c.JSON(http.StatusOK, info)
}

// @Summary test patch InfoUpdate API
// @Description just test only , not really update
// @Param InfoUpdate body models.InfoUpdate true "InfoUpdate"
// @Success 200 {string} ok
// @Router /info [patch]
func (a *InfoHandler) UpdateInfo(c *gin.Context) {
	body := models.InfoUpdate{}
	//gin example bind body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//call usecase method
	err := a.CUsecase.Update(&body)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}
