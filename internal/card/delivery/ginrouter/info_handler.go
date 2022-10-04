package ginrouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/internal/config"
	"github.com/minilabmemo/go-rest-arch/internal/models"
)

// ResponseError represent the reseponse error struct

// ArticleHandler  represent the httphandler for article
type InfoHandler struct {
	CUsecase models.InfoUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewInfoHandler(base *gin.RouterGroup, us models.InfoUsecase) {
	handler := &InfoHandler{
		CUsecase: us,
	}

	base.GET("/info", handler.FetchArticle) //http://127.0.0.1:8888/service/api/v1/info

}
func (a *InfoHandler) FetchArticle(c *gin.Context) {

	info, _ := a.CUsecase.Fetch(*config.ConfigData)

	c.JSON(http.StatusOK, info)
}

// func (*InfoHandler) Fetch() (models.Info, error) {

// }
