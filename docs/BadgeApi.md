# \BadgeApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectPolicyViolationsBadge**](BadgeApi.md#GetProjectPolicyViolationsBadge) | **Get** /v1/badge/violations/project/{name}/{version} | Returns a policy violations badge for a specific project
[**GetProjectPolicyViolationsBadge1**](BadgeApi.md#GetProjectPolicyViolationsBadge1) | **Get** /v1/badge/violations/project/{uuid} | Returns a policy violations badge for a specific project
[**GetProjectVulnerabilitiesBadge**](BadgeApi.md#GetProjectVulnerabilitiesBadge) | **Get** /v1/badge/vulns/project/{uuid} | Returns current metrics for a specific project
[**GetProjectVulnerabilitiesBadge1**](BadgeApi.md#GetProjectVulnerabilitiesBadge1) | **Get** /v1/badge/vulns/project/{name}/{version} | Returns current metrics for a specific project


# **GetProjectPolicyViolationsBadge**
> string GetProjectPolicyViolationsBadge(ctx, name, version)
Returns a policy violations badge for a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the project to query on | 
  **version** | **string**| The version of the project to query on | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectPolicyViolationsBadge1**
> string GetProjectPolicyViolationsBadge1(ctx, uuid)
Returns a policy violations badge for a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to retrieve a badge for | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectVulnerabilitiesBadge**
> ProjectMetrics GetProjectVulnerabilitiesBadge(ctx, uuid)
Returns current metrics for a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to retrieve metrics for | 

### Return type

[**ProjectMetrics**](ProjectMetrics.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectVulnerabilitiesBadge1**
> ProjectMetrics GetProjectVulnerabilitiesBadge1(ctx, name, version)
Returns current metrics for a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the project to query on | 
  **version** | **string**| The version of the project to query on | 

### Return type

[**ProjectMetrics**](ProjectMetrics.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

