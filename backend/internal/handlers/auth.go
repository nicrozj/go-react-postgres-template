package handlers

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandlers struct {
	svc services.AuthServiceInterface
}

func NewAuthHandlers() *AuthHandlers {
	return &AuthHandlers{
		svc: services.NewAuthService(),
	}
}

func (h AuthHandlers) Greet(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello!")
}

func (h AuthHandlers) CreateUser(c *gin.Context) {
	var body models.AuthRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("please provide valid input")))
		return
	}

	result, err := h.svc.CreateUser(&body)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}

func (h AuthHandlers) LogoutUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, err := h.svc.LogoutUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}

func (h AuthHandlers) DeleteUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, err := h.svc.DeleteUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}
func (h AuthHandlers) GetUserByID(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, err := h.svc.GetUserByID(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}
func (h AuthHandlers) LoginUser(c *gin.Context) {
	var body *models.AuthRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("please provide valid input")))
		return
	}

	tokensResponse, err := h.svc.LoginUser(body)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("refresh_token", tokensResponse.RefreshToken, 0, "", config.Envs.WEB_URL, true, true)
	c.JSON(http.StatusOK, &models.Response{Success: true, Status: http.StatusOK, Data: tokensResponse})
}

func (h AuthHandlers) RefreshToken(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil || cookie == "" {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusUnauthorized, fmt.Errorf("please login again")))
	}

	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))
	tokensResponse, er := h.svc.GenerateTokens(userID, cookie)

	if er != nil {
		c.JSON(er.Status, er)
		return
	}

	c.SetCookie("refresh_token", tokensResponse.RefreshToken, 0, "", config.Envs.WEB_URL, true, true)
	c.JSON(http.StatusOK, &models.Response{Success: true, Status: http.StatusOK, Data: tokensResponse})
}
