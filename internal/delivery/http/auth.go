package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todo_app/internal/usecase"
	"todo_app/pkg/utils"
)

type RegisterRequest struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}
type LoginRequest struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

type AuthResponse struct {
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
	Name     string `json:"name,omitempty"`
}

type AuthHandler struct {
	authUsecase usecase.UserUsecaseImpl
}

func NewAuthHandler(authUsecase usecase.UserUsecaseImpl) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

// Register handles user registration.
// @Summary Register new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	authRequest := RegisterRequest{}

	if err := c.ShouldBindJSON(&authRequest); err != nil {
		log.Println("Validation Error", err)
		c.JSON(http.StatusBadRequest, utils.BuildValidatorErrorResponse(http.StatusBadRequest, "Validation Error", err))
		return
	}

	user, err := h.authUsecase.Register(c.Request.Context(), authRequest.Name, authRequest.Username, authRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response := AuthResponse{
		Name:     user.Username,
		Username: user.Username,
	}

	c.JSON(200, utils.BuildResponse(http.StatusCreated, "Register was successfully", response))
}

// Login handles user login.
// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	authRequest := LoginRequest{}

	if err := c.ShouldBindJSON(&authRequest); err != nil {
		log.Println("Validation Error", err)
		c.JSON(http.StatusBadRequest, utils.BuildValidatorErrorResponse(http.StatusBadRequest, "Validation Error", err))
		return
	}

	user, token, err := h.authUsecase.Login(c.Request.Context(), authRequest.Username, authRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response := AuthResponse{
		Name:     user.Name,
		Username: user.Username,
		Token:    token,
	}

	c.JSON(200, utils.BuildResponse(http.StatusOK, "Login was successfully", response))
}
