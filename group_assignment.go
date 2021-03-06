/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Group assignment
type GroupAssignment struct {

	// Id of an instance
	InstanceId int32 `json:"InstanceId,omitempty"`

	// Id of ip
	IpId int32 `json:"IpId,omitempty"`

	// Instance ip v4 address attached to group
	IpV4 string `json:"IpV4,omitempty"`

	// Instance ip v6 address attached to group
	IpV6 string `json:"IpV6,omitempty"`
}
