/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Group autoscaler settings
type GroupAutoscaler struct {

	// Expansion type
	ExpansionType *DictionaryItem `json:"ExpansionType,omitempty"`

	// Minimum instance number in group
	MinimumInstanceNumber int32 `json:"MinimumInstanceNumber,omitempty"`

	// Minimum instance number in group
	MaximumInstanceNumber int32 `json:"MaximumInstanceNumber,omitempty"`

	// Instances to scaling
	Instances []GroupAutoscalerInstance `json:"Instances,omitempty"`
}
