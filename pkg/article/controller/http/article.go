package http

import (
	"clean-arch/pkg/article/dto"
	"clean-arch/pkg/article/repository"
	"clean-arch/pkg/article/usecase"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"net/http"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type ArticleController struct {
	ArticleuCase usecase.IArticleUCase
}

func NewArticleHandler(conn *gorm.DB) *ArticleController {
	articleRepo := repository.NewArticleRepository(conn)
	return &ArticleController{
		ArticleuCase: usecase.NewArticleUCase(articleRepo),
	}
}

func (a *ArticleController) Get(c echo.Context) error {
	req := dto.GetArticleRequest{}
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	res, err := a.ArticleuCase.Get(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
