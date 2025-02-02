/*
 * Dependency-Track API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.10.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AnalysisRequest struct {
	Project string `json:"project,omitempty"`
	Component string `json:"component"`
	Vulnerability string `json:"vulnerability"`
	AnalysisState string `json:"analysisState,omitempty"`
	AnalysisJustification string `json:"analysisJustification,omitempty"`
	AnalysisResponse string `json:"analysisResponse,omitempty"`
	AnalysisDetails string `json:"analysisDetails,omitempty"`
	Comment string `json:"comment,omitempty"`
	Suppressed bool `json:"suppressed,omitempty"`
}
