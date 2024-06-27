# \TeamApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeam**](TeamApi.md#CreateTeam) | **Put** /v1/team | Creates a new team along with an associated API key
[**DeleteApiKey**](TeamApi.md#DeleteApiKey) | **Delete** /v1/team/key/{apikey} | Deletes the specified API key
[**DeleteTeam**](TeamApi.md#DeleteTeam) | **Delete** /v1/team | Deletes a team
[**GenerateApiKey**](TeamApi.md#GenerateApiKey) | **Put** /v1/team/{uuid}/key | Generates an API key and returns its value
[**GetSelf**](TeamApi.md#GetSelf) | **Get** /v1/team/self | Returns information about the current team.
[**GetTeam**](TeamApi.md#GetTeam) | **Get** /v1/team/{uuid} | Returns a specific team
[**GetTeams**](TeamApi.md#GetTeams) | **Get** /v1/team | Returns a list of all teams
[**RegenerateApiKey**](TeamApi.md#RegenerateApiKey) | **Post** /v1/team/key/{apikey} | Regenerates an API key by removing the specified key, generating a new one and returning its value
[**UpdateTeam**](TeamApi.md#UpdateTeam) | **Post** /v1/team | Updates a team&#39;s fields including


# **CreateTeam**
> Team CreateTeam(ctx, optional)
Creates a new team along with an associated API key



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamApiCreateTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamApiCreateTeamOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Team**](Team.md)|  | 

### Return type

[**Team**](Team.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiKey**
> DeleteApiKey(ctx, apikey)
Deletes the specified API key



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| The API key to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTeam**
> DeleteTeam(ctx, optional)
Deletes a team



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamApiDeleteTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamApiDeleteTeamOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Team**](Team.md)|  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateApiKey**
> ApiKey GenerateApiKey(ctx, uuid)
Generates an API key and returns its value



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the team to generate a key for | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSelf**
> TeamSelfResponse GetSelf(ctx, )
Returns information about the current team.



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TeamSelfResponse**](TeamSelfResponse.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeam**
> Team GetTeam(ctx, uuid)
Returns a specific team



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the team to retrieve | 

### Return type

[**Team**](Team.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeams**
> []Team GetTeams(ctx, )
Returns a list of all teams



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Team**](Team.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateApiKey**
> ApiKey RegenerateApiKey(ctx, apikey)
Regenerates an API key by removing the specified key, generating a new one and returning its value



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| The API key to regenerate | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeam**
> Team UpdateTeam(ctx, optional)
Updates a team's fields including



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamApiUpdateTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamApiUpdateTeamOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Team**](Team.md)|  | 

### Return type

[**Team**](Team.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

