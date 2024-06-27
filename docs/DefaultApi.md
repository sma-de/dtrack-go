# \DefaultApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProjectToRule**](DefaultApi.md#AddProjectToRule) | **Post** /v1/notification/rule/{ruleUuid}/project/{projectUuid} | Adds a project to a notification rule
[**AddTeamToRule**](DefaultApi.md#AddTeamToRule) | **Post** /v1/notification/rule/{ruleUuid}/team/{teamUuid} | Adds a team to a notification rule
[**CreateNotificationPublisher**](DefaultApi.md#CreateNotificationPublisher) | **Put** /v1/notification/publisher | Creates a new notification publisher
[**CreateNotificationRule**](DefaultApi.md#CreateNotificationRule) | **Put** /v1/notification/rule | Creates a new notification rule
[**DeleteNotificationPublisher**](DefaultApi.md#DeleteNotificationPublisher) | **Delete** /v1/notification/publisher/{notificationPublisherUuid} | Deletes a notification publisher and all related notification rules
[**DeleteNotificationRule**](DefaultApi.md#DeleteNotificationRule) | **Delete** /v1/notification/rule | Deletes a notification rule
[**GetAllNotificationPublishers**](DefaultApi.md#GetAllNotificationPublishers) | **Get** /v1/notification/publisher | Returns a list of all notification publishers
[**GetAllNotificationRules**](DefaultApi.md#GetAllNotificationRules) | **Get** /v1/notification/rule | Returns a list of all notification rules
[**RemoveProjectFromRule**](DefaultApi.md#RemoveProjectFromRule) | **Delete** /v1/notification/rule/{ruleUuid}/project/{projectUuid} | Removes a project from a notification rule
[**RemoveTeamFromRule**](DefaultApi.md#RemoveTeamFromRule) | **Delete** /v1/notification/rule/{ruleUuid}/team/{teamUuid} | Removes a team from a notification rule
[**RestoreDefaultTemplates**](DefaultApi.md#RestoreDefaultTemplates) | **Post** /v1/notification/publisher/restoreDefaultTemplates | Restore the default notification publisher templates using the ones in the solution classpath
[**TestSmtpPublisherConfig**](DefaultApi.md#TestSmtpPublisherConfig) | **Post** /v1/notification/publisher/test/smtp | Dispatches a SMTP notification test
[**UpdateNotificationPublisher**](DefaultApi.md#UpdateNotificationPublisher) | **Post** /v1/notification/publisher | Updates a notification publisher
[**UpdateNotificationRule**](DefaultApi.md#UpdateNotificationRule) | **Post** /v1/notification/rule | Updates a notification rule


# **AddProjectToRule**
> NotificationRule AddProjectToRule(ctx, ruleUuid, projectUuid)
Adds a project to a notification rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleUuid** | **string**| The UUID of the rule to add a project to | 
  **projectUuid** | **string**| The UUID of the project to add to the rule | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTeamToRule**
> NotificationRule AddTeamToRule(ctx, ruleUuid, teamUuid)
Adds a team to a notification rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleUuid** | **string**| The UUID of the rule to add a team to | 
  **teamUuid** | **string**| The UUID of the team to add to the rule | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNotificationPublisher**
> NotificationPublisher CreateNotificationPublisher(ctx, optional)
Creates a new notification publisher



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiCreateNotificationPublisherOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiCreateNotificationPublisherOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationPublisher**](NotificationPublisher.md)|  | 

### Return type

[**NotificationPublisher**](NotificationPublisher.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNotificationRule**
> NotificationRule CreateNotificationRule(ctx, optional)
Creates a new notification rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiCreateNotificationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiCreateNotificationRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationRule**](NotificationRule.md)|  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationPublisher**
> DeleteNotificationPublisher(ctx, notificationPublisherUuid)
Deletes a notification publisher and all related notification rules



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationPublisherUuid** | **string**| The UUID of the notification publisher to delete | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationRule**
> DeleteNotificationRule(ctx, optional)
Deletes a notification rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiDeleteNotificationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiDeleteNotificationRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationRule**](NotificationRule.md)|  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllNotificationPublishers**
> []NotificationPublisher GetAllNotificationPublishers(ctx, )
Returns a list of all notification publishers



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]NotificationPublisher**](NotificationPublisher.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllNotificationRules**
> []NotificationRule GetAllNotificationRules(ctx, )
Returns a list of all notification rules



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]NotificationRule**](NotificationRule.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveProjectFromRule**
> NotificationRule RemoveProjectFromRule(ctx, ruleUuid, projectUuid)
Removes a project from a notification rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleUuid** | **string**| The UUID of the rule to remove the project from | 
  **projectUuid** | **string**| The UUID of the project to remove from the rule | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTeamFromRule**
> NotificationRule RemoveTeamFromRule(ctx, ruleUuid, teamUuid)
Removes a team from a notification rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleUuid** | **string**| The UUID of the rule to remove the project from | 
  **teamUuid** | **string**| The UUID of the project to remove from the rule | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreDefaultTemplates**
> RestoreDefaultTemplates(ctx, )
Restore the default notification publisher templates using the ones in the solution classpath



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestSmtpPublisherConfig**
> TestSmtpPublisherConfig(ctx, optional)
Dispatches a SMTP notification test



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiTestSmtpPublisherConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTestSmtpPublisherConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destination** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationPublisher**
> NotificationRule UpdateNotificationPublisher(ctx, optional)
Updates a notification publisher



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiUpdateNotificationPublisherOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUpdateNotificationPublisherOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationPublisher**](NotificationPublisher.md)|  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationRule**
> NotificationRule UpdateNotificationRule(ctx, optional)
Updates a notification rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiUpdateNotificationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUpdateNotificationRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationRule**](NotificationRule.md)|  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

