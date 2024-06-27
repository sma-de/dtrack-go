# \MetricsApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponentCurrentMetrics**](MetricsApi.md#GetComponentCurrentMetrics) | **Get** /v1/metrics/component/{uuid}/current | Returns current metrics for a specific component
[**GetComponentMetricsSince**](MetricsApi.md#GetComponentMetricsSince) | **Get** /v1/metrics/component/{uuid}/since/{date} | Returns historical metrics for a specific component from a specific date
[**GetComponentMetricsXDays**](MetricsApi.md#GetComponentMetricsXDays) | **Get** /v1/metrics/component/{uuid}/days/{days} | Returns X days of historical metrics for a specific component
[**GetPortfolioCurrentMetrics**](MetricsApi.md#GetPortfolioCurrentMetrics) | **Get** /v1/metrics/portfolio/current | Returns current metrics for the entire portfolio
[**GetPortfolioMetricsSince**](MetricsApi.md#GetPortfolioMetricsSince) | **Get** /v1/metrics/portfolio/since/{date} | Returns historical metrics for the entire portfolio from a specific date
[**GetPortfolioMetricsXDays**](MetricsApi.md#GetPortfolioMetricsXDays) | **Get** /v1/metrics/portfolio/{days}/days | Returns X days of historical metrics for the entire portfolio
[**GetProjectCurrentMetrics**](MetricsApi.md#GetProjectCurrentMetrics) | **Get** /v1/metrics/project/{uuid}/current | Returns current metrics for a specific project
[**GetProjectMetricsSince**](MetricsApi.md#GetProjectMetricsSince) | **Get** /v1/metrics/project/{uuid}/since/{date} | Returns historical metrics for a specific project from a specific date
[**GetProjectMetricsXDays**](MetricsApi.md#GetProjectMetricsXDays) | **Get** /v1/metrics/project/{uuid}/days/{days} | Returns X days of historical metrics for a specific project
[**GetVulnerabilityMetrics**](MetricsApi.md#GetVulnerabilityMetrics) | **Get** /v1/metrics/vulnerability | Returns the sum of all vulnerabilities in the database by year and month
[**RefreshComponentMetrics**](MetricsApi.md#RefreshComponentMetrics) | **Get** /v1/metrics/component/{uuid}/refresh | Requests a refresh of a specific components metrics
[**RefreshPortfolioMetrics**](MetricsApi.md#RefreshPortfolioMetrics) | **Get** /v1/metrics/portfolio/refresh | Requests a refresh of the portfolio metrics
[**RefreshProjectMetrics**](MetricsApi.md#RefreshProjectMetrics) | **Get** /v1/metrics/project/{uuid}/refresh | Requests a refresh of a specific projects metrics


# **GetComponentCurrentMetrics**
> DependencyMetrics GetComponentCurrentMetrics(ctx, uuid)
Returns current metrics for a specific component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the component to retrieve metrics for | 

### Return type

[**DependencyMetrics**](DependencyMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentMetricsSince**
> []DependencyMetrics GetComponentMetricsSince(ctx, uuid, date)
Returns historical metrics for a specific component from a specific date

Date format must be YYYYMMDD

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the component to retrieve metrics for | 
  **date** | **string**| The start date to retrieve metrics for | 

### Return type

[**[]DependencyMetrics**](DependencyMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentMetricsXDays**
> []DependencyMetrics GetComponentMetricsXDays(ctx, uuid, days)
Returns X days of historical metrics for a specific component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the component to retrieve metrics for | 
  **days** | **int32**| The number of days back to retrieve metrics for | 

### Return type

[**[]DependencyMetrics**](DependencyMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortfolioCurrentMetrics**
> PortfolioMetrics GetPortfolioCurrentMetrics(ctx, )
Returns current metrics for the entire portfolio



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PortfolioMetrics**](PortfolioMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortfolioMetricsSince**
> []PortfolioMetrics GetPortfolioMetricsSince(ctx, date)
Returns historical metrics for the entire portfolio from a specific date

Date format must be YYYYMMDD

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **date** | **string**| The start date to retrieve metrics for | 

### Return type

[**[]PortfolioMetrics**](PortfolioMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortfolioMetricsXDays**
> []PortfolioMetrics GetPortfolioMetricsXDays(ctx, days)
Returns X days of historical metrics for the entire portfolio



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **days** | **int32**| The number of days back to retrieve metrics for | 

### Return type

[**[]PortfolioMetrics**](PortfolioMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectCurrentMetrics**
> ProjectMetrics GetProjectCurrentMetrics(ctx, uuid)
Returns current metrics for a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to retrieve metrics for | 

### Return type

[**ProjectMetrics**](ProjectMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectMetricsSince**
> []ProjectMetrics GetProjectMetricsSince(ctx, uuid, date)
Returns historical metrics for a specific project from a specific date

Date format must be YYYYMMDD

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to retrieve metrics for | 
  **date** | **string**| The start date to retrieve metrics for | 

### Return type

[**[]ProjectMetrics**](ProjectMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectMetricsXDays**
> []ProjectMetrics GetProjectMetricsXDays(ctx, uuid, days)
Returns X days of historical metrics for a specific project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to retrieve metrics for | 
  **days** | **int32**| The number of days back to retrieve metrics for | 

### Return type

[**[]ProjectMetrics**](ProjectMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVulnerabilityMetrics**
> []VulnerabilityMetrics GetVulnerabilityMetrics(ctx, )
Returns the sum of all vulnerabilities in the database by year and month



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]VulnerabilityMetrics**](VulnerabilityMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshComponentMetrics**
> RefreshComponentMetrics(ctx, uuid)
Requests a refresh of a specific components metrics



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the component to refresh metrics on | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshPortfolioMetrics**
> PortfolioMetrics RefreshPortfolioMetrics(ctx, )
Requests a refresh of the portfolio metrics



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PortfolioMetrics**](PortfolioMetrics.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshProjectMetrics**
> RefreshProjectMetrics(ctx, uuid)
Requests a refresh of a specific projects metrics



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the project to refresh metrics on | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

