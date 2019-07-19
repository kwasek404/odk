/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Create/update http health check command
type CreateUpdateHealthCheckHttpCommand struct {

	// Health check http method type (Dictionary 166)
	HttpMethodId int32 `json:"HttpMethodId"`

	// The content has to match the expression (GET and POST methods only)
	ContentRegularExpression string `json:"ContentRegularExpression,omitempty"`

	// The content cannot match the expression (GET and POST methods only)
	ContentNegativeRegularExpression string `json:"ContentNegativeRegularExpression,omitempty"`

	// Port
	Port int32 `json:"Port"`

	// Time the server has to complete the request in [ms]
	Timeout int32 `json:"Timeout"`

	// Content
	Content string `json:"Content,omitempty"`

	// Content type
	ContentType string `json:"ContentType,omitempty"`

	// How many (%) locations have to report an error to consider it a breakdown
	ErrorTolerance int32 `json:"ErrorTolerance"`

	// Health check name
	Name string `json:"Name"`

	// URL or IP address of the monitored service
	Address string `json:"Address"`

	// Time interval between health checks, in seconds
	Interval int32 `json:"Interval"`

	// Is paused
	Paused bool `json:"Paused"`

	// Use random substitute locations in case of location breakdown
	LocationsFailoverEnabled bool `json:"LocationsFailoverEnabled"`

	// Notification types enabled for a health check
	NotificationTypeIds []int32 `json:"NotificationTypeIds,omitempty"`

	// Event types with enabled notification
	NotificationEventTypeIds []int32 `json:"NotificationEventTypeIds,omitempty"`

	// Time when notification is sent
	NotificationTimeId int32 `json:"NotificationTimeId"`
}
