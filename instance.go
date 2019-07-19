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

// Instance
type Instance struct {

	// ID
	Id int32 `json:"Id,omitempty"`

	// Name
	Name string `json:"Name,omitempty"`

	// Creation date
	CreationDate time.Time `json:"CreationDate,omitempty"`

	// User who created the instance
	CreationUser *UserResource `json:"CreationUser,omitempty"`

	// If the instance is locked by a running operation
	IsLocked bool `json:"IsLocked,omitempty"`

	// Locking date
	LockingDate time.Time `json:"LockingDate,omitempty"`

	// Template from which the instance was created
	Template *BaseResource `json:"Template,omitempty"`

	// Subregion that is running the instance
	Subregion *BaseResource `json:"Subregion,omitempty"`

	// Instance type. Defines the configuration of CPU and RAM
	Type_ *DictionaryItem `json:"Type,omitempty"`

	// Status
	Status *DictionaryItem `json:"Status,omitempty"`

	// Operating system category installed on the instance
	SystemCategory *DictionaryItem `json:"SystemCategory,omitempty"`

	// Autoscaling type
	AutoscalingType *DictionaryItem `json:"AutoscalingType,omitempty"`

	// VMware Tools status
	VmWareToolsStatus *DictionaryItem `json:"VmWareToolsStatus,omitempty"`

	// Monitoring status
	MonitStatus *DictionaryItem `json:"MonitStatus,omitempty"`

	// Template type eg marketplace, oci instance
	TemplateType *DictionaryItem `json:"TemplateType,omitempty"`

	// IP address
	IpAddress string `json:"IpAddress,omitempty"`

	// DNS address
	DnsAddress string `json:"DnsAddress,omitempty"`

	// Total disks capacity in GB
	TotalDisksCapacity int32 `json:"TotalDisksCapacity,omitempty"`

	// Payment type
	PaymentType *DictionaryItem `json:"PaymentType,omitempty"`

	// Health check
	HealthCheck *BaseResource `json:"HealthCheck,omitempty"`

	// SCSI controller type
	ScsiControllerType *DictionaryItem `json:"ScsiControllerType,omitempty"`

	// Is freemium
	IsFreemium bool `json:"IsFreemium,omitempty"`

	// Number of CPUs
	CpuNumber int32 `json:"CpuNumber,omitempty"`

	// Memory in MB
	RamMb int32 `json:"RamMb,omitempty"`

	// Support type
	SupportType *Software `json:"SupportType,omitempty"`
}
