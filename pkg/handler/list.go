package handler

import (
	structs "Golangcrud/Structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		return
	}

	var input structs.Todolist

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.services.Todolist.Create(userid, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
func (h *Handler) getAllList(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.Todolist.GetAll(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lists": lists,
	})

}

func (h *Handler) getListById(c *gin.Context) {
	id := c.Param("id")
	userlist, err := h.services.Todolist.GetListById(id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "User not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": userlist,
	})
}
func (h *Handler) updateList(c *gin.Context) {
	var input structs.UpdateListItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id := c.Param("id")
	userlist, err := h.services.Todolist.UpdateList(input, id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "User not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": userlist,
	})
}
func (h *Handler) deleteList(c *gin.Context) {
	id := c.Param("id")
	userlist, err := h.services.Todolist.DeleteList(id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "User not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": userlist,
	})
}
