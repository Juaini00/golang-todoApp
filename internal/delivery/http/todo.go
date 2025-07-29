package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_app/internal/domain/entity"
	"todo_app/internal/usecase"
	"todo_app/pkg/utils"
)

type TodoRequest struct {
	Title string `json:"title" binding:"required"`
}

type TodoHandler struct {
	usecase usecase.TodoUsecaseImpl
}

func NewTodoHandler(u usecase.TodoUsecaseImpl) *TodoHandler {
	return &TodoHandler{usecase: u}
}

// Create adds a new todo.
// @Summary Create todo
// @Tags todo
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body TodoRequest true "New todo"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /todo [post]
func (h *TodoHandler) Create(c *gin.Context) {
	req := TodoRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildValidatorErrorResponse(http.StatusBadRequest, "Validation Error", err))
		return
	}
	userCtxVal, _ := c.Get("userContext")
	userCtx := userCtxVal.(*entity.User)
	todo, err := h.usecase.CreateTodo(c.Request.Context(), userCtx.ID, req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.BuildResponse(http.StatusCreated, "Todo created", todo))
}

// List returns user's todos.
// @Summary List todos
// @Tags todo
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /todo [get]
func (h *TodoHandler) List(c *gin.Context) {
	userCtxVal, _ := c.Get("userContext")
	userCtx := userCtxVal.(*entity.User)
	todos, err := h.usecase.ListTodo(c.Request.Context(), userCtx.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Todos fetched", todos))
}
