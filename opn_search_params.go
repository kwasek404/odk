/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Search params for OPNs
type OpnSearchParams struct {

	// Instance id
	InstanceId int32 `json:"InstanceId,omitempty"`

	// Query
	Query string `json:"Query,omitempty"`

	// Page size
	PageSize int32 `json:"PageSize,omitempty"`

	// Page number
	PageNumber int32 `json:"PageNumber,omitempty"`

	// Order by
	OrderBy string `json:"OrderBy,omitempty"`

	// Sort expression
	SortExpression string `json:"SortExpression,omitempty"`

	// Sort direction
	SortDirection string `json:"SortDirection,omitempty"`
}
