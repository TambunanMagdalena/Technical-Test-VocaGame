package repositories

import (
	"context"
	"template-go/app/models"
)

type itemRepository repository

type ItemInterface interface {
	CreateMasterItem(ctx context.Context, item models.MasterItem) (models.MasterItem, error)
	GetMasterItemByFilter(ctx context.Context, filter models.FilterRequest) ([]models.MasterItem, error)
	UpdateMasterItem(ctx context.Context, code, name string, update models.MasterItem) error
	DeleteMasterItem(ctx context.Context, code, name string) error
	ExportData(ctx context.Context, filter models.FilterRequest) ([]models.MasterItem, error)
}

func (r *itemRepository) CreateMasterItem(ctx context.Context, item models.MasterItem) (models.MasterItem, error) {
	err := r.Options.Postgres.Create(&item).Error
	return item, err
}

func (r *itemRepository) GetMasterItemByFilter(ctx context.Context, filter models.FilterRequest) ([]models.MasterItem, error) {
	var items []models.MasterItem
	db := r.Options.Postgres
	if filter.Code != "" {
		db = db.Where("code = ?", filter.Code)
	}
	if filter.Name != "" {
		db = db.Where("name = ?", filter.Name)
	}
	if filter.Category != "" {
		db = db.Where("category = ?", filter.Category)
	}
	err := db.Find(&items).Error
	return items, err
}

func (r *itemRepository) UpdateMasterItem(ctx context.Context, code, name string, update models.MasterItem) error {
	return r.Options.Postgres.Model(&models.MasterItem{}).
		Where("code = ? AND name = ?", code, name).
		Updates(update).Error
}

func (r *itemRepository) DeleteMasterItem(ctx context.Context, code, name string) error {
	return r.Options.Postgres.Where("code = ? AND name = ?", code, name).
		Delete(&models.MasterItem{}).Error
}

func (r *itemRepository) ExportData(ctx context.Context, filter models.FilterRequest) ([]models.MasterItem, error) {
	var items []models.MasterItem
	db := r.Options.Postgres
	if filter.Code != "" {
		db = db.Where("code = ?", filter.Code)
	}
	if filter.Name != "" {
		db = db.Where("name = ?", filter.Name)
	}
	if filter.Category != "" {
		db = db.Where("category = ?", filter.Category)
	}
	err := db.Find(&items).Error
	return items, err
}
