package db

import (
	"sync"

	"github.com/mayureshucsb2019/bookstore/service/common"
)

var authorRepoInstance *AuthorRepository
var authorRepoOnce sync.Once

func NewAuthorRepository(db *common.DBConnection) *AuthorRepository {
	authorRepoOnce.Do(func() {
		authorRepoInstance = &AuthorRepository{
			DB: db.DB,
		}
	})
	return authorRepoInstance
}
