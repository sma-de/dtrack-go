# \ComponentApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComponent**](ComponentApi.md#CreateComponent) | **Put** /v1/component/project/{uuid} | Creates a new component
[**DeleteComponent**](ComponentApi.md#DeleteComponent) | **Delete** /v1/component/{uuid} | Deletes a component
[**GetAllComponents**](ComponentApi.md#GetAllComponents) | **Get** /v1/component/project/{uuid} | Returns a list of all components for a given project
[**GetComponentByHash**](ComponentApi.md#GetComponentByHash) | **Get** /v1/component/hash/{hash} | Returns a list of components that have the specified hash value
[**GetComponentByIdentity**](ComponentApi.md#GetComponentByIdentity) | **Get** /v1/component/identity | Returns a list of components that have the specified component identity. This resource accepts coordinates (group, name, version) or purl, cpe, or swidTagId
[**GetComponentByUuid**](ComponentApi.md#GetComponentByUuid) | **Get** /v1/component/{uuid} | Returns a specific component
[**GetDependencyGraphForComponent**](ComponentApi.md#GetDependencyGraphForComponent) | **Get** /v1/component/project/{projectUuid}/dependencyGraph/{componentUuid} | Returns the expanded dependency graph to every occurrence of a component
[**IdentifyInternalComponents**](ComponentApi.md#IdentifyInternalComponents) | **Get** /v1/component/internal/identify | Requests the identification of internal components in the portfolio
[**UpdateComponent**](ComponentApi.md#UpdateComponent) | **Post** /v1/component | Updates a component


# **CreateComponent**
> Component CreateComponent(ctx, uuid, optional)
Creates a new component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 
 **optional** | ***ComponentApiCreateComponentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentApiCreateComponentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Component**](Component.md)|  | 

### Return type

[**Component**](Component.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteComponent**
> DeleteComponent(ctx, uuid)
Deletes a component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the component to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllComponents**
> []Component GetAllComponents(ctx, uuid, optional)
Returns a list of all components for a given project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to retrieve components for | 
 **optional** | ***ComponentApiGetAllComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentApiGetAllComponentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onlyOutdated** | **optional.Bool**| Optionally exclude recent components so only outdated components are returned | 
 **onlyDirect** | **optional.Bool**| Optionally exclude transitive dependencies so only direct dependencies are returned | 

### Return type

[**[]Component**](Component.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentByHash**
> []Component GetComponentByHash(ctx, hash)
Returns a list of components that have the specified hash value



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| The MD5, SHA-1, SHA-256, SHA-384, SHA-512, SHA3-256, SHA3-384, SHA3-512, BLAKE2b-256, BLAKE2b-384, BLAKE2b-512, or BLAKE3 hash of the component to retrieve | 

### Return type

[**[]Component**](Component.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentByIdentity**
> []Component GetComponentByIdentity(ctx, optional)
Returns a list of components that have the specified component identity. This resource accepts coordinates (group, name, version) or purl, cpe, or swidTagId



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComponentApiGetComponentByIdentityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentApiGetComponentByIdentityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **optional.String**| The group of the component | 
 **name** | **optional.String**| The name of the component | 
 **version** | **optional.String**| The version of the component | 
 **purl** | **optional.String**| The purl of the component | 
 **cpe** | **optional.String**| The cpe of the component | 
 **swidTagId** | **optional.String**| The swidTagId of the component | 
 **project** | **optional.String**| The project the component belongs to | 

### Return type

[**[]Component**](Component.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentByUuid**
> Component GetComponentByUuid(ctx, uuid, optional)
Returns a specific component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the component to retrieve | 
 **optional** | ***ComponentApiGetComponentByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentApiGetComponentByUuidOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeRepositoryMetaData** | **optional.Bool**| Optionally includes third-party metadata about the component from external repositories | 

### Return type

[**Component**](Component.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDependencyGraphForComponent**
> map[string]Component GetDependencyGraphForComponent(ctx, projectUuid, componentUuid)
Returns the expanded dependency graph to every occurrence of a component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUuid** | **string**| The UUID of the project to get the expanded dependency graph for | 
  **componentUuid** | **string**| The UUID of the component to get the expanded dependency graph for | 

### Return type

[**map[string]Component**](Component.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IdentifyInternalComponents**
> IdentifyInternalComponents(ctx, )
Requests the identification of internal components in the portfolio



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateComponent**
> Component UpdateComponent(ctx, optional)
Updates a component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComponentApiUpdateComponentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentApiUpdateComponentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Component**](Component.md)|  | 

### Return type

[**Component**](Component.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

