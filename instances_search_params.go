/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Search params for instances
type InstancesSearchParams struct {

	// Template type id eg marketplace, oci instance
	TemplateTypeId int32 `json:"TemplateTypeId,omitempty"`

	// Indicates wether an instance is turned on
	IsTurnedOn bool `json:"IsTurnedOn,omitempty"`

	// Subregion Id
	SubregionId int32 `json:"SubregionId,omitempty"`

	// Type Id
	TypeId int32 `json:"TypeId,omitempty"`

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
