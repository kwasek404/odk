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

// SSH key
type SshKey struct {

	// Id
	Id int32 `json:"Id,omitempty"`

	// Name
	Name string `json:"Name,omitempty"`

	// Key value trimmed
	Value string `json:"Value,omitempty"`

	// User owning the key
	OwnerUser *UserResource `json:"OwnerUser,omitempty"`

	// Creation date
	CreationDate time.Time `json:"CreationDate,omitempty"`
}
