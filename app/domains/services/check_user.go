package services

import (
	"github.com/vality-prototype/vality-user-service/app/domains/entities/models"
	"github.com/vality-prototype/vality-utility-service/pkg"
	"gorm.io/gorm"
)

type CheckUser interface {
	FindByEmail(email string) (*models.User, error)
}

type checkUser struct {
	userRepo CheckUser
	appCtx   pkg.AppContext
}

func NewCheckUser(appCtx pkg.AppContext, userRepo CheckUser) *checkUser {
	return &checkUser{
		userRepo: userRepo,
		appCtx:   appCtx,
	}
}

func (c *checkUser) IsEmailExist(email string) (bool, error) {
	user, err := c.userRepo.FindByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if err == gorm.ErrRecordNotFound || user == nil {
		return false, nil
	}

	return true, nil
}
