# \AclApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMapping**](AclApi.md#AddMapping) | **Put** /v1/acl/mapping | Adds an ACL mapping
[**DeleteMapping**](AclApi.md#DeleteMapping) | **Delete** /v1/acl/mapping/team/{teamUuid}/project/{projectUuid} | Removes an ACL mapping
[**RetrieveProjects**](AclApi.md#RetrieveProjects) | **Get** /v1/acl/team/{uuid} | Returns the projects assigned to the specified team


# **AddMapping**
> AclMappingRequest AddMapping(ctx, optional)
Adds an ACL mapping



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AclApiAddMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclApiAddMappingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AclMappingRequest**](AclMappingRequest.md)|  | 

### Return type

[**AclMappingRequest**](AclMappingRequest.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMapping**
> DeleteMapping(ctx, teamUuid, projectUuid)
Removes an ACL mapping



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamUuid** | **string**| The UUID of the team to delete the mapping for | 
  **projectUuid** | **string**| The UUID of the project to delete the mapping for | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveProjects**
> []string RetrieveProjects(ctx, uuid, optional)
Returns the projects assigned to the specified team



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the team to retrieve mappings for | 
 **optional** | ***AclApiRetrieveProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclApiRetrieveProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeInactive** | **optional.Bool**| Optionally excludes inactive projects from being returned | 
 **onlyRoot** | **optional.Bool**| Optionally excludes children projects from being returned | 

### Return type

**[]string**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

