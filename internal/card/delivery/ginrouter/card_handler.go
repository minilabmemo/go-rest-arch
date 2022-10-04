package ginrouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/internal/models"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type CardHandler struct {
	CUsecase models.CardUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(base *gin.RouterGroup, us models.CardUsecase) {
	handler := &CardHandler{
		CUsecase: us,
	}

	baseCardGrp := base.Group("card")
	{
		baseCardGrp.GET("/articles", handler.FetchArticle)
	}
	// e.POST("/articles", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
}

// // FetchArticle will fetch the article based on given params
func (a *CardHandler) FetchArticle(c *gin.Context) {

	// a.CUsecase.Fetch()

	c.JSON(http.StatusOK, "OK")
}

// // GetByID will get article by given id
// func (a *ArticleHandler) GetByID(c echo.Context) error {
// 	idP, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
// 	}

// 	id := int64(idP)
// 	ctx := c.Request().Context()

// 	art, err := a.AUsecase.GetByID(ctx, id)
// 	if err != nil {
// 		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, art)
// }

// func isRequestValid(m *domain.Article) (bool, error) {
// 	validate := validator.New()
// 	err := validate.Struct(m)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// // Store will store the article by given request body
// func (a *ArticleHandler) Store(c echo.Context) (err error) {
// 	var article domain.Article
// 	err = c.Bind(&article)
// 	if err != nil {
// 		return c.JSON(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	var ok bool
// 	if ok, err = isRequestValid(&article); !ok {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	ctx := c.Request().Context()
// 	err = a.AUsecase.Store(ctx, &article)
// 	if err != nil {
// 		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusCreated, article)
// }

// // Delete will delete article by given param
// func (a *ArticleHandler) Delete(c echo.Context) error {
// 	idP, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
// 	}

// 	id := int64(idP)
// 	ctx := c.Request().Context()

// 	err = a.AUsecase.Delete(ctx, id)
// 	if err != nil {
// 		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}

// 	return c.NoContent(http.StatusNoContent)
// }

// func getStatusCode(err error) int {
// 	if err == nil {
// 		return http.StatusOK
// 	}

// 	logrus.Error(err)
// 	switch err {
// 	case domain.ErrInternalServerError:
// 		return http.StatusInternalServerError
// 	case domain.ErrNotFound:
// 		return http.StatusNotFound
// 	case domain.ErrConflict:
// 		return http.StatusConflict
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }
