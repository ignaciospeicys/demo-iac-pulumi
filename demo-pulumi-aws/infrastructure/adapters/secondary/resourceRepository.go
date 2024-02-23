package secondary

import (
	"demo-pulumi-aws/model"
	"gorm.io/gorm"
)

type ResourceRepository struct {
	db *gorm.DB
}

func NewResourceRepository(db *gorm.DB) *ResourceRepository {
	return &ResourceRepository{db: db}
}

// Create inserts a new Resource into the database.
func (r *ResourceRepository) Create(resource *model.Resource) error {
	return r.db.Create(resource).Error
}

// FindByID returns a single Resource matching the given ID.
func (r *ResourceRepository) FindByID(id uint) (*model.Resource, error) {
	var resource model.Resource
	err := r.db.Preload("Configurations").Preload("ResourceStacks").First(&resource, id).Error
	return &resource, err
}

// Update modifies an existing Resource.
func (r *ResourceRepository) Update(resource *model.Resource) error {
	return r.db.Save(resource).Error
}

// Delete removes a Resource from the database.
func (r *ResourceRepository) Delete(id uint) error {
	return r.db.Delete(&model.Resource{}, id).Error
}
