package secondary

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
		resourceConfigurationRepository: resourceConfigurationRepository
	}
}

// FIXME return a list of all of the resources so they can be built into the pulumi program as well
// TODO With configurations!
func (service ResourceDBService) fetchAllResources() ([]error, error) {
	return nil, nil
}

func (service ResourceDBService) saveResource() error {
	return nil
}
