# \PolicyConditionApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyCondition**](PolicyConditionApi.md#CreatePolicyCondition) | **Put** /v1/policy/{uuid}/condition | Creates a new policy condition
[**DeletePolicyCondition**](PolicyConditionApi.md#DeletePolicyCondition) | **Delete** /v1/policy/condition/{uuid} | Deletes a policy condition
[**UpdatePolicyCondition**](PolicyConditionApi.md#UpdatePolicyCondition) | **Post** /v1/policy/condition | Updates a policy condition


# **CreatePolicyCondition**
> PolicyCondition CreatePolicyCondition(ctx, uuid, optional)
Creates a new policy condition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the policy | 
 **optional** | ***PolicyConditionApiCreatePolicyConditionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyConditionApiCreatePolicyConditionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PolicyCondition**](PolicyCondition.md)|  | 

### Return type

[**PolicyCondition**](PolicyCondition.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyCondition**
> DeletePolicyCondition(ctx, uuid)
Deletes a policy condition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the policy condition to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyCondition**
> PolicyCondition UpdatePolicyCondition(ctx, optional)
Updates a policy condition



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyConditionApiUpdatePolicyConditionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyConditionApiUpdatePolicyConditionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PolicyCondition**](PolicyCondition.md)|  | 

### Return type

[**PolicyCondition**](PolicyCondition.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

