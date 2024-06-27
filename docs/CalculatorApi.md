# \CalculatorApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCvssScores**](CalculatorApi.md#GetCvssScores) | **Get** /v1/calculator/cvss | Returns the CVSS base score, impact sub-score and exploitability sub-score
[**GetOwaspRRScores**](CalculatorApi.md#GetOwaspRRScores) | **Get** /v1/calculator/owasp | Returns the OWASP Risk Rating likelihood score, technical impact score and business impact score


# **GetCvssScores**
> Score GetCvssScores(ctx, vector)
Returns the CVSS base score, impact sub-score and exploitability sub-score



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vector** | **string**| A valid CVSSv2 or CVSSv3 vector | 

### Return type

[**Score**](Score.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOwaspRRScores**
> Score GetOwaspRRScores(ctx, vector)
Returns the OWASP Risk Rating likelihood score, technical impact score and business impact score



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vector** | **string**| A valid OWASP Risk Rating vector | 

### Return type

[**Score**](Score.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

