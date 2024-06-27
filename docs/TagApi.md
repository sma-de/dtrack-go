# \TagApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTags**](TagApi.md#GetTags) | **Get** /v1/tag/{policyUuid} | Returns a list of all tags


# **GetTags**
> []Tag GetTags(ctx, policyUuid)
Returns a list of all tags



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyUuid** | **string**| The UUID of the policy | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

