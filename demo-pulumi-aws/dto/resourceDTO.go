package dto

type ResourceDTO struct {
	ResourceName   string
	ResourceType   string
	StackName      string
	Status         string
	Configurations []ConfigurationDTO
}

type ConfigurationDTO struct {
	ConfigKey   string
	ConfigValue string
}
