package service

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/leeexeo/kon/id"
	"github.com/leeexeo/kon/log"
	"github.com/leeexeo/kuchiki/pkg/common"
	"github.com/leeexeo/kuchiki/pkg/domain"
	"github.com/leeexeo/kuchiki/pkg/models"
)

func AddUser(c *gin.Context, input *domain.AddUserInput) (*domain.AddUserOutput, error) {
	// check RoleId

	// check avatar

	userId := id.NewWithPrefix("u-")
	user := &models.User{
		ControlBy: models.ControlBy{
			CreateBy: input.CreateBy,
			UpdateBy: input.UpdateBy,
		},
		Id:       userId,
		Username: input.Username,
		Password: input.Password,
		NickName: input.NickName,
		Phone:    input.Phone,
		RoleId:   input.RoleId,
		Salt:     input.Salt,
		Avatar:   input.Avatar,
		Sex:      input.Sex,
		Email:    input.Email,
		Remark:   input.Remark,
		Status:   common.UserStatusNormal,
	}
	if err := userDao.Save(user); err != nil {
		log.Error(context.TODO(), "save user failed", "name", user.Username)
		return nil, errors.New("save user failed")
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error: save user failed"})
	} else {
		return &domain.AddUserOutput{UserId: userId}, nil
	}
}

func DeleteUser(c *gin.Context, input *domain.DeleteUserInput) error {
	userId := input.UserId
	if err := userDao.Delete(userId); err != nil {
		log.Error(context.TODO(), "delete user failed", "id", userId)
		return errors.New("delete user failed")
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error: delete user failed"})
	}
	return nil
}
