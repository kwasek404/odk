/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Convert instance to template command
type ConvertInstanceToTemplateCommand struct {

	// Template name
	TemplateName string `json:"TemplateName"`

	// Template descriptions
	TemplateDescriptions []TemplateDescription `json:"TemplateDescriptions"`

	// Template icon
	TemplateIcon *TemplateIcon `json:"TemplateIcon,omitempty"`

	// Template version
	TemplateVersion string `json:"TemplateVersion"`

	// Template system category id
	TemplateSystemCategoryId int32 `json:"TemplateSystemCategoryId"`

	// Template windows type id
	TemplateWindowsTypeId int32 `json:"TemplateWindowsTypeId,omitempty"`

	// Template default type id
	TemplateDefaultTypeId int32 `json:"TemplateDefaultTypeId"`

	// Template minimum type id
	TemplateMinimumTypeId int32 `json:"TemplateMinimumTypeId"`

	// Template with initialization
	TemplateWithInitialization bool `json:"TemplateWithInitialization"`

	// Account 'tech-support' password. Required for templates with initialization
	TechSupportPassword string `json:"TechSupportPassword,omitempty"`
}
