# \LicenseGroupApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLicenseToLicenseGroup**](LicenseGroupApi.md#AddLicenseToLicenseGroup) | **Post** /v1/licenseGroup/{uuid}/license/{licenseUuid} | Adds the license to the specified license group.
[**CreateLicenseGroup**](LicenseGroupApi.md#CreateLicenseGroup) | **Put** /v1/licenseGroup | Creates a new license group
[**DeleteLicenseGroup**](LicenseGroupApi.md#DeleteLicenseGroup) | **Delete** /v1/licenseGroup/{uuid} | Deletes a license group
[**GetLicenseGroup**](LicenseGroupApi.md#GetLicenseGroup) | **Get** /v1/licenseGroup/{uuid} | Returns a specific license group
[**GetLicenseGroups**](LicenseGroupApi.md#GetLicenseGroups) | **Get** /v1/licenseGroup | Returns a list of all license groups
[**RemoveLicenseFromLicenseGroup**](LicenseGroupApi.md#RemoveLicenseFromLicenseGroup) | **Delete** /v1/licenseGroup/{uuid}/license/{licenseUuid} | Removes the license from the license group.
[**UpdateLicenseGroup**](LicenseGroupApi.md#UpdateLicenseGroup) | **Post** /v1/licenseGroup | Updates a license group


# **AddLicenseToLicenseGroup**
> LicenseGroup AddLicenseToLicenseGroup(ctx, uuid, licenseUuid)
Adds the license to the specified license group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| A valid license group | 
  **licenseUuid** | **string**| A valid license | 

### Return type

[**LicenseGroup**](LicenseGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLicenseGroup**
> LicenseGroup CreateLicenseGroup(ctx, optional)
Creates a new license group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LicenseGroupApiCreateLicenseGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LicenseGroupApiCreateLicenseGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LicenseGroup**](LicenseGroup.md)|  | 

### Return type

[**LicenseGroup**](LicenseGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLicenseGroup**
> DeleteLicenseGroup(ctx, uuid)
Deletes a license group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the license group to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseGroup**
> License GetLicenseGroup(ctx, uuid)
Returns a specific license group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the license group to retrieve | 

### Return type

[**License**](License.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseGroups**
> []LicenseGroup GetLicenseGroups(ctx, )
Returns a list of all license groups



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LicenseGroup**](LicenseGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveLicenseFromLicenseGroup**
> LicenseGroup RemoveLicenseFromLicenseGroup(ctx, uuid, licenseUuid)
Removes the license from the license group.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| A valid license group | 
  **licenseUuid** | **string**| A valid license | 

### Return type

[**LicenseGroup**](LicenseGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLicenseGroup**
> LicenseGroup UpdateLicenseGroup(ctx, optional)
Updates a license group



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LicenseGroupApiUpdateLicenseGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LicenseGroupApiUpdateLicenseGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LicenseGroup**](LicenseGroup.md)|  | 

### Return type

[**LicenseGroup**](LicenseGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

