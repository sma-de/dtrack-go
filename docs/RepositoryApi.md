# \RepositoryApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRepository**](RepositoryApi.md#CreateRepository) | **Put** /v1/repository | Creates a new repository
[**DeleteRepository**](RepositoryApi.md#DeleteRepository) | **Delete** /v1/repository/{uuid} | Deletes a repository
[**GetRepositories**](RepositoryApi.md#GetRepositories) | **Get** /v1/repository | Returns a list of all repositories
[**GetRepositoriesByType**](RepositoryApi.md#GetRepositoriesByType) | **Get** /v1/repository/{type} | Returns repositories that support the specific type
[**GetRepositoryMetaComponent**](RepositoryApi.md#GetRepositoryMetaComponent) | **Get** /v1/repository/latest | Attempts to resolve the latest version of the component available in the configured repositories
[**UpdateRepository**](RepositoryApi.md#UpdateRepository) | **Post** /v1/repository | Updates a repository


# **CreateRepository**
> Repository CreateRepository(ctx, optional)
Creates a new repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryApiCreateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCreateRepositoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Repository**](Repository.md)|  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepository**
> DeleteRepository(ctx, uuid)
Deletes a repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the repository to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositories**
> []Repository GetRepositories(ctx, )
Returns a list of all repositories



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Repository**](Repository.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoriesByType**
> []Repository GetRepositoriesByType(ctx, type_)
Returns repositories that support the specific type



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The type of repositories to retrieve | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryMetaComponent**
> RepositoryMetaComponent GetRepositoryMetaComponent(ctx, purl)
Attempts to resolve the latest version of the component available in the configured repositories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purl** | **string**| The Package URL for the component to query | 

### Return type

[**RepositoryMetaComponent**](RepositoryMetaComponent.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository**
> Repository UpdateRepository(ctx, optional)
Updates a repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryApiUpdateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiUpdateRepositoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Repository**](Repository.md)|  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

