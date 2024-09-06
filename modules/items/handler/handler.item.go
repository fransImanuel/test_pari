package handler

import (
	"fmt"
	"net/http"
	"strconv"
	items "test-pari/modules/items"
	"test-pari/schemas"
	"test-pari/utils"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	ItemService items.Service
}

func InitItemHandler(g *gin.Engine, itemService items.Service) {
	handler := &ItemHandler{
		ItemService: itemService,
	}

	routeAPI := g.Group("/api/v1/items")
	routeAPI.GET("/", handler.GetAllItemsHandler)
	routeAPI.GET("/:id", handler.GetItemByIDHandler)
	routeAPI.POST("/create", handler.CreateItemHandler)
	routeAPI.PUT("/:id", handler.UpdateItemByIDHandler)
	routeAPI.DELETE("/:id", handler.DeleteItemByIDHandler)
}

// Create Item
// @Tags Items
// @Summary Create Item
// @Description Create Item
// @ID Item-Create
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.CreateItemRequest true "body data"
// @Success 200  {object} schemas.Response
// @Router /v1/items/create [post]
func (h *ItemHandler) CreateItemHandler(c *gin.Context) {
	var req schemas.CreateItemRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Required field is empty", nil)
		return
	}

	err, ID := h.ItemService.CreateItemService(req)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Create Item", map[string]interface{}{
		"id": ID,
	})

}

// Get Item
// @Tags Items
// @Summary Get Item
// @Description Get Item
// @ID Item-Get
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200  {object} schemas.Response
// @Router /v1/items/ [get]
func (h *ItemHandler) GetAllItemsHandler(c *gin.Context) {
	items, err := h.ItemService.GetItemsService()
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Get Items", items)
}

// Get Item By ID
// @Tags Items
// @Summary Get Item By ID
// @Description Get Item By id
// @ID ItemByID-Get
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Item ID"
// @Success 200  {object} schemas.Response
// @Router /v1/items/{id} [get]
func (h *ItemHandler) GetItemByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Param Path Empty", nil)
		return
	}

	intValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", fmt.Sprintf("Error converting string to int64: %v", err), nil)
		return
	}

	items, err := h.ItemService.GetItemByIDService(intValue)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Get Item By ID", items)
}

// Update Item By ID
// @Tags Items
// @Summary Update Item By ID
// @Description Update Item By id
// @ID ItemByID-Update
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Item ID"
// @Param data body schemas.UpdateItemRequest true "body data"
// @Success 200  {object} schemas.Response
// @Router /v1/items/{id} [put]
func (h *ItemHandler) UpdateItemByIDHandler(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Param Path Empty", nil)
		return
	}

	var req schemas.UpdateItemRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Required field is empty", nil)
		return
	}

	intValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", fmt.Sprintf("Error converting string to int64: %v", err), nil)
		return
	}

	if err := h.ItemService.UpdateItemByIDService(intValue, req); err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}

	utils.APIResponse(c, http.StatusOK, "success", "Success Update Item By ID", nil)
}

// Delete Item By ID
// @Tags Items
// @Summary Delete Item By ID
// @Description Delete Item By id
// @ID ItemByID-Delete
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Item ID"
// @Success 200  {object} schemas.Response
// @Router /v1/items/{id} [delete]
func (h *ItemHandler) DeleteItemByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Param Path Empty", nil)
		return
	}

	intValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", fmt.Sprintf("Error converting string to int64: %v", err), nil)
		return
	}

	if err := h.ItemService.DeleteItemByIDService(intValue); err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Delete Item By ID", nil)
}
