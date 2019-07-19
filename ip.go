/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Public IP address
type Ip struct {

	// ID
	Id int32 `json:"Id,omitempty"`

	// IP address version 4
	Address string `json:"Address,omitempty"`

	// IP address version 6
	AddressV6 string `json:"AddressV6,omitempty"`

	// Gateway address
	Gateway string `json:"Gateway,omitempty"`

	// Netmask
	NetMask string `json:"NetMask,omitempty"`

	// Instance
	Instance *BaseResource `json:"Instance,omitempty"`

	// The MAC address of the network card associated with that IP address
	MacAddress string `json:"MacAddress,omitempty"`

	// Network card number
	InterfaceId int32 `json:"InterfaceId,omitempty"`

	// DNS prefix
	DnsPrefix string `json:"DnsPrefix,omitempty"`

	// Dhcp branch address
	DhcpBranch string `json:"DhcpBranch,omitempty"`

	// Subregion
	Subregion *BaseResource `json:"Subregion,omitempty"`

	// Type
	Type_ *DictionaryItem `json:"Type,omitempty"`

	// Account that is owner IP address
	OwnerAccount *BaseResource `json:"OwnerAccount,omitempty"`

	// Owner comment
	Comment string `json:"Comment,omitempty"`

	// Reverse DNS IP version 4
	RevDns string `json:"RevDns,omitempty"`

	// Reverse DNS IP version 6
	RevDnsV6 string `json:"RevDnsV6,omitempty"`

	// User who created the IP
	CreationUser *UserResource `json:"CreationUser,omitempty"`
}