/*
 * Dependency-Track API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.10.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Component struct {
	Author string `json:"author,omitempty"`
	Publisher string `json:"publisher,omitempty"`
	Supplier *OrganizationalEntity `json:"supplier,omitempty"`
	Group string `json:"group,omitempty"`
	Name string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Classifier string `json:"classifier,omitempty"`
	Filename string `json:"filename,omitempty"`
	Extension string `json:"extension,omitempty"`
	Md5 string `json:"md5,omitempty"`
	Sha1 string `json:"sha1,omitempty"`
	Sha256 string `json:"sha256,omitempty"`
	Sha384 string `json:"sha384,omitempty"`
	Sha512 string `json:"sha512,omitempty"`
	Sha3256 string `json:"sha3_256,omitempty"`
	Sha3384 string `json:"sha3_384,omitempty"`
	Sha3512 string `json:"sha3_512,omitempty"`
	Blake2b256 string `json:"blake2b_256,omitempty"`
	Blake2b384 string `json:"blake2b_384,omitempty"`
	Blake2b512 string `json:"blake2b_512,omitempty"`
	Blake3 string `json:"blake3,omitempty"`
	Cpe string `json:"cpe,omitempty"`
	Purl *PackageUrl `json:"purl,omitempty"`
	PurlCoordinates *PackageUrl `json:"purlCoordinates,omitempty"`
	SwidTagId string `json:"swidTagId,omitempty"`
	Description string `json:"description,omitempty"`
	Copyright string `json:"copyright,omitempty"`
	License string `json:"license,omitempty"`
	LicenseExpression string `json:"licenseExpression,omitempty"`
	LicenseUrl string `json:"licenseUrl,omitempty"`
	ResolvedLicense *License `json:"resolvedLicense,omitempty"`
	DirectDependencies string `json:"directDependencies,omitempty"`
	ExternalReferences []ExternalReference `json:"externalReferences,omitempty"`
	Parent *Component `json:"parent,omitempty"`
	Children []Component `json:"children,omitempty"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities,omitempty"`
	Project *Project `json:"project"`
	LastInheritedRiskScore float64 `json:"lastInheritedRiskScore,omitempty"`
	Notes string `json:"notes,omitempty"`
	Uuid string `json:"uuid"`
	Metrics *DependencyMetrics `json:"metrics,omitempty"`
	RepositoryMeta *RepositoryMetaComponent `json:"repositoryMeta,omitempty"`
	BomRef string `json:"bomRef,omitempty"`
	UsedBy int32 `json:"usedBy,omitempty"`
	DependencyGraph []string `json:"dependencyGraph,omitempty"`
	ExpandDependencyGraph bool `json:"expandDependencyGraph,omitempty"`
	IsInternal bool `json:"isInternal,omitempty"`
}
