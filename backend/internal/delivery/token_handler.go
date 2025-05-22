package delivery

import (
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	TokenUC usecase.TokenUsecase
}

func NewTokenHandler(tokenUC usecase.TokenUsecase) *TokenHandler {
	return &TokenHandler{TokenUC: tokenUC}
}

func (h *TokenHandler) RefreshToken(c *gin.Context)
