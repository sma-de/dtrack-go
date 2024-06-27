# \DependencyGraphApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponentsAndServicesByComponentUuid**](DependencyGraphApi.md#GetComponentsAndServicesByComponentUuid) | **Get** /v1/dependencyGraph/component/{uuid}/directDependencies | Returns a list of specific components and services from component UUID
[**GetComponentsAndServicesByProjectUuid**](DependencyGraphApi.md#GetComponentsAndServicesByProjectUuid) | **Get** /v1/dependencyGraph/project/{uuid}/directDependencies | Returns a list of specific components and services from project UUID


# **GetComponentsAndServicesByComponentUuid**
> []DependencyGraphResponse GetComponentsAndServicesByComponentUuid(ctx, uuid)
Returns a list of specific components and services from component UUID



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 

### Return type

[**[]DependencyGraphResponse**](DependencyGraphResponse.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentsAndServicesByProjectUuid**
> []DependencyGraphResponse GetComponentsAndServicesByProjectUuid(ctx, uuid)
Returns a list of specific components and services from project UUID



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 

### Return type

[**[]DependencyGraphResponse**](DependencyGraphResponse.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

