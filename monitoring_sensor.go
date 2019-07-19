/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Monitoring sensor
type MonitoringSensor struct {

	// Identifier
	Id int32 `json:"Id,omitempty"`

	// Country
	Country string `json:"Country,omitempty"`

	// City
	City string `json:"City,omitempty"`

	// IPAddress
	IPAddress string `json:"IPAddress,omitempty"`

	// Name
	Name string `json:"Name,omitempty"`

	// Is available
	IsAvailable bool `json:"IsAvailable,omitempty"`
}
