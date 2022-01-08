package handler

import (
	structs "Golangcrud/Structs"
	"net/http"
	"github.com/gin-gonic/gin"
)

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
