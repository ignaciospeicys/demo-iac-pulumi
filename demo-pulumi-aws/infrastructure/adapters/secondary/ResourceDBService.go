package secondary

import (
	"demo-pulumi-aws/dto"
	"demo-pulumi-aws/model"
)

// ResourceDBService aggregates operations for resources, configurations, stacks, and resource stacks.
type ResourceDBService struct {
	resourceRepository              *ResourceRepository
	resourceConfigurationRepository *ConfigurationRepository
}

// NewResourceDBService creates a new instance of InfrastructureService.
func NewResourceDBService(
	resourceRepository *ResourceRepository,
	resourceConfigurationRepository *ConfigurationRepository) *ResourceDBService {
	return &ResourceDBService{
		resourceRepository:              resourceRepository,
		resourceConfigurationRepository: resourceConfigurationRepository,
	}
}

func (service *ResourceDBService) FetchAllResources(stack string) ([]dto.ResourceDTO, error) {
	var resources []model.Resource
	var resourceDTOs []dto.ResourceDTO

	// Fetch all resources from the database
	if err := service.resourceRepository.db.Preload("Configurations").Where("stack_name = ?", stack).Find(&resources).Error; err != nil {
		return nil, err
	}

	// Map resources and their configurations to DTOs
	for _, resource := range resources {
		var configDTOs []dto.ConfigurationDTO
		for _, config := range resource.Configurations {
			configDTO := dto.ConfigurationDTO{
				ConfigKey:   config.ConfigKey,
				ConfigValue: config.ConfigValue,
			}
			configDTOs = append(configDTOs, configDTO)
		}

		resourceDTO := dto.ResourceDTO{
			ResourceName:   resource.ResourceName,
			ResourceType:   resource.ResourceType,
			StackName:      resource.StackName,
			Status:         resource.Status,
			Configurations: configDTOs,
		}
		resourceDTOs = append(resourceDTOs, resourceDTO)
	}

	return resourceDTOs, nil
}

func (service *ResourceDBService) SaveResource(resourceDTO dto.ResourceDTO) error {
	// Start a new transaction
	tx := service.resourceRepository.db.Begin()

	// Check for transaction error
	if tx.Error != nil {
		return tx.Error
	}

	// Save the resource
	resource := model.Resource{
		ResourceName: resourceDTO.ResourceName,
		ResourceType: resourceDTO.ResourceType,
		StackName:    resourceDTO.StackName,
		Status:       resourceDTO.Status,
	}

	if err := tx.Create(&resource).Error; err != nil {
		tx.Rollback() // Rollback the transaction on error
		return err
	}

	// Save each configuration
	for _, configDTO := range resourceDTO.Configurations {
		config := model.Configuration{
			ResourceID:  resource.ResourceID, // Set the foreign key to the newly created resource
			ConfigKey:   configDTO.ConfigKey,
			ConfigValue: configDTO.ConfigValue,
		}

		if err := tx.Create(&config).Error; err != nil {
			tx.Rollback() // Rollback the transaction on error
			return err
		}
	}
	// Commit the transaction
	return tx.Commit().Error
}
