package handlers

import (
	"net/http"
	"strconv"

	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "user id error", err)
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid id param", err)
		return
	}

	var body models.TodoItem
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "parse error", err)
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, body)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "todo item create error", err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "user id error", err)
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid id param", err)
		return
	}
	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "get all items error", err)
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: items,
	})
}

func (h *Handler) getItemByID(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "user id error", err)
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid list id param", err)
		return
	}
	itemId, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid item id param", err)
		return
	}
	list, err := h.services.TodoItem.GetById(userId, listId,itemId)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "get item by id error", err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": list,
	})
}

func (h *Handler) updateItem(c *gin.Context) {

	userId, err := getUserID(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "user id error", err)
		return
	}


	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid id param", err)
		return
	} 
	itemId, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid item id param", err)
		return
	}

	var body models.UpdateItemInput
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "parse error", err)
		return
	}

	err = h.services.TodoItem.Update(userId, listId,itemId, body)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "update item error", err)
		return
	} 
	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {

	userId, err := getUserID(c)
	if err != nil {	
		h.handleErrorResponse(c, http.StatusInternalServerError, "user id error", err)
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid id param", err)
		return
	} 
	itemId, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid item id param", err)
		return
	}
	err = h.services.TodoItem.Delete(userId, listId,itemId)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "delete item error", err)
		return
	}
	c.JSON(http.StatusOK,statusResponse{
		Status: "ok",
	})
}

