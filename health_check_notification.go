/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Heatlth check notification
type HealthCheckNotification struct {

	// Id
	Id int32 `json:"Id,omitempty"`

	// Address
	Address string `json:"Address,omitempty"`

	// Notification type
	NotificationType *DictionaryItem `json:"NotificationType,omitempty"`
}
