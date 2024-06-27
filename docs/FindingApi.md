# \FindingApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeProject**](FindingApi.md#AnalyzeProject) | **Post** /v1/finding/project/{uuid}/analyze | Triggers Vulnerability Analysis on a specific project
[**ExportFindingsByProject**](FindingApi.md#ExportFindingsByProject) | **Get** /v1/finding/project/{uuid}/export | Returns the findings for the specified project as FPF
[**GetFindingsByProject**](FindingApi.md#GetFindingsByProject) | **Get** /v1/finding/project/{uuid} | Returns a list of all findings for a specific project


# **AnalyzeProject**
> Project AnalyzeProject(ctx, uuid)
Triggers Vulnerability Analysis on a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to analyze | 

### Return type

[**Project**](Project.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportFindingsByProject**
> ExportFindingsByProject(ctx, uuid)
Returns the findings for the specified project as FPF



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFindingsByProject**
> []Finding GetFindingsByProject(ctx, uuid, optional)
Returns a list of all findings for a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 
 **optional** | ***FindingApiGetFindingsByProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindingApiGetFindingsByProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressed** | **optional.Bool**| Optionally includes suppressed findings | 
 **source** | **optional.String**| Optionally limit findings to specific sources of vulnerability intelligence | 

### Return type

[**[]Finding**](Finding.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

