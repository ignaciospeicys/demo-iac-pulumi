package secondary

import (
	"demo-pulumi-aws/model"
	"gorm.io/gorm"
)

type ResourceStackRepository struct {
	db *gorm.DB
}

func NewResourceStackRepository(db *gorm.DB) *ResourceStackRepository {
	return &ResourceStackRepository{db: db}
}

// Create inserts a new ResourceStack into the database.
func (rs *ResourceStackRepository) Create(resourceStack *model.ResourceStack) error {
	return rs.db.Create(resourceStack).Error
}

// FindByID returns a single ResourceStack matching the given ID.
func (rs *ResourceStackRepository) FindByID(id uint) (*model.ResourceStack, error) {
	var resourceStack model.ResourceStack
	err := rs.db.First(&resourceStack, id).Error
	return &resourceStack, err
}

// Update modifies an existing ResourceStack.
func (rs *ResourceStackRepository) Update(resourceStack *model.ResourceStack) error {
	return rs.db.Save(resourceStack).Error
}

// Delete removes a ResourceStack from the database.
func (rs *ResourceStackRepository) Delete(id uint) error {
	return rs.db.Delete(&model.ResourceStack{}, id).Error
}
