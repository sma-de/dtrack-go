# \ViolationApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetViolations**](ViolationApi.md#GetViolations) | **Get** /v1/violation | Returns a list of all policy violations for the entire portfolio
[**GetViolationsByComponent**](ViolationApi.md#GetViolationsByComponent) | **Get** /v1/violation/component/{uuid} | Returns a list of all policy violations for a specific component
[**GetViolationsByProject**](ViolationApi.md#GetViolationsByProject) | **Get** /v1/violation/project/{uuid} | Returns a list of all policy violations for a specific project


# **GetViolations**
> []PolicyViolation GetViolations(ctx, optional)
Returns a list of all policy violations for the entire portfolio



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ViolationApiGetViolationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViolationApiGetViolationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **suppressed** | **optional.Bool**| Optionally includes suppressed violations | 

### Return type

[**[]PolicyViolation**](PolicyViolation.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetViolationsByComponent**
> []PolicyViolation GetViolationsByComponent(ctx, uuid, optional)
Returns a list of all policy violations for a specific component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 
 **optional** | ***ViolationApiGetViolationsByComponentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViolationApiGetViolationsByComponentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressed** | **optional.Bool**| Optionally includes suppressed violations | 

### Return type

[**[]PolicyViolation**](PolicyViolation.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetViolationsByProject**
> []PolicyViolation GetViolationsByProject(ctx, uuid, optional)
Returns a list of all policy violations for a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 
 **optional** | ***ViolationApiGetViolationsByProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViolationApiGetViolationsByProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressed** | **optional.Bool**| Optionally includes suppressed violations | 

### Return type

[**[]PolicyViolation**](PolicyViolation.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

