/*
 * Dependency-Track API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.10.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ConfigProperty struct {
	GroupName string `json:"groupName,omitempty"`
	PropertyName string `json:"propertyName,omitempty"`
	PropertyValue string `json:"propertyValue,omitempty"`
	PropertyType string `json:"propertyType"`
	Description string `json:"description,omitempty"`
}
