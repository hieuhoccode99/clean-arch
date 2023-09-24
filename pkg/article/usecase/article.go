package usecase

import (
	"clean-arch/pkg/article/dto"
	"clean-arch/pkg/article/repository"
	"clean-arch/utils"
	"context"
)

type articleUCase struct {
	articleRepo repository.IArticleRepository
}

func NewArticleUCase(a repository.IArticleRepository) IArticleUCase {
	return &articleUCase{
		articleRepo: a,
	}
}

type IArticleUCase interface {
	Get(c context.Context, req dto.GetArticleRequest) (res dto.GetArticleResponse, err error)
}

func (a *articleUCase) Get(c context.Context, req dto.GetArticleRequest) (res dto.GetArticleResponse, err error) {
	listAr, err := a.articleRepo.Get(c, req)
	if err != nil {
		return
	}

	if err = utils.Copy(res, listAr); err != nil {
		return res, err
	}

	return res, nil
}
