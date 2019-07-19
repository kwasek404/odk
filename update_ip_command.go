/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Instance IP configuration update command
type UpdateIpCommand struct {

	// IP should be static
	SetStatic bool `json:"SetStatic,omitempty"`

	// IP comment
	Comment string `json:"Comment,omitempty"`

	// Restore default value for revDNS IPv4
	RestoreRevDns bool `json:"RestoreRevDns,omitempty"`

	// Restore default value for revDNS IPv6
	RestoreRevDnsV6 bool `json:"RestoreRevDnsV6,omitempty"`

	// revDNS for IPv4
	RevDns string `json:"RevDns,omitempty"`

	// revDNS for IPv6
	RevDnsV6 string `json:"RevDnsV6,omitempty"`
}
