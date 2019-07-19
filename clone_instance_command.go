/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Clone instance command
type CloneInstanceCommand struct {

	// Clone name
	CloneName string `json:"CloneName"`

	// Clone type
	CloneType int32 `json:"CloneType"`

	// Subregion Id
	SubregionId int32 `json:"SubregionId,omitempty"`

	// Instance should be power on after clone
	PowerOn bool `json:"PowerOn,omitempty"`
}