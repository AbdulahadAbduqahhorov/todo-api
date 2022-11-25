package handlers

import (
	"net/http"
	"strconv"

	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"

	"github.com/gin-gonic/gin"
)
type statusResponse struct{
	Status string `json:"status"`
}

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "user id error", err)
		return
	}
	var body models.TodoList
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "parse error", err)
		return
	}

	id, err := h.services.TodoList.Create(userId, body)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "todo list create error", err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct{
	Data []models.TodoList `json:"data"`
}
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "user id error", err)
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "get lists error", err)
		return
	}
	c.JSON(http.StatusOK,getAllListsResponse{
		Data: lists,
	})
}
func (h *Handler) getListByID(c *gin.Context) {
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
	list, err := h.services.TodoList.GetById(userId, listId)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "get list by id error", err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": list,
	})
}

func (h *Handler) updateList(c *gin.Context) {

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

	var body models.UpdateListInput
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "parse error", err)
		return
	}

	err = h.services.TodoList.Update(userId, listId, body)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "update list error", err)
		return
	} 
	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})


}


func (h *Handler) deleteList(c *gin.Context) {

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
	err = h.services.TodoList.Delete(userId, listId)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "delete list error", err)
		return
	}
	c.JSON(http.StatusOK,statusResponse{
		Status: "ok",
	})
}
