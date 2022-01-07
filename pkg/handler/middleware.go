package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx = "UserId"
)

func (h *Handler) userIdenity(c *gin.Context) {
	header := c.GetHeader(authHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	headerparts := strings.Split(header, " ")
	if len(headerparts) != 2 || headerparts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	UserId, err := h.services.Authorization.Parsetoken(headerparts[1])
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return;
	}
	c.Set(userCtx, UserId)
}
