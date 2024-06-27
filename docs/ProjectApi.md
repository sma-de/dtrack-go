# \ProjectApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneProject**](ProjectApi.md#CloneProject) | **Put** /v1/project/clone | Clones a project
[**CreateProject**](ProjectApi.md#CreateProject) | **Put** /v1/project | Creates a new project
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /v1/project/{uuid} | Deletes a project
[**GetChildrenProjects**](ProjectApi.md#GetChildrenProjects) | **Get** /v1/project/{uuid}/children | Returns a list of all children for a project
[**GetChildrenProjectsByClassifier**](ProjectApi.md#GetChildrenProjectsByClassifier) | **Get** /v1/project/{uuid}/children/classifier/{classifier} | Returns a list of all children for a project by classifier
[**GetChildrenProjectsByTag**](ProjectApi.md#GetChildrenProjectsByTag) | **Get** /v1/project/{uuid}/children/tag/{tag} | Returns a list of all children for a project by tag
[**GetProject**](ProjectApi.md#GetProject) | **Get** /v1/project/{uuid} | Returns a specific project
[**GetProjectByNameAndVersion**](ProjectApi.md#GetProjectByNameAndVersion) | **Get** /v1/project/lookup | Returns a specific project by its name and version
[**GetProjects**](ProjectApi.md#GetProjects) | **Get** /v1/project | Returns a list of all projects
[**GetProjectsByClassifier**](ProjectApi.md#GetProjectsByClassifier) | **Get** /v1/project/classifier/{classifier} | Returns a list of all projects by classifier
[**GetProjectsByTag**](ProjectApi.md#GetProjectsByTag) | **Get** /v1/project/tag/{tag} | Returns a list of all projects by tag
[**GetProjectsWithoutDescendantsOf**](ProjectApi.md#GetProjectsWithoutDescendantsOf) | **Get** /v1/project/withoutDescendantsOf/{uuid} | Returns a list of all projects without the descendants of the selected project
[**PatchProject**](ProjectApi.md#PatchProject) | **Patch** /v1/project/{uuid} | Partially updates a project
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Post** /v1/project | Updates a project


# **CloneProject**
> Project CloneProject(ctx, optional)
Clones a project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectApiCloneProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCloneProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CloneProjectRequest**](CloneProjectRequest.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> Project CreateProject(ctx, optional)
Creates a new project

If a parent project exists, the UUID of the parent project is required 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectApiCreateProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCreateProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Project**](Project.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(ctx, uuid)
Deletes a project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChildrenProjects**
> []Project GetChildrenProjects(ctx, uuid, optional)
Returns a list of all children for a project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to get the children from | 
 **optional** | ***ProjectApiGetChildrenProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetChildrenProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeInactive** | **optional.Bool**| Optionally excludes inactive projects from being returned | 

### Return type

[**[]Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChildrenProjectsByClassifier**
> []Project GetChildrenProjectsByClassifier(ctx, classifier, uuid, optional)
Returns a list of all children for a project by classifier



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **classifier** | **string**| The classifier to query on | 
  **uuid** | **string**| The UUID of the project to get the children from | 
 **optional** | ***ProjectApiGetChildrenProjectsByClassifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetChildrenProjectsByClassifierOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludeInactive** | **optional.Bool**| Optionally excludes inactive projects from being returned | 

### Return type

[**[]Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChildrenProjectsByTag**
> []Project GetChildrenProjectsByTag(ctx, tag, uuid, optional)
Returns a list of all children for a project by tag



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**| The tag to query on | 
  **uuid** | **string**| The UUID of the project to get the children from | 
 **optional** | ***ProjectApiGetChildrenProjectsByTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetChildrenProjectsByTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludeInactive** | **optional.Bool**| Optionally excludes inactive projects from being returned | 

### Return type

[**[]Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> Project GetProject(ctx, uuid)
Returns a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to retrieve | 

### Return type

[**Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectByNameAndVersion**
> Project GetProjectByNameAndVersion(ctx, name, version)
Returns a specific project by its name and version



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the project to query on | 
  **version** | **string**| The version of the project to query on | 

### Return type

[**Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjects**
> []Project GetProjects(ctx, optional)
Returns a list of all projects



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectApiGetProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The optional name of the project to query on | 
 **excludeInactive** | **optional.Bool**| Optionally excludes inactive projects from being returned | 
 **onlyRoot** | **optional.Bool**| Optionally excludes children projects from being returned | 

### Return type

[**[]Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectsByClassifier**
> []Project GetProjectsByClassifier(ctx, classifier, optional)
Returns a list of all projects by classifier



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **classifier** | **string**| The classifier to query on | 
 **optional** | ***ProjectApiGetProjectsByClassifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetProjectsByClassifierOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeInactive** | **optional.Bool**| Optionally excludes inactive projects from being returned | 
 **onlyRoot** | **optional.Bool**| Optionally excludes children projects from being returned | 

### Return type

[**[]Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectsByTag**
> []Project GetProjectsByTag(ctx, tag, optional)
Returns a list of all projects by tag



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**| The tag to query on | 
 **optional** | ***ProjectApiGetProjectsByTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetProjectsByTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeInactive** | **optional.Bool**| Optionally excludes inactive projects from being returned | 
 **onlyRoot** | **optional.Bool**| Optionally excludes children projects from being returned | 

### Return type

[**[]Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectsWithoutDescendantsOf**
> []Project GetProjectsWithoutDescendantsOf(ctx, uuid, optional)
Returns a list of all projects without the descendants of the selected project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project which descendants will be excluded | 
 **optional** | ***ProjectApiGetProjectsWithoutDescendantsOfOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetProjectsWithoutDescendantsOfOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The optional name of the project to query on | 
 **excludeInactive** | **optional.Bool**| Optionally excludes inactive projects from being returned | 

### Return type

[**[]Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchProject**
> Project PatchProject(ctx, uuid, optional)
Partially updates a project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to modify | 
 **optional** | ***ProjectApiPatchProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiPatchProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Project**](Project.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProject**
> Project UpdateProject(ctx, optional)
Updates a project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectApiUpdateProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiUpdateProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Project**](Project.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

