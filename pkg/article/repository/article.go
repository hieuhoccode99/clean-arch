package repository

import (
	"clean-arch/domain"
	"clean-arch/domain/models"
	"clean-arch/pkg/article/dto"
	"context"
	"gorm.io/gorm"
)

type articleRepo struct {
	Conn *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) IArticleRepository {
	return &articleRepo{conn}
}

type IArticleRepository interface {
	Get(ctx context.Context, req dto.GetArticleRequest) (res models.Article, err error)
}

func (m *articleRepo) fetch(ctx context.Context, query string, args ...interface{}) (result []models.Article, err error) {
	return result, nil
}

func (m *articleRepo) Get(ctx context.Context, req dto.GetArticleRequest) (res models.Article, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM pkg WHERE title = ?`

	list, err := m.fetch(ctx, query, req.Title)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}
