package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leeexeo/kuchiki/pkg/common/exception"
	"github.com/leeexeo/kuchiki/pkg/domain"
	"github.com/leeexeo/kuchiki/pkg/service"
)

func validateAddUserInput(input *domain.AddUserInput) error {
	if input.Password == "" {
		return exception.ErrUserPasswordEmpty
	}
	return nil
}

func AddUser(c *gin.Context) {
	var input domain.AddUserInput
	if err := c.Bind(&input); err != nil {
		Response(c, nil, err)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := validateAddUserInput(&input); err != nil {
		Response(c, nil, err)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	output, err := service.AddUser(c, &input)
	Response(c, output, err)
}

func validateDeleteUserInput(input *domain.DeleteUserInput) error {
	return nil
}
func DeleteUser(c *gin.Context) {
	var input domain.DeleteUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		Response(c, nil, err)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := validateDeleteUserInput(&input); err != nil {
		Response(c, nil, err)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := service.DeleteUser(c, &input)
	Response(c, nil, err)
}
