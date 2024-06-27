# \ConfigPropertyApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigProperties**](ConfigPropertyApi.md#GetConfigProperties) | **Get** /v1/configProperty | Returns a list of all ConfigProperties for the specified groupName
[**UpdateConfigProperty**](ConfigPropertyApi.md#UpdateConfigProperty) | **Post** /v1/configProperty/aggregate | Updates an array of config properties
[**UpdateConfigProperty1**](ConfigPropertyApi.md#UpdateConfigProperty1) | **Post** /v1/configProperty | Updates a config property


# **GetConfigProperties**
> []ConfigProperty GetConfigProperties(ctx, )
Returns a list of all ConfigProperties for the specified groupName



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConfigProperty**](ConfigProperty.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfigProperty**
> []ConfigProperty UpdateConfigProperty(ctx, optional)
Updates an array of config properties



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigPropertyApiUpdateConfigPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigPropertyApiUpdateConfigPropertyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []ConfigProperty**](ConfigProperty.md)|  | 

### Return type

[**[]ConfigProperty**](ConfigProperty.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfigProperty1**
> ConfigProperty UpdateConfigProperty1(ctx, optional)
Updates a config property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigPropertyApiUpdateConfigProperty1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigPropertyApiUpdateConfigProperty1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ConfigProperty**](ConfigProperty.md)|  | 

### Return type

[**ConfigProperty**](ConfigProperty.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

