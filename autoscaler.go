/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Instance autoscaler settings
type Autoscaler struct {

	// Instance autoscaling mode
	AutoscalingMode *DictionaryItem `json:"AutoscalingMode,omitempty"`

	// Is hotplug enabled
	HotPlugEnabled bool `json:"HotPlugEnabled,omitempty"`

	// Instance autoscaling parameter
	AutoscalingParameter *InstanceAutoscalingParameter `json:"AutoscalingParameter,omitempty"`
}
