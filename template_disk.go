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

// Template disk
type TemplateDisk struct {

	// Id
	Id int32 `json:"Id,omitempty"`

	// Name
	Name string `json:"Name,omitempty"`

	// Space capacity in GB
	SpaceCapacity int32 `json:"SpaceCapacity,omitempty"`

	// Tier
	Tier *DictionaryItem `json:"Tier,omitempty"`

	// Creation date
	CreationDate time.Time `json:"CreationDate,omitempty"`

	// Controller
	Controller int32 `json:"Controller,omitempty"`

	// Slot
	Slot int32 `json:"Slot,omitempty"`

	// If is system disk
	IsSystemDisk bool `json:"IsSystemDisk,omitempty"`
}
