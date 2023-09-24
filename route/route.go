package route

import (
	config2 "clean-arch/config"
	"clean-arch/infrastructure"
	"clean-arch/pkg/article/controller/http"
	middleware2 "clean-arch/pkg/middleware"
	v1 "clean-arch/route/v1"
	"github.com/labstack/echo"
	"log"
)

func NewRouters() {
	config := config2.NewConfig()
	dbConn := infrastructure.NewDatabase(config)

	e := echo.New()
	middL := middleware2.InitMiddleware(config)
	articleCtrl := http.NewArticleHandler(dbConn.DB)

	v1.NewArticleRoutes(e, articleCtrl, middL)
	log.Fatal(e.Start(":8010"))
}
