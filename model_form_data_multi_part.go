/*
 * Dependency-Track API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.10.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type FormDataMultiPart struct {
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`
	Entity interface{} `json:"entity,omitempty"`
	Headers map[string][]string `json:"headers,omitempty"`
	MediaType *MediaType `json:"mediaType,omitempty"`
	MessageBodyWorkers *MessageBodyWorkers `json:"messageBodyWorkers,omitempty"`
	Parent *MultiPart `json:"parent,omitempty"`
	Providers *Providers `json:"providers,omitempty"`
	BodyParts []BodyPart `json:"bodyParts,omitempty"`
	Fields map[string][]FormDataBodyPart `json:"fields,omitempty"`
	ParameterizedHeaders map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`
}
