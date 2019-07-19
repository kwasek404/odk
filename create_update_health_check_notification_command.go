/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Create/update health check notification command
type CreateUpdateHealthCheckNotificationCommand struct {

	// Address
	Address string `json:"Address"`

	// Health check notification type (Dictionary 178)
	NotificationTypeId int32 `json:"NotificationTypeId"`
}
