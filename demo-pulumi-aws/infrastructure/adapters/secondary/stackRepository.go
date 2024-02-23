package secondary

import (
	"demo-pulumi-aws/model"
	"gorm.io/gorm"
)

type StackRepository struct {
	db *gorm.DB
}

func NewStackRepository(db *gorm.DB) *StackRepository {
	return &StackRepository{db: db}
}

// Create inserts a new Stack into the database.
func (s *StackRepository) Create(stack *model.Stack) error {
	return s.db.Create(stack).Error
}

// FindByID returns a single Stack matching the given ID.
func (s *StackRepository) FindByID(id uint) (*model.Stack, error) {
	var stack model.Stack
	err := s.db.First(&stack, id).Error
	return &stack, err
}

// Update modifies an existing Stack.
func (s *StackRepository) Update(stack *model.Stack) error {
	return s.db.Save(stack).Error
}

// Delete removes a Stack from the database.
func (s *StackRepository) Delete(id uint) error {
	return s.db.Delete(&model.Stack{}, id).Error
}
