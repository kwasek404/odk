/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Create export command
type CreateExportCommand struct {

	// Name
	Name string `json:"Name"`

	// OCS location
	OcsLocation string `json:"OcsLocation"`

	// OCS Project Id
	OcsProjectId string `json:"OcsProjectId"`
}