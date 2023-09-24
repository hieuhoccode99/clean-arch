package v1

import (
	"clean-arch/pkg/article/controller/http"
	"clean-arch/pkg/middleware"
	"github.com/labstack/echo"
)

func NewArticleRoutes(e *echo.Echo, ctrl *http.ArticleController, middL *middleware.GoMiddleware) {
	g := e.Group("/articles")
	g.GET("", ctrl.Get, middL.Auth(true))
}
