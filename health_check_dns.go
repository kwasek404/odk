/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

import (
	"time"
)

// Dns health check
type HealthCheckDns struct {

	// Timeout
	Timeout int32 `json:"Timeout,omitempty"`

	// Record type
	RecordType *DictionaryItem `json:"RecordType,omitempty"`

	// Query domain
	QueryDomain string `json:"QueryDomain,omitempty"`

	// Expected DNS record value
	ExpectedResponseValue string `json:"ExpectedResponseValue,omitempty"`

	// Recursive query
	Recurse bool `json:"Recurse,omitempty"`

	// Notification types enabled for a health check
	NotificationTypes []DictionaryItem `json:"NotificationTypes,omitempty"`

	// Event types with enabled notification
	NotificationEventTypes []DictionaryItem `json:"NotificationEventTypes,omitempty"`

	// Time when notification is sent
	NotificationTime *DictionaryItem `json:"NotificationTime,omitempty"`

	// Use random substitute locations in case of location breakdown
	LocationsFailoverEnabled bool `json:"LocationsFailoverEnabled,omitempty"`

	// How many (%) locations have to report an error to consider it a breakdown
	ErrorTolerance int32 `json:"ErrorTolerance,omitempty"`

	// Id
	Id int32 `json:"Id,omitempty"`

	// Interval
	Interval int32 `json:"Interval,omitempty"`

	// Name
	Name string `json:"Name,omitempty"`

	// Address
	Address string `json:"Address,omitempty"`

	// Type
	ServiceType *DictionaryItem `json:"ServiceType,omitempty"`

	// State
	State *DictionaryItem `json:"State,omitempty"`

	// Details url
	DetailsLocation string `json:"DetailsLocation,omitempty"`

	// Is paused
	Paused bool `json:"Paused,omitempty"`

	// Is suspended
	Suspended bool `json:"Suspended,omitempty"`

	// Last invalid check
	LastInvalidCheck time.Time `json:"LastInvalidCheck,omitempty"`

	// Last valid check
	LastValidCheck time.Time `json:"LastValidCheck,omitempty"`
}
