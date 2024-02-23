package secondary

import (
	"demo-pulumi-aws/model"
	"gorm.io/gorm"
)

type ConfigurationRepository struct {
	db *gorm.DB
}

func NewConfigurationRepository(db *gorm.DB) *ConfigurationRepository {
	return &ConfigurationRepository{db: db}
}

// Create inserts a new Configuration into the database.
func (c *ConfigurationRepository) Create(configuration *model.Configuration) error {
	return c.db.Create(configuration).Error
}

// FindByID returns a single Configuration matching the given ID.
func (c *ConfigurationRepository) FindByID(id uint) (*model.Configuration, error) {
	var configuration model.Configuration
	err := c.db.First(&configuration, id).Error
	return &configuration, err
}

// Update modifies an existing Configuration.
func (c *ConfigurationRepository) Update(configuration *model.Configuration) error {
	return c.db.Save(configuration).Error
}

// Delete removes a Configuration from the database.
func (c *ConfigurationRepository) Delete(id uint) error {
	return c.db.Delete(&model.Configuration{}, id).Error
}
