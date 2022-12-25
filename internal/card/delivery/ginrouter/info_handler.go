package ginrouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/internal/models"
)

// InfoHandler  represent the httphandler for Info
type InfoHandler struct { // not interface?
	CUsecase models.InfoUsecase
}

// NewInfoHandler will initialize the service/api/v1/ endpoint
func NewInfoHandler(base *gin.RouterGroup, us models.InfoUsecase) {
	handler := &InfoHandler{
		CUsecase: us,
	}

	base.GET("/info", handler.GetInfo) //e.g.: http://127.0.0.1:8888/service/api/v1/info
	base.PUT("/info", handler.UpdateInfo)

}

// @Summary Get info API
// @Schemes http xx
// @Description Get info API
// @Tags infos
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /service/api/v1/info [get]
func (a *InfoHandler) GetInfo(c *gin.Context) {
	//call usecase method
	info, _ := a.CUsecase.GetInfo()

	c.JSON(http.StatusOK, info)
}

// @Summary put info API
// @Schemes http
// @Description put info API service name
// @Tags infos
// @Param InfoUpdate body models.InfoUpdate true "InfoUpdate"
// @Success 200 {string} ok
// @Router /service/api/v1/info [put]
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
