# \ProjectPropertyApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProperty**](ProjectPropertyApi.md#CreateProperty) | **Put** /v1/project/{uuid}/property | Creates a new project property
[**DeleteProperty**](ProjectPropertyApi.md#DeleteProperty) | **Delete** /v1/project/{uuid}/property | Deletes a config property
[**GetProperties**](ProjectPropertyApi.md#GetProperties) | **Get** /v1/project/{uuid}/property | Returns a list of all ProjectProperties for the specified project
[**UpdateProperty**](ProjectPropertyApi.md#UpdateProperty) | **Post** /v1/project/{uuid}/property | Updates a project property


# **CreateProperty**
> ProjectProperty CreateProperty(ctx, uuid, optional)
Creates a new project property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to create a property for | 
 **optional** | ***ProjectPropertyApiCreatePropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectPropertyApiCreatePropertyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProjectProperty**](ProjectProperty.md)|  | 

### Return type

[**ProjectProperty**](ProjectProperty.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProperty**
> ProjectProperty DeleteProperty(ctx, uuid, optional)
Deletes a config property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to delete a property from | 
 **optional** | ***ProjectPropertyApiDeletePropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectPropertyApiDeletePropertyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProjectProperty**](ProjectProperty.md)|  | 

### Return type

[**ProjectProperty**](ProjectProperty.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProperties**
> []ProjectProperty GetProperties(ctx, uuid)
Returns a list of all ProjectProperties for the specified project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to retrieve properties for | 

### Return type

[**[]ProjectProperty**](ProjectProperty.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProperty**
> ProjectProperty UpdateProperty(ctx, uuid, optional)
Updates a project property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to create a property for | 
 **optional** | ***ProjectPropertyApiUpdatePropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectPropertyApiUpdatePropertyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProjectProperty**](ProjectProperty.md)|  | 

### Return type

[**ProjectProperty**](ProjectProperty.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

