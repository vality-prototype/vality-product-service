package usecases

import (
	"errors"

	"github.com/vality-prototype/vality-user-service/app/domains/entities/models"
	"github.com/vality-prototype/vality-user-service/app/domains/repository"
	"github.com/vality-prototype/vality-user-service/app/domains/services"
	"github.com/vality-prototype/vality-utility-service/pkg"
)

type CreateUser interface {
	Create(user *models.User) (*models.User, error)
}

type createUser struct {
	userRepo CreateUser
	appCtx   pkg.AppContext
}

func NewCreateUser(appCtx pkg.AppContext, userRepo CreateUser) *createUser {
	return &createUser{
		userRepo: userRepo,
		appCtx:   appCtx,
	}
}

func (c *createUser) CreateUser(user *models.User) (*models.User, error) {
	var (
		checkUserService = services.NewCheckUser(c.appCtx, repository.NewUserRepo(c.appCtx))
	)
	isExist, err := checkUserService.IsEmailExist(user.Email)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, errors.New("email already exist")
	}
	return c.userRepo.Create(user)
}
