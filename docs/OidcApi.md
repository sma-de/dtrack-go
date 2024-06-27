# \OidcApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMapping2**](OidcApi.md#AddMapping2) | **Put** /v1/oidc/mapping | Adds a mapping
[**CreateGroup**](OidcApi.md#CreateGroup) | **Put** /v1/oidc/group | Creates group
[**DeleteGroup**](OidcApi.md#DeleteGroup) | **Delete** /v1/oidc/group/{uuid} | Deletes a group
[**DeleteMapping2**](OidcApi.md#DeleteMapping2) | **Delete** /v1/oidc/group/{groupUuid}/team/{teamUuid}/mapping | Deletes a mapping
[**DeleteMappingByUuid**](OidcApi.md#DeleteMappingByUuid) | **Delete** /v1/oidc/mapping/{uuid} | Deletes a mapping
[**IsAvailable**](OidcApi.md#IsAvailable) | **Get** /v1/oidc/available | Indicates if OpenID Connect is available for this application
[**RetrieveGroups**](OidcApi.md#RetrieveGroups) | **Get** /v1/oidc/group | Returns a list of all groups
[**RetrieveTeamsMappedToGroup**](OidcApi.md#RetrieveTeamsMappedToGroup) | **Get** /v1/oidc/group/{uuid}/team | Returns a list of teams associated with the specified group
[**UpdateGroup**](OidcApi.md#UpdateGroup) | **Post** /v1/oidc/group | Updates group


# **AddMapping2**
> MappedOidcGroup AddMapping2(ctx, optional)
Adds a mapping



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OidcApiAddMapping2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OidcApiAddMapping2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MappedOidcGroupRequest**](MappedOidcGroupRequest.md)|  | 

### Return type

[**MappedOidcGroup**](MappedOidcGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> OidcGroup CreateGroup(ctx, optional)
Creates group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OidcApiCreateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OidcApiCreateGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OidcGroup**](OidcGroup.md)|  | 

### Return type

[**OidcGroup**](OidcGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup(ctx, uuid)
Deletes a group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the group to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMapping2**
> DeleteMapping2(ctx, groupUuid, teamUuid)
Deletes a mapping



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupUuid** | **string**| The UUID of the group to delete a mapping for | 
  **teamUuid** | **string**| The UUID of the team to delete a mapping for | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMappingByUuid**
> DeleteMappingByUuid(ctx, uuid)
Deletes a mapping



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the mapping to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsAvailable**
> bool IsAvailable(ctx, )
Indicates if OpenID Connect is available for this application



### Required Parameters
This endpoint does not need any parameter.

### Return type

**bool**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveGroups**
> []OidcGroup RetrieveGroups(ctx, )
Returns a list of all groups



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OidcGroup**](OidcGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveTeamsMappedToGroup**
> []Team RetrieveTeamsMappedToGroup(ctx, uuid)
Returns a list of teams associated with the specified group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the mapping to retrieve the team for | 

### Return type

[**[]Team**](Team.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> OidcGroup UpdateGroup(ctx, optional)
Updates group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OidcApiUpdateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OidcApiUpdateGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OidcGroup**](OidcGroup.md)|  | 

### Return type

[**OidcGroup**](OidcGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

