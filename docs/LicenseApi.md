# \LicenseApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLicense**](LicenseApi.md#CreateLicense) | **Put** /v1/license | Creates a new custom license
[**DeleteLicense**](LicenseApi.md#DeleteLicense) | **Delete** /v1/license/{licenseId} | Deletes a custom license
[**GetLicense**](LicenseApi.md#GetLicense) | **Get** /v1/license/{licenseId} | Returns a specific license
[**GetLicenseListing**](LicenseApi.md#GetLicenseListing) | **Get** /v1/license/concise | Returns a concise listing of all licenses
[**GetLicenses**](LicenseApi.md#GetLicenses) | **Get** /v1/license | Returns a list of all licenses with complete metadata for each license


# **CreateLicense**
> License CreateLicense(ctx, optional)
Creates a new custom license



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LicenseApiCreateLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LicenseApiCreateLicenseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of License**](License.md)|  | 

### Return type

[**License**](License.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLicense**
> DeleteLicense(ctx, licenseId)
Deletes a custom license



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **licenseId** | **string**| The SPDX License ID of the license to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicense**
> License GetLicense(ctx, licenseId)
Returns a specific license



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **licenseId** | **string**| The SPDX License ID of the license to retrieve | 

### Return type

[**License**](License.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseListing**
> []License GetLicenseListing(ctx, )
Returns a concise listing of all licenses



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]License**](License.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenses**
> []License GetLicenses(ctx, )
Returns a list of all licenses with complete metadata for each license



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]License**](License.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

