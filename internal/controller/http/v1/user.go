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

type userRoutes struct {
	t usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.User, l logger.Interface) {
	r := &userRoutes{t, l}

	h := handler.Group("/user")
	{
		h.GET("/", r.listUser)
		h.POST("/", r.store)
		h.PUT("/:id", r.update)
		h.DELETE("/:id", r.delete)
	}
}

// @Summary     Show user
// @Description Show all user user
// @ID          user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} response.UserResponse
// @Failure     500 {object} response
// @Router      /user/user [get]
func (r *userRoutes) listUser(c *gin.Context) {
	users, err := r.t.User(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - user")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, new(responseCustom.UserResponse).Makes(users))
}

func (r *userRoutes) store(c *gin.Context) {
	var request requestCustom.UserStore
	if err := c.ShouldBindJSON(&request); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			r.l.Error(err.Error(), "http - v1 - create user")
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
		entity.User{
			Name:     request.Name,
			Gender:   request.Gender,
			Address:  request.Address,
			Email:    request.Email,
			Password: request.Password,
		},
	)

	if err != nil {
		r.l.Error(err, "http - v1 - create user")
		errorResponse(c, http.StatusBadRequest, "invalid store to db, check DB user")

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (r *userRoutes) update(c *gin.Context) {
	var request requestCustom.UserUpdate

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			r.l.Error(err.Error(), "http - v1 - update user")
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

	userIDString := c.Param("id")
	userID, _ := strconv.ParseInt(userIDString, 0, 64)
	_, err := r.t.Update(
		c.Request.Context(),
		int(userID),
		entity.User{
			Name:     request.Name,
			Gender:   request.Gender,
			Address:  request.Address,
			Email:    request.Email,
			Password: request.Password,
		},
	)

	if err != nil {
		r.l.Error(err, "http - v1 - update user")
		errorResponse(c, http.StatusBadRequest, "invalid update to db, because not found, check DB user")

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (r *userRoutes) delete(c *gin.Context) {
	userIDString := c.Param("id")
	userID, _ := strconv.ParseInt(userIDString, 0, 64)
	err := r.t.Delete(
		c.Request.Context(),
		int(userID),
	)

	if err != nil {
		r.l.Error(err, "http - v1 - update user")
		errorResponse(c, http.StatusBadRequest, "invalid delete to db, because not found, check DB user")

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
