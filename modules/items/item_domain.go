package users

import (
	"test-pari/modules/items/model"
	"test-pari/schemas"
)

type Repository interface {
	CreateItemRepository(item *model.Items) (error, int64)
	GetItemsRepository() (*[]model.Items, error)
	GetItemByIDRepository(id int64) (*model.Items, error)
	UpdateItemByIDRepository(id int64, item *model.Items) error
	DeleteItemByIDRepository(id int64) error
}

type Service interface {
	CreateItemService(item schemas.CreateItemRequest) (error, int64)
	GetItemsService() (*[]model.Items, error)
	GetItemByIDService(id int64) (*model.Items, error)
	UpdateItemByIDService(id int64, item schemas.UpdateItemRequest) error
	DeleteItemByIDService(id int64) error
}
