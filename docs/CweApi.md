# \CweApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCwe**](CweApi.md#GetCwe) | **Get** /v1/cwe/{cweId} | Returns a specific CWE
[**GetCwes**](CweApi.md#GetCwes) | **Get** /v1/cwe | Returns a list of all CWEs


# **GetCwe**
> Cwe GetCwe(ctx, cweId)
Returns a specific CWE



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cweId** | **int32**| The CWE ID of the CWE to retrieve | 

### Return type

[**Cwe**](Cwe.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCwes**
> []Cwe GetCwes(ctx, )
Returns a list of all CWEs



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Cwe**](Cwe.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

