/*
 * Dependency-Track API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.10.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VexSubmitRequest struct {
	Project string `json:"project"`
	ProjectName string `json:"projectName,omitempty"`
	ProjectVersion string `json:"projectVersion,omitempty"`
	Vex string `json:"vex"`
}
