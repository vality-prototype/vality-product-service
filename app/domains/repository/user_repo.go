package repository

import (
	"github.com/vality-prototype/vality-user-service/app/domains/entities/models"
	"github.com/vality-prototype/vality-utility-service/pkg"
)

type userRepo struct {
	appCtx pkg.AppContext
}

func NewUserRepo(appCtx pkg.AppContext) *userRepo {
	return &userRepo{
		appCtx: appCtx,
	}
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	var (
		user models.User
		db   = r.appCtx.GetMainDBConnection()
	)
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Create(user *models.User) (*models.User, error) {
	var (
		db          = r.appCtx.GetMainDBConnection()
		userCreated models.User
	)
	err := db.Create(user).Find(&userCreated).Error
	return &userCreated, err
}

func (r *userRepo) FindAll() ([]*models.User, error) {
	var (
		users []*models.User
		db    = r.appCtx.GetMainDBConnection()
	)
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
