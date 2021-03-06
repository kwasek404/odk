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

// Ticket representing operation on service eg. instance, disk
type Ticket struct {

	// ID
	Id int64 `json:"Id,omitempty"`

	// Creation date
	CreationDate time.Time `json:"CreationDate,omitempty"`

	// User who created the ticket
	CreationUser *UserResource `json:"CreationUser,omitempty"`

	// End date of operation
	EndDate time.Time `json:"EndDate,omitempty"`

	// Status
	Status *DictionaryItem `json:"Status,omitempty"`

	// Operation type
	OperationType *DictionaryItem `json:"OperationType,omitempty"`

	// Object identifier on which the operation is performed
	ObjectId int32 `json:"ObjectId,omitempty"`

	// Object type
	ObjectType *DictionaryItem `json:"ObjectType,omitempty"`

	// Object name
	ObjectName string `json:"ObjectName,omitempty"`

	// Operation progress
	Progress int32 `json:"Progress,omitempty"`
}
