package repository

import (
	"errors"
	"test-pari/constant"
	items "test-pari/modules/items"
	"test-pari/modules/items/model"

	"gorm.io/gorm"
)

type ItemRepository struct {
	DBPostgres *gorm.DB
	//DBMongoDB
	//DBMinio, etc
}

func InitItemRepository(db *gorm.DB) items.Repository {
	return &ItemRepository{
		DBPostgres: db,
	}
}

func (u *ItemRepository) CreateItemRepository(item *model.Items) (error, int64) {
	db := u.DBPostgres

	item.InitAudit(constant.OPERATION_SQL_INSERT)

	results := db.Create(&item)
	if results.Error != nil {
		return results.Error, 0
	}

	return nil, item.ID

}

func (u *ItemRepository) GetItemsRepository() (*[]model.Items, error) {
	var items *[]model.Items
	db := u.DBPostgres

	// Get all records
	results := db.Find(&items)
	// SELECT * FROM ITEMS;
	if results.Error != nil {
		return nil, results.Error
	}

	return items, nil
}

func (u *ItemRepository) GetItemByIDRepository(id int64) (*model.Items, error) {
	var items model.Items
	db := u.DBPostgres

	// Get all records
	results := db.First(&items, id)
	// SELECT * FROM ITEMS;
	if results.Error != nil {
		return nil, results.Error
	}

	return &items, nil
}

func (u *ItemRepository) UpdateItemByIDRepository(id int64, item *model.Items) error {
	var items model.Items
	db := u.DBPostgres

	result := db.Model(&items).Where("id = ?", id).Updates(model.Items{Name: item.Name, Description: item.Description, Price: item.Price, Qty: item.Qty, IsActive: item.IsActive})

	if result.RowsAffected < 1 {
		return errors.New("No Data Affected")
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *ItemRepository) DeleteItemByIDRepository(id int64) error {
	db := u.DBPostgres

	if err := db.Delete(&model.Items{}, id).Error; err != nil {
		return err
	}

	return nil
}
