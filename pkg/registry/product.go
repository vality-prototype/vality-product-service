package registry

import (
	"github.com/vality-prototype/vality-product-service/pkg/adapters/controller"
	"github.com/vality-prototype/vality-product-service/pkg/adapters/repository"
	"github.com/vality-prototype/vality-product-service/pkg/usecases"
)

func (r *registry) NewProductController() controller.ProductController {
	var (
		repo = repository.NewProductRepo(r.appCtx)
		u    = usecases.NewCreateUser(r.appCtx, repo)
	)
	return controller.NewProductController(r.appCtx, u)
}
