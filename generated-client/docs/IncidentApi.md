# IncidentApi

All URIs are relative to *https://incident.deltalab.dev/v0*

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**incidentHelpPost**](IncidentApi.md#incidentHelpPost) | **POST** incident/help | Help |
| [**incidentHelpQueryPost**](IncidentApi.md#incidentHelpQueryPost) | **POST** incident/help/query | Query help |
| [**incidentImageIncidentIdGet**](IncidentApi.md#incidentImageIncidentIdGet) | **GET** incident/image/{incident_id} | Query incident image |
| [**incidentQueryPost**](IncidentApi.md#incidentQueryPost) | **POST** incident/query | Query incidents |
| [**incidentReportDelete**](IncidentApi.md#incidentReportDelete) | **DELETE** incident/report | Delete an incident report |
| [**incidentReportPost**](IncidentApi.md#incidentReportPost) | **POST** incident/report | Report an incident |
| [**incidentVotePost**](IncidentApi.md#incidentVotePost) | **POST** incident/vote | Vote an incident |



Help

Help

### Example
```kotlin
// Import classes:
//import org.openapitools.client.*
//import org.openapitools.client.infrastructure.*
//import org.openapitools.client.models.*

val apiClient = ApiClient()
val webService = apiClient.createWebservice(IncidentApi::class.java)
val request : IncidentHelpRequest =  // IncidentHelpRequest | Help request

webService.incidentHelpPost(request)
```

### Parameters
| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **request** | [**IncidentHelpRequest**](IncidentHelpRequest.md)| Help request | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*


Query help

Query help

### Example
```kotlin
// Import classes:
//import org.openapitools.client.*
//import org.openapitools.client.infrastructure.*
//import org.openapitools.client.models.*

val apiClient = ApiClient()
val webService = apiClient.createWebservice(IncidentApi::class.java)
val request : IncidentQueryHelpRequest =  // IncidentQueryHelpRequest | Query help request

val result : kotlin.collections.List<IncidentQueryHelpResponse> = webService.incidentHelpQueryPost(request)
```

### Parameters
| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **request** | [**IncidentQueryHelpRequest**](IncidentQueryHelpRequest.md)| Query help request | |

### Return type

[**kotlin.collections.List&lt;IncidentQueryHelpResponse&gt;**](IncidentQueryHelpResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


Query incident image

Query incident image

### Example
```kotlin
// Import classes:
//import org.openapitools.client.*
//import org.openapitools.client.infrastructure.*
//import org.openapitools.client.models.*

val apiClient = ApiClient()
val webService = apiClient.createWebservice(IncidentApi::class.java)
val incidentId : kotlin.String = incidentId_example // kotlin.String | Incident ID

val result : kotlin.String = webService.incidentImageIncidentIdGet(incidentId)
```

### Parameters
| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **incidentId** | **kotlin.String**| Incident ID | |

### Return type

**kotlin.String**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*


Query incidents

Query incidents

### Example
```kotlin
// Import classes:
//import org.openapitools.client.*
//import org.openapitools.client.infrastructure.*
//import org.openapitools.client.models.*

val apiClient = ApiClient()
val webService = apiClient.createWebservice(IncidentApi::class.java)
val request : IncidentQueryRequest =  // IncidentQueryRequest | Query request

val result : kotlin.collections.List<IncidentQueryResponse> = webService.incidentQueryPost(request)
```

### Parameters
| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **request** | [**IncidentQueryRequest**](IncidentQueryRequest.md)| Query request | |

### Return type

[**kotlin.collections.List&lt;IncidentQueryResponse&gt;**](IncidentQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


Delete an incident report

Delete an incident report

### Example
```kotlin
// Import classes:
//import org.openapitools.client.*
//import org.openapitools.client.infrastructure.*
//import org.openapitools.client.models.*

val apiClient = ApiClient()
val webService = apiClient.createWebservice(IncidentApi::class.java)
val request : IncidentDeleteReportRequest =  // IncidentDeleteReportRequest | Delete report request

webService.incidentReportDelete(request)
```

### Parameters
| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **request** | [**IncidentDeleteReportRequest**](IncidentDeleteReportRequest.md)| Delete report request | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*


Report an incident

Report an incident

### Example
```kotlin
// Import classes:
//import org.openapitools.client.*
//import org.openapitools.client.infrastructure.*
//import org.openapitools.client.models.*

val apiClient = ApiClient()
val webService = apiClient.createWebservice(IncidentApi::class.java)
val request : IncidentReportRequest =  // IncidentReportRequest | Report request

val result : IncidentReportResponse = webService.incidentReportPost(request)
```

### Parameters
| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **request** | [**IncidentReportRequest**](IncidentReportRequest.md)| Report request | |

### Return type

[**IncidentReportResponse**](IncidentReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


Vote an incident

Vote an incident

### Example
```kotlin
// Import classes:
//import org.openapitools.client.*
//import org.openapitools.client.infrastructure.*
//import org.openapitools.client.models.*

val apiClient = ApiClient()
val webService = apiClient.createWebservice(IncidentApi::class.java)
val request : IncidentVoteRequest =  // IncidentVoteRequest | Vote request

webService.incidentVotePost(request)
```

### Parameters
| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **request** | [**IncidentVoteRequest**](IncidentVoteRequest.md)| Vote request | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

