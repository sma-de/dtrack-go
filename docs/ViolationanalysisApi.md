# \ViolationanalysisApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveAnalysis1**](ViolationanalysisApi.md#RetrieveAnalysis1) | **Get** /v1/violation/analysis | Retrieves a violation analysis trail
[**UpdateAnalysis1**](ViolationanalysisApi.md#UpdateAnalysis1) | **Put** /v1/violation/analysis | Records a violation analysis decision


# **RetrieveAnalysis1**
> ViolationAnalysis RetrieveAnalysis1(ctx, component, policyViolation)
Retrieves a violation analysis trail



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **component** | **string**| The UUID of the component | 
  **policyViolation** | **string**| The UUID of the policy violation | 

### Return type

[**ViolationAnalysis**](ViolationAnalysis.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAnalysis1**
> ViolationAnalysis UpdateAnalysis1(ctx, optional)
Records a violation analysis decision



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ViolationanalysisApiUpdateAnalysis1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViolationanalysisApiUpdateAnalysis1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ViolationAnalysisRequest**](ViolationAnalysisRequest.md)|  | 

### Return type

[**ViolationAnalysis**](ViolationAnalysis.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

