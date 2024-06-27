# \LdapApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMapping1**](LdapApi.md#AddMapping1) | **Put** /v1/ldap/mapping | Adds a mapping
[**DeleteMapping1**](LdapApi.md#DeleteMapping1) | **Delete** /v1/ldap/mapping/{uuid} | Removes a mapping
[**RetrieveLdapGroups**](LdapApi.md#RetrieveLdapGroups) | **Get** /v1/ldap/groups | Returns the DNs of all accessible groups within the directory
[**RetrieveLdapGroups1**](LdapApi.md#RetrieveLdapGroups1) | **Get** /v1/ldap/team/{uuid} | Returns the DNs of all groups mapped to the specified team


# **AddMapping1**
> MappedLdapGroup AddMapping1(ctx, optional)
Adds a mapping



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapApiAddMapping1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapApiAddMapping1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MappedLdapGroupRequest**](MappedLdapGroupRequest.md)|  | 

### Return type

[**MappedLdapGroup**](MappedLdapGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMapping1**
> MappedLdapGroup DeleteMapping1(ctx, uuid)
Removes a mapping



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the mapping to delete | 

### Return type

[**MappedLdapGroup**](MappedLdapGroup.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLdapGroups**
> []string RetrieveLdapGroups(ctx, )
Returns the DNs of all accessible groups within the directory

This API performs a pass-thru query to the configured LDAP server. Search criteria results are cached using default Alpine CacheManager policy

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLdapGroups1**
> []string RetrieveLdapGroups1(ctx, uuid)
Returns the DNs of all groups mapped to the specified team



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The UUID of the team to retrieve mappings for | 

### Return type

**[]string**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

