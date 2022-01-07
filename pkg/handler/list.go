package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {

}
func (h *Handler) getAllList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
func (h *Handler) getListById(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
