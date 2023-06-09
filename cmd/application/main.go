package main

import (
	"log"

	"github.com/vality-prototype/vality-product-service/pkg/infrastructure/datastore"
	"github.com/vality-prototype/vality-product-service/pkg/infrastructure/transport"
	"github.com/vality-prototype/vality-product-service/pkg/registry"
	"github.com/vality-prototype/vality-utility-service/configs"
	"github.com/vality-prototype/vality-utility-service/pkg"
	"github.com/vality-prototype/vality-utility-service/pkg/s3_provider"
)

func main() {
	env := configs.Init(".env")

	// init sql connection, this connection will keep alive until the app is closed
	db := datastore.NewDB(env)
	sql, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sql.Close()

	provider := s3_provider.NewS3Provider(
		env.S3BucketName,
		env.S3Region,
		env.S3APIKey,
		env.S3Secret,
		env.S3Domain,
	)

	appCtx := pkg.NewAppContext(db, provider, &env)
	r := registry.NewRegistry(appCtx)
	transport.NewRouter(r.NewAppController())
}
