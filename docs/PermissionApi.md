# \PermissionApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPermissionToTeam**](PermissionApi.md#AddPermissionToTeam) | **Post** /v1/permission/{permission}/team/{uuid} | Adds the permission to the specified team.
[**AddPermissionToUser**](PermissionApi.md#AddPermissionToUser) | **Post** /v1/permission/{permission}/user/{username} | Adds the permission to the specified username.
[**GetAllPermissions**](PermissionApi.md#GetAllPermissions) | **Get** /v1/permission | Returns a list of all permissions
[**RemovePermissionFromTeam**](PermissionApi.md#RemovePermissionFromTeam) | **Delete** /v1/permission/{permission}/team/{uuid} | Removes the permission from the team.
[**RemovePermissionFromUser**](PermissionApi.md#RemovePermissionFromUser) | **Delete** /v1/permission/{permission}/user/{username} | Removes the permission from the user.


# **AddPermissionToTeam**
> Team AddPermissionToTeam(ctx, uuid, permission)
Adds the permission to the specified team.

Requires 'manage users' permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| A valid team uuid | 
  **permission** | **string**| A valid permission | 

### Return type

[**Team**](Team.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPermissionToUser**
> UserPrincipal AddPermissionToUser(ctx, username, permission)
Adds the permission to the specified username.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| A valid username | 
  **permission** | **string**| A valid permission | 

### Return type

[**UserPrincipal**](UserPrincipal.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPermissions**
> []Permission GetAllPermissions(ctx, )
Returns a list of all permissions



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Permission**](Permission.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePermissionFromTeam**
> Team RemovePermissionFromTeam(ctx, uuid, permission)
Removes the permission from the team.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| A valid team uuid | 
  **permission** | **string**| A valid permission | 

### Return type

[**Team**](Team.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePermissionFromUser**
> UserPrincipal RemovePermissionFromUser(ctx, username, permission)
Removes the permission from the user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| A valid username | 
  **permission** | **string**| A valid permission | 

### Return type

[**UserPrincipal**](UserPrincipal.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

