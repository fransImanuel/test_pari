package service

import (
	items "test-pari/modules/items"
	"test-pari/modules/items/model"
	"test-pari/schemas"
)

type ItemService struct {
	ItemsRepository items.Repository
}

func InitItemsService(ItemsRepository items.Repository) items.Service {
	return &ItemService{
		ItemsRepository: ItemsRepository,
	}
}

func (i *ItemService) CreateItemService(item schemas.CreateItemRequest) (error, int64) {
	itemModel := &model.Items{
		Name:        &item.Name,
		Description: &item.Description,
		Price:       &item.Price,
		Qty:         &item.Qty,
	}
	err, id := i.ItemsRepository.CreateItemRepository(itemModel)
	if err != nil {
		return err, 0
	}

	return nil, id
}

func (i *ItemService) GetItemsService() (*[]model.Items, error) {
	items, err := i.ItemsRepository.GetItemsRepository()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (i *ItemService) GetItemByIDService(id int64) (*model.Items, error) {
	item, err := i.ItemsRepository.GetItemByIDRepository(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *ItemService) UpdateItemByIDService(id int64, item schemas.UpdateItemRequest) error {
	itemModel := &model.Items{
		Name:        &item.Name,
		Description: &item.Description,
		Price:       &item.Price,
		Qty:         &item.Qty,
		IsActive:    &item.IsActive,
	}
	if err := i.ItemsRepository.UpdateItemByIDRepository(id, itemModel); err != nil {
		return err
	}

	return nil
}

func (i *ItemService) DeleteItemByIDService(id int64) error {
	if err := i.ItemsRepository.DeleteItemByIDRepository(id); err != nil {
		return err
	}

	return nil
}
