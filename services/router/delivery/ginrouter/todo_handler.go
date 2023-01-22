package ginrouter

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/services/models"
)

type TodoHandler struct {
	CardUsecase models.CardUsecase
}

// NewTodoHandler will initialize the service/api/v1/ endpoint
func NewTodoHandler(base *gin.RouterGroup, cu models.CardUsecase) {
	handler := &TodoHandler{
		CardUsecase: cu,
	}
	base.POST("/cards", handler.AddCard)
	base.GET("/cards", handler.GetCards)
	base.PUT("/cards/:id", handler.UpdateCard)
	base.DELETE("/cards/:id", handler.DeleteCard)

}

// @Tags card
// @Summary post Card API
// @Schemes http
// @Description post Card
// @Param CardUpdate body models.CardUpdate true "CardUpdate"
// @Success 200 {string} ok
// @Router /cards [post]
func (a *TodoHandler) AddCard(c *gin.Context) {
	body := models.CardUpdate{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id, err := a.CardUsecase.Store(context.TODO(), &body)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
		return
	}

	c.JSON(http.StatusOK, id)
}

// @Tags card
// @Summary Get info API
// @Description Get info API
// @Success 200 {string} ok
// @Router /cards [get]
func (a *TodoHandler) GetCards(c *gin.Context) {
	//call usecase method
	res, _, err := a.CardUsecase.Fetch(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Tags card
// @Summary put Card API
// @Description put Card API
// @Param  id  path   string  true  "Card ID"
// @Param CardUpdate body models.CardUpdate true "CardUpdate"
// @Success 200 {string} ok
// @Router /cards/{id} [put]
func (a *TodoHandler) UpdateCard(c *gin.Context) {

	body := models.CardUpdate{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	_, err := a.CardUsecase.Update(c.Request.Context(), id, &body)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// @Tags card
// @Summary Delete Card API
// @Description Delete Card API
// @Param  id  path   string  true  "Card ID"
// @Success 200 {string} ok
// @Router /cards/{id} [delete]
func (a *TodoHandler) DeleteCard(c *gin.Context) {

	id := c.Param("id")
	_, err := a.CardUsecase.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
