package test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/vality-prototype/vality-product-service/pkg/entities/models"
	"github.com/vality-prototype/vality-product-service/pkg/registry"
	"github.com/vality-prototype/vality-utility-service/configs"
	"github.com/vality-prototype/vality-utility-service/pkg"

	// "github.com/go-playground/assert/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// this is integration test

var _ = Describe("CreateProd", func() {
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
	It("should create prod", func() {
		var (
			r    = registry.NewRegistry(appCtx)
			prod = &models.Product{
				Name:  "Prod1",
				Price: 1000,
			}
			appController = r.NewAppController()
		)
		resp, err := appController.Product.CreateProduct(context.Background(), prod)
		if err != nil {
			Fail(err.Error())
		}
		data, ok := resp.Data.(*models.Product)
		Expect(ok).To(Equal(true))
		Expect(data.Name).To(Equal(prod.Name))
		Expect(data.Price).To(Equal(prod.Price))
		Expect(data.ID).NotTo(Equal(0))

		err = db.Model(data).Delete(data).Error
		if err != nil {
			Fail(err.Error())
		}
	})
})
