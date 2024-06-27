# \PolicyApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProjectToPolicy**](PolicyApi.md#AddProjectToPolicy) | **Post** /v1/policy/{policyUuid}/project/{projectUuid} | Adds a project to a policy
[**AddTagToPolicy**](PolicyApi.md#AddTagToPolicy) | **Post** /v1/policy/{policyUuid}/tag/{tagName} | Adds a tag to a policy
[**CreatePolicy**](PolicyApi.md#CreatePolicy) | **Put** /v1/policy | Creates a new policy
[**DeletePolicy**](PolicyApi.md#DeletePolicy) | **Delete** /v1/policy/{uuid} | Deletes a policy
[**GetPolicies**](PolicyApi.md#GetPolicies) | **Get** /v1/policy | Returns a list of all policies
[**GetPolicy**](PolicyApi.md#GetPolicy) | **Get** /v1/policy/{uuid} | Returns a specific policy
[**RemoveProjectFromPolicy**](PolicyApi.md#RemoveProjectFromPolicy) | **Delete** /v1/policy/{policyUuid}/project/{projectUuid} | Removes a project from a policy
[**RemoveTagFromPolicy**](PolicyApi.md#RemoveTagFromPolicy) | **Delete** /v1/policy/{policyUuid}/tag/{tagName} | Removes a tag from a policy
[**UpdatePolicy**](PolicyApi.md#UpdatePolicy) | **Post** /v1/policy | Updates a policy


# **AddProjectToPolicy**
> Policy AddProjectToPolicy(ctx, policyUuid, projectUuid)
Adds a project to a policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyUuid** | **string**| The UUID of the policy to add a project to | 
  **projectUuid** | **string**| The UUID of the project to add to the rule | 

### Return type

[**Policy**](Policy.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTagToPolicy**
> Policy AddTagToPolicy(ctx, policyUuid, tagName)
Adds a tag to a policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyUuid** | **string**| The UUID of the policy to add a project to | 
  **tagName** | **string**| The name of the tag to add to the rule | 

### Return type

[**Policy**](Policy.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicy**
> Policy CreatePolicy(ctx, optional)
Creates a new policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiCreatePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiCreatePolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Policy**](Policy.md)|  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicy**
> DeletePolicy(ctx, uuid)
Deletes a policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the policy to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicies**
> []Policy GetPolicies(ctx, )
Returns a list of all policies



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Policy**](Policy.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicy**
> Policy GetPolicy(ctx, uuid)
Returns a specific policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the policy to retrieve | 

### Return type

[**Policy**](Policy.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveProjectFromPolicy**
> Policy RemoveProjectFromPolicy(ctx, policyUuid, projectUuid)
Removes a project from a policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyUuid** | **string**| The UUID of the policy to remove the project from | 
  **projectUuid** | **string**| The UUID of the project to remove from the policy | 

### Return type

[**Policy**](Policy.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTagFromPolicy**
> Policy RemoveTagFromPolicy(ctx, policyUuid, tagName)
Removes a tag from a policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyUuid** | **string**| The UUID of the policy to remove the tag from | 
  **tagName** | **string**| The name of the tag to remove from the policy | 

### Return type

[**Policy**](Policy.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicy**
> Policy UpdatePolicy(ctx, optional)
Updates a policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiUpdatePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiUpdatePolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Policy**](Policy.md)|  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

