package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	requestCustom "github.com/evrone/go-clean-template/internal/controller/request"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/hashing"
	"github.com/evrone/go-clean-template/pkg/logger"
)

type authRoutes struct {
	t usecase.Auth
	l logger.Interface
}

func AuthRoutes(handler *gin.RouterGroup, t usecase.Auth, l logger.Interface) {
	r := &authRoutes{t, l}

	h := handler.Group("/auth")
	{
		h.POST("/login", r.login)
	}
}

func (r *authRoutes) login(c *gin.Context) {
	var request requestCustom.Login

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			r.l.Error(err.Error(), "http - v1 - login")
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

	user, err := r.t.Login(
		c.Request.Context(),
		request.Email,
		request.Password,
	)

	if user == (entity.User{}) {
		errorResponse(c, http.StatusBadRequest, "User not found")

		return
	} else if err != nil {
		r.l.Error(err, "http - v1 - update user")
		errorResponse(c, http.StatusBadRequest, "invalid update to db, because not found, check DB user")

		return
	}

	token := hashing.GenerateToken(user.Email, user.Password)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
