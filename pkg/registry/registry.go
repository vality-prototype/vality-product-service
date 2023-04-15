package registry

import (
	"github.com/vality-prototype/vality-product-service/pkg/adapters/controller"
	"github.com/vality-prototype/vality-utility-service/pkg"
)

type registry struct {
	appCtx pkg.AppContext
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(appCtx pkg.AppContext) Registry {
	return &registry{
		appCtx: appCtx,
	}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Product: r.NewProductController(),
	}
}
