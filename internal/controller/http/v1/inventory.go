package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	requestCustom "github.com/evrone/go-clean-template/internal/controller/request"
	responseCustom "github.com/evrone/go-clean-template/internal/controller/response"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type inventoryRoutes struct {
	t usecase.Inventory
	l logger.Interface
}

func newInventoryRoutes(handler *gin.RouterGroup, t usecase.Inventory, l logger.Interface) {
	r := &inventoryRoutes{t, l}

	h := handler.Group("/inventory")
	{
		h.GET("/", r.listInventory)
		h.POST("/", r.store)
		h.PUT("/:id", r.update)
		h.DELETE("/:id", r.delete)
	}
}

// @Summary     Show inventory
// @Description Show all inventory inventory
// @ID          inventory
// @Tags  	    inventory
// @Accept      json
// @Produce     json
// @Success     200 {object} response.InventoryResponse
// @Failure     500 {object} response
// @Router      /inventory/inventory [get]
func (r *inventoryRoutes) listInventory(c *gin.Context) {
	inventorys, err := r.t.Inventory(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - inventory")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, new(responseCustom.InventoryResponse).Make(inventorys))
}

func (r *inventoryRoutes) store(c *gin.Context) {
	var request requestCustom.InventoryStore
	if err := c.ShouldBindJSON(&request); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			r.l.Error(err.Error(), "http - v1 - create inventory")
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something Wrong!!!"})
			return
		}

		var errorMessages = make(map[string]string)
		for _, e := range validationErr {
			fieldJSONName := request.GetJsonFieldName(e.Field())
			errorMessages[fieldJSONName] = request.ErrMessages()[fieldJSONName][e.ActualTag()]
		}
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"Errors": errorMessages},
		)

		return
	}

	_, err := r.t.Create(
		c.Request.Context(),
		entity.Inventory{
			Name:        request.Name,
			UserID:      request.UserID,
			Stock:       request.Stock,
			Price:       request.Price,
			Description: request.Description,
		},
	)

	if err != nil {
		r.l.Error(err, "http - v1 - create inventory")
		errorResponse(c, http.StatusBadRequest, "invalid store to db, check DB inventory")

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (r *inventoryRoutes) update(c *gin.Context) {
	var request requestCustom.InventoryUpdate

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			r.l.Error(err.Error(), "http - v1 - update inventory")
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something Wrong!!!"})
			return
		}

		var errorMessages = make(map[string]string)
		for _, e := range validationErr {
			fieldJSONName := request.GetJsonFieldName(e.Field())
			errorMessages[fieldJSONName] = request.ErrMessages()[fieldJSONName][e.ActualTag()]
		}
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"Errors": errorMessages},
		)

		return
	}

	inventoryIDString := c.Param("id")
	inventoryID, _ := strconv.ParseInt(inventoryIDString, 0, 64)
	_, err := r.t.Update(
		c.Request.Context(),
		int(inventoryID),
		entity.Inventory{
			Name:        request.Name,
			UserID:      request.UserID,
			Stock:       request.Stock,
			Price:       request.Price,
			Description: request.Description,
		},
	)

	if err != nil {
		r.l.Error(err, "http - v1 - update inventory")
		errorResponse(c, http.StatusBadRequest, "invalid update to db, because not found, check DB inventory")

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (r *inventoryRoutes) delete(c *gin.Context) {
	inventoryIDString := c.Param("id")
	inventoryID, _ := strconv.ParseInt(inventoryIDString, 0, 64)
	err := r.t.Delete(
		c.Request.Context(),
		int(inventoryID),
	)

	if err != nil {
		r.l.Error(err, "http - v1 - update inventory")
		errorResponse(c, http.StatusBadRequest, "invalid delete to db, because not found, check DB inventory")

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
