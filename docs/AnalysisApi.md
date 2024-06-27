# \AnalysisApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveAnalysis**](AnalysisApi.md#RetrieveAnalysis) | **Get** /v1/analysis | Retrieves an analysis trail
[**UpdateAnalysis**](AnalysisApi.md#UpdateAnalysis) | **Put** /v1/analysis | Records an analysis decision


# **RetrieveAnalysis**
> Analysis RetrieveAnalysis(ctx, component, vulnerability, optional)
Retrieves an analysis trail



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **component** | **string**| The UUID of the component | 
  **vulnerability** | **string**| The UUID of the vulnerability | 
 **optional** | ***AnalysisApiRetrieveAnalysisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiRetrieveAnalysisOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **project** | **optional.String**| The UUID of the project | 

### Return type

[**Analysis**](Analysis.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAnalysis**
> Analysis UpdateAnalysis(ctx, optional)
Records an analysis decision



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AnalysisApiUpdateAnalysisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiUpdateAnalysisOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AnalysisRequest**](AnalysisRequest.md)|  | 

### Return type

[**Analysis**](Analysis.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

