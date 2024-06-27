/*
 * Dependency-Track API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.10.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Project struct {
	Author string `json:"author,omitempty"`
	Publisher string `json:"publisher,omitempty"`
	Manufacturer *OrganizationalEntity `json:"manufacturer,omitempty"`
	Supplier *OrganizationalEntity `json:"supplier,omitempty"`
	Group string `json:"group,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version string `json:"version,omitempty"`
	Classifier string `json:"classifier,omitempty"`
	Cpe string `json:"cpe,omitempty"`
	Purl *PackageUrl `json:"purl,omitempty"`
	SwidTagId string `json:"swidTagId,omitempty"`
	DirectDependencies string `json:"directDependencies,omitempty"`
	Uuid string `json:"uuid"`
	Parent *Project `json:"parent,omitempty"`
	Children []Project `json:"children,omitempty"`
	Properties []ProjectProperty `json:"properties,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
	LastBomImport time.Time `json:"lastBomImport,omitempty"`
	LastBomImportFormat string `json:"lastBomImportFormat,omitempty"`
	LastInheritedRiskScore float64 `json:"lastInheritedRiskScore,omitempty"`
	Active bool `json:"active,omitempty"`
	ExternalReferences []ExternalReference `json:"externalReferences,omitempty"`
	Metadata *ProjectMetadata `json:"metadata,omitempty"`
	Versions []ProjectVersion `json:"versions,omitempty"`
	Metrics *ProjectMetrics `json:"metrics,omitempty"`
}
