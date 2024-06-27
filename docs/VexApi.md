# \VexApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportProjectAsCycloneDx1**](VexApi.md#ExportProjectAsCycloneDx1) | **Get** /v1/vex/cyclonedx/project/{uuid} | Returns a VEX for a project in CycloneDX format
[**UploadVex**](VexApi.md#UploadVex) | **Post** /v1/vex | Upload a supported VEX document
[**UploadVex1**](VexApi.md#UploadVex1) | **Put** /v1/vex | Upload a supported VEX document


# **ExportProjectAsCycloneDx1**
> string ExportProjectAsCycloneDx1(ctx, uuid, optional)
Returns a VEX for a project in CycloneDX format



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to export | 
 **optional** | ***VexApiExportProjectAsCycloneDx1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VexApiExportProjectAsCycloneDx1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **download** | **optional.Bool**| Force the resulting VEX to be downloaded as a file (defaults to &#39;false&#39;) | 

### Return type

**string**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.cyclonedx+json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadVex**
> UploadVex(ctx, optional)
Upload a supported VEX document

Expects CycloneDX along and a valid project UUID. If a UUID is not specified, than the projectName and projectVersion must be specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VexApiUploadVexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VexApiUploadVexOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.String**|  | 
 **projectName** | **optional.String**|  | 
 **projectVersion** | **optional.String**|  | 
 **body** | [**optional.Interface of FormDataMultiPart**](FormDataMultiPart.md)|  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadVex1**
> UploadVex1(ctx, optional)
Upload a supported VEX document

Expects CycloneDX along and a valid project UUID. If a UUID is not specified then the projectName and projectVersion must be specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VexApiUploadVex1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VexApiUploadVex1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VexSubmitRequest**](VexSubmitRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

