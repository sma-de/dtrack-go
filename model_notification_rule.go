/*
 * Dependency-Track API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.10.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NotificationRule struct {
	Name string `json:"name"`
	Enabled bool `json:"enabled,omitempty"`
	NotifyChildren bool `json:"notifyChildren,omitempty"`
	LogSuccessfulPublish bool `json:"logSuccessfulPublish,omitempty"`
	Scope string `json:"scope"`
	NotificationLevel string `json:"notificationLevel,omitempty"`
	Projects []Project `json:"projects,omitempty"`
	Teams []Team `json:"teams,omitempty"`
	NotifyOn []string `json:"notifyOn,omitempty"`
	Message string `json:"message,omitempty"`
	Publisher *NotificationPublisher `json:"publisher,omitempty"`
	PublisherConfig string `json:"publisherConfig,omitempty"`
	Uuid string `json:"uuid"`
}
