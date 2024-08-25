package db

import (
	"sync"

	"github.com/mayureshucsb2019/bookstore/service/common"
)

var customerRepoInstance *CustomerRepository
var customerRepoOnce sync.Once

func NewCustomerRepository(db *common.DBConnection) *CustomerRepository {
	customerRepoOnce.Do(func() {
		customerRepoInstance = &CustomerRepository{
			DB: db.DB,
		}
	})
	return customerRepoInstance
}
