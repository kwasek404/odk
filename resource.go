/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Resource with HATEOAS links
type Resource struct {

	Links []Link `json:"Links,omitempty"`

	// Resurce id
	Id int32 `json:"Id,omitempty"`
}
