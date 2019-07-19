/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// User
type UserResource struct {

	// User id
	Id int32 `json:"Id,omitempty"`

	// User login
	Login string `json:"Login,omitempty"`

	// First name
	FirstName string `json:"FirstName,omitempty"`

	// Last name
	LastName string `json:"LastName,omitempty"`
}
