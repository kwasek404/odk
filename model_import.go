/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

import (
	"time"
)

// Import
type ModelImport struct {

	// Id
	Id int32 `json:"Id,omitempty"`

	// Name
	Name string `json:"Name,omitempty"`

	// Creation date
	CreationDate time.Time `json:"CreationDate,omitempty"`

	// Start date
	StartDate time.Time `json:"StartDate,omitempty"`

	// End date
	EndDate time.Time `json:"EndDate,omitempty"`

	// Total size in GB
	Size int32 `json:"Size,omitempty"`

	// Network interfaces count
	NetworkInterfacesCount int32 `json:"NetworkInterfacesCount,omitempty"`

	// Disks count
	DisksCount int32 `json:"DisksCount,omitempty"`

	// Status
	Status *DictionaryItem `json:"Status,omitempty"`

	// User who created the import
	CreationUser *UserResource `json:"CreationUser,omitempty"`
}
