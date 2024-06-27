# \BomApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportComponentAsCycloneDx**](BomApi.md#ExportComponentAsCycloneDx) | **Get** /v1/bom/cyclonedx/component/{uuid} | Returns dependency metadata for a specific component in CycloneDX format
[**ExportProjectAsCycloneDx**](BomApi.md#ExportProjectAsCycloneDx) | **Get** /v1/bom/cyclonedx/project/{uuid} | Returns dependency metadata for a project in CycloneDX format
[**IsTokenBeingProcessed**](BomApi.md#IsTokenBeingProcessed) | **Get** /v1/bom/token/{uuid} | Determines if there are any tasks associated with the token that are being processed, or in the queue to be processed.
[**UploadBom**](BomApi.md#UploadBom) | **Post** /v1/bom | Upload a supported bill of material format document
[**UploadBomBase64Encoded**](BomApi.md#UploadBomBase64Encoded) | **Put** /v1/bom | Upload a supported bill of material format document


# **ExportComponentAsCycloneDx**
> string ExportComponentAsCycloneDx(ctx, uuid, optional)
Returns dependency metadata for a specific component in CycloneDX format



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the component to export | 
 **optional** | ***BomApiExportComponentAsCycloneDxOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BomApiExportComponentAsCycloneDxOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**| The format to output (defaults to JSON) | 

### Return type

**string**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.cyclonedx+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportProjectAsCycloneDx**
> string ExportProjectAsCycloneDx(ctx, uuid, optional)
Returns dependency metadata for a project in CycloneDX format



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to export | 
 **optional** | ***BomApiExportProjectAsCycloneDxOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BomApiExportProjectAsCycloneDxOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**| The format to output (defaults to JSON) | 
 **variant** | **optional.String**| Specifies the CycloneDX variant to export. Value options are &#39;inventory&#39; and &#39;withVulnerabilities&#39;. (defaults to &#39;inventory&#39;) | 
 **download** | **optional.Bool**| Force the resulting BOM to be downloaded as a file (defaults to &#39;false&#39;) | 

### Return type

**string**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.cyclonedx+xml, application/vnd.cyclonedx+json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsTokenBeingProcessed**
> IsTokenBeingProcessedResponse IsTokenBeingProcessed(ctx, uuid)
Determines if there are any tasks associated with the token that are being processed, or in the queue to be processed.

This endpoint is intended to be used in conjunction with uploading a supported BOM document. Upon upload, a token will be returned. The token can then be queried using this endpoint to determine if any tasks (such as vulnerability analysis) is being performed on the BOM. A value of true indicates processing is occurring. A value of false indicates that no processing is occurring for the specified token. However, a value of false also does not confirm the token is valid, only that no processing is associated with the specified token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the token to query | 

### Return type

[**IsTokenBeingProcessedResponse**](IsTokenBeingProcessedResponse.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBom**
> BomUploadResponse UploadBom(ctx, optional)
Upload a supported bill of material format document

Expects CycloneDX along and a valid project UUID. If a UUID is not specified, then the projectName and projectVersion must be specified. Optionally, if autoCreate is specified and 'true' and the project does not exist, the project will be created. In this scenario, the principal making the request will additionally need the PORTFOLIO_MANAGEMENT or PROJECT_CREATION_UPLOAD permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BomApiUploadBomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BomApiUploadBomOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.String**|  | 
 **autoCreate** | **optional.Bool**|  | [default to false]
 **projectName** | **optional.String**|  | 
 **projectVersion** | **optional.String**|  | 
 **parentName** | **optional.String**|  | 
 **parentVersion** | **optional.String**|  | 
 **parentUUID** | **optional.String**|  | 
 **body** | [**optional.Interface of FormDataMultiPart**](FormDataMultiPart.md)|  | 

### Return type

[**BomUploadResponse**](BomUploadResponse.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBomBase64Encoded**
> BomUploadResponse UploadBomBase64Encoded(ctx, optional)
Upload a supported bill of material format document

Expects CycloneDX along and a valid project UUID. If a UUID is not specified, then the projectName and projectVersion must be specified. Optionally, if autoCreate is specified and 'true' and the project does not exist, the project will be created. In this scenario, the principal making the request will additionally need the PORTFOLIO_MANAGEMENT or PROJECT_CREATION_UPLOAD permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BomApiUploadBomBase64EncodedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BomApiUploadBomBase64EncodedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BomSubmitRequest**](BomSubmitRequest.md)|  | 

### Return type

[**BomUploadResponse**](BomUploadResponse.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

