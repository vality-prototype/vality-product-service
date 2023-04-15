package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/vality-prototype/vality-user-service/app/domains/entities/models"
	"github.com/vality-prototype/vality-user-service/app/domains/repository"
	"github.com/vality-prototype/vality-user-service/app/usecases"
	"github.com/vality-prototype/vality-utility-service/configs"
	"github.com/vality-prototype/vality-utility-service/pkg"

	// "github.com/go-playground/assert/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// this is integration test

// func TestCreateUser(t *testing.T) {
// 	var (
// 		env = configs.Init(".env-ut")
// 		// init sql connection, this connection will keep alive until the app is closed
// 		connStr = env.DBConnectionStr
// 		db, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
// 		appCtx  = pkg.NewAppContext(db, nil, &env)
// 	)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// create user
// 	user := &models.User{
// 		Email:    "",
// 		Password: "",
// 	}
// 	userRepo := repository.NewUserRepo(appCtx)
// 	createUser := usecases.NewCreateUser(appCtx, userRepo)
// 	userCreated, err := createUser.CreateUser(user)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.Equal(t, user.Email, userCreated.Email)
// 	assert.Equal(t, user.Password, userCreated.Password)
// 	assert.NotEqual(t, user.ID, 0)

// 	err = db.Delete(userCreated).Error
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

var _ = Describe("CreateUser", func() {
	var (
		env    configs.Env
		db     *gorm.DB
		appCtx pkg.AppContext
		err    error
	)
	BeforeEach(func() {
		env = configs.Init(".env-ut")
		connStr := env.DBConnectionStr
		db, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
		if err != nil {
			Fail(err.Error())
		}
		appCtx = pkg.NewAppContext(db, nil, &env)
	})
	AfterEach(func() {
		current, err := db.DB()
		if err != nil {
			Fail(err.Error())
		}
		err = current.Close()
		if err != nil {
			Fail(err.Error())
		}
	})
	It("should create user", func() {
		var (
			user = &models.User{
				Email:    "",
				Password: "",
			}
			userRepo   = repository.NewUserRepo(appCtx)
			createUser = usecases.NewCreateUser(appCtx, userRepo)
		)

		userCreated, err := createUser.CreateUser(user)
		if err != nil {
			Fail(err.Error())
		}
		Expect(user.Email).To(Equal(userCreated.Email))
		Expect(user.Password).To(Equal(userCreated.Password))
		Expect(user.ID).NotTo(Equal(0))
		// assert.Equal(t, user.Email, userCreated.Email)
		// assert.Equal(t, user.Password, userCreated.Password)
		// assert.NotEqual(t, user.ID, 0)

		err = db.Delete(userCreated).Error
		if err != nil {
			Fail(err.Error())
		}
	})
})
