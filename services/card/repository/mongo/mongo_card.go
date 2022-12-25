package mongo

import (
	"context"

	"github.com/minilabmemo/go-rest-arch/services/models"
)

type mysqlArticleRepository struct {
	//Conn *sql.DB
	Conn bool
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
// func NewMysqlArticleRepository(Conn *sql.DB) domain.ArticleRepository {
// 	return &mysqlArticleRepository{Conn}
// }
func NewMongoCardRepository(Conn bool) models.ArticleRepository {
	return &mysqlArticleRepository{Conn}
}

func (m *mysqlArticleRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []models.Card, err error) {

	return
}
func (m *mysqlArticleRepository) Fetch(ctx context.Context, cursor string, num int64) (res []models.Card, nextCursor string, err error) {
	// query := `SELECT id,title,content, author_id, updated_at, created_at
	// 						FROM article WHERE created_at > ? ORDER BY created_at LIMIT ? `

	// decodedCursor, err := repository.DecodeCursor(cursor)
	// if err != nil && cursor != "" {
	// 	return nil, "", domain.ErrBadParamInput
	// }

	// res, err = m.fetch(ctx, query, decodedCursor, num)
	// if err != nil {
	// 	return nil, "", err
	// }

	// if len(res) == int(num) {
	// 	nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	// }

	return
}
