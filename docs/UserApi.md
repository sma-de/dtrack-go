# \UserApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeamToUser**](UserApi.md#AddTeamToUser) | **Post** /v1/user/{username}/membership | Adds the username to the specified team.
[**CreateLdapUser**](UserApi.md#CreateLdapUser) | **Put** /v1/user/ldap | Creates a new user that references an existing LDAP object.
[**CreateManagedUser**](UserApi.md#CreateManagedUser) | **Put** /v1/user/managed | Creates a new user.
[**CreateOidcUser**](UserApi.md#CreateOidcUser) | **Put** /v1/user/oidc | Creates a new user that references an existing OpenID Connect user.
[**DeleteLdapUser**](UserApi.md#DeleteLdapUser) | **Delete** /v1/user/ldap | Deletes a user.
[**DeleteManagedUser**](UserApi.md#DeleteManagedUser) | **Delete** /v1/user/managed | Deletes a user.
[**DeleteOidcUser**](UserApi.md#DeleteOidcUser) | **Delete** /v1/user/oidc | Deletes an OpenID Connect user.
[**ForceChangePassword**](UserApi.md#ForceChangePassword) | **Post** /v1/user/forceChangePassword | Asserts login credentials and upon successful authentication, verifies passwords match and changes users password
[**GetLdapUsers**](UserApi.md#GetLdapUsers) | **Get** /v1/user/ldap | Returns a list of all LDAP users
[**GetManagedUsers**](UserApi.md#GetManagedUsers) | **Get** /v1/user/managed | Returns a list of all managed users
[**GetOidcUsers**](UserApi.md#GetOidcUsers) | **Get** /v1/user/oidc | Returns a list of all OIDC users
[**GetSelf1**](UserApi.md#GetSelf1) | **Get** /v1/user/self | Returns information about the current logged in user.
[**RemoveTeamFromUser**](UserApi.md#RemoveTeamFromUser) | **Delete** /v1/user/{username}/membership | Removes the username from the specified team.
[**UpdateManagedUser**](UserApi.md#UpdateManagedUser) | **Post** /v1/user/managed | Updates a managed user.
[**UpdateSelf**](UserApi.md#UpdateSelf) | **Post** /v1/user/self | Updates information about the current logged in user.
[**ValidateCredentials**](UserApi.md#ValidateCredentials) | **Post** /v1/user/login | Assert login credentials
[**ValidateOidcAccessToken**](UserApi.md#ValidateOidcAccessToken) | **Post** /v1/user/oidc/login | Login with OpenID Connect


# **AddTeamToUser**
> UserPrincipal AddTeamToUser(ctx, username, body)
Adds the username to the specified team.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| A valid username | 
  **body** | [**IdentifiableObject**](IdentifiableObject.md)| The UUID of the team to associate username with | 

### Return type

[**UserPrincipal**](UserPrincipal.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLdapUser**
> LdapUser CreateLdapUser(ctx, optional)
Creates a new user that references an existing LDAP object.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiCreateLdapUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateLdapUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LdapUser**](LdapUser.md)|  | 

### Return type

[**LdapUser**](LdapUser.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateManagedUser**
> ManagedUser CreateManagedUser(ctx, optional)
Creates a new user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiCreateManagedUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateManagedUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ManagedUser**](ManagedUser.md)|  | 

### Return type

[**ManagedUser**](ManagedUser.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOidcUser**
> OidcUser CreateOidcUser(ctx, optional)
Creates a new user that references an existing OpenID Connect user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiCreateOidcUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateOidcUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OidcUser**](OidcUser.md)|  | 

### Return type

[**OidcUser**](OidcUser.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLdapUser**
> DeleteLdapUser(ctx, optional)
Deletes a user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiDeleteLdapUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiDeleteLdapUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LdapUser**](LdapUser.md)|  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteManagedUser**
> DeleteManagedUser(ctx, optional)
Deletes a user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiDeleteManagedUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiDeleteManagedUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ManagedUser**](ManagedUser.md)|  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOidcUser**
> DeleteOidcUser(ctx, optional)
Deletes an OpenID Connect user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiDeleteOidcUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiDeleteOidcUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OidcUser**](OidcUser.md)|  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForceChangePassword**
> string ForceChangePassword(ctx, optional)
Asserts login credentials and upon successful authentication, verifies passwords match and changes users password

Upon a successful login, a JSON Web Token will be returned in the response body. This functionality requires authentication to be enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiForceChangePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiForceChangePasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**|  | 
 **password** | **optional.String**|  | 
 **newPassword** | **optional.String**|  | 
 **confirmPassword** | **optional.String**|  | 

### Return type

**string**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLdapUsers**
> []LdapUser GetLdapUsers(ctx, )
Returns a list of all LDAP users



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LdapUser**](LdapUser.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetManagedUsers**
> []ManagedUser GetManagedUsers(ctx, )
Returns a list of all managed users



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ManagedUser**](ManagedUser.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOidcUsers**
> []OidcUser GetOidcUsers(ctx, )
Returns a list of all OIDC users



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OidcUser**](OidcUser.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSelf1**
> UserPrincipal GetSelf1(ctx, )
Returns information about the current logged in user.



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserPrincipal**](UserPrincipal.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTeamFromUser**
> UserPrincipal RemoveTeamFromUser(ctx, username, body)
Removes the username from the specified team.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| A valid username | 
  **body** | [**IdentifiableObject**](IdentifiableObject.md)| The UUID of the team to un-associate username from | 

### Return type

[**UserPrincipal**](UserPrincipal.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateManagedUser**
> ManagedUser UpdateManagedUser(ctx, optional)
Updates a managed user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiUpdateManagedUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateManagedUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ManagedUser**](ManagedUser.md)|  | 

### Return type

[**ManagedUser**](ManagedUser.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSelf**
> UserPrincipal UpdateSelf(ctx, optional)
Updates information about the current logged in user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiUpdateSelfOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateSelfOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ManagedUser**](ManagedUser.md)|  | 

### Return type

[**UserPrincipal**](UserPrincipal.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateCredentials**
> string ValidateCredentials(ctx, optional)
Assert login credentials

Upon a successful login, a JSON Web Token will be returned in the response body. This functionality requires authentication to be enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiValidateCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiValidateCredentialsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**|  | 
 **password** | **optional.String**|  | 

### Return type

**string**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateOidcAccessToken**
> string ValidateOidcAccessToken(ctx, idToken, optional)
Login with OpenID Connect

Upon a successful login, a JSON Web Token will be returned in the response body. This functionality requires authentication to be enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idToken** | **string**| An OAuth2 access token | 
 **optional** | ***UserApiValidateOidcAccessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiValidateOidcAccessTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**|  | 

### Return type

**string**

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

