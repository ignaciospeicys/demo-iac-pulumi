package dto

type ObjectStorageCreateRequest struct {
	Name       string `json:"name"`
	Versioning bool   `json:"enable_versioning"`
}
