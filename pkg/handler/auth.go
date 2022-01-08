package handler

import (
	structs "Golangcrud/Structs"
	"net/http"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body structs.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input structs.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, 400, err.Error())
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type SignInType struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {

	var input SignInType

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, 400, err.Error())
	}

	token, err := h.services.Authorization.GenerateToken(input.Username,input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
