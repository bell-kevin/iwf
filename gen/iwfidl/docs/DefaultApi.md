# \DefaultApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1WorkflowCancelPost**](DefaultApi.md#ApiV1WorkflowCancelPost) | **Post** /api/v1/workflow/cancel | cancel a workflow
[**ApiV1WorkflowGetPost**](DefaultApi.md#ApiV1WorkflowGetPost) | **Post** /api/v1/workflow/get | get a workflow&#39;s status and results(if completed &amp; requested)
[**ApiV1WorkflowGetWithWaitPost**](DefaultApi.md#ApiV1WorkflowGetWithWaitPost) | **Post** /api/v1/workflow/getWithWait | get a workflow&#39;s status and results(if completed &amp; requested), wait if the workflow is still running
[**ApiV1WorkflowQueryattributesGetPost**](DefaultApi.md#ApiV1WorkflowQueryattributesGetPost) | **Post** /api/v1/workflow/queryattributes/get | get workflow query attributes
[**ApiV1WorkflowResetPost**](DefaultApi.md#ApiV1WorkflowResetPost) | **Post** /api/v1/workflow/reset | reset a workflow
[**ApiV1WorkflowSearchPost**](DefaultApi.md#ApiV1WorkflowSearchPost) | **Post** /api/v1/workflow/search | search for workflows by a search attribute query
[**ApiV1WorkflowSearchattributesGetPost**](DefaultApi.md#ApiV1WorkflowSearchattributesGetPost) | **Post** /api/v1/workflow/searchattributes/get | get workflow search attributes
[**ApiV1WorkflowSignalPost**](DefaultApi.md#ApiV1WorkflowSignalPost) | **Post** /api/v1/workflow/signal | signal a workflow
[**ApiV1WorkflowStartPost**](DefaultApi.md#ApiV1WorkflowStartPost) | **Post** /api/v1/workflow/start | start a workflow
[**ApiV1WorkflowStateDecidePost**](DefaultApi.md#ApiV1WorkflowStateDecidePost) | **Post** /api/v1/workflowState/decide | for invoking WorkflowState.decide API
[**ApiV1WorkflowStateStartPost**](DefaultApi.md#ApiV1WorkflowStateStartPost) | **Post** /api/v1/workflowState/start | for invoking WorkflowState.start API



## ApiV1WorkflowCancelPost

> ApiV1WorkflowCancelPost(ctx).WorkflowCancelRequest(workflowCancelRequest).Execute()

cancel a workflow

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowCancelRequest := *openapiclient.NewWorkflowCancelRequest("WorkflowId_example") // WorkflowCancelRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowCancelPost(context.Background()).WorkflowCancelRequest(workflowCancelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowCancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowCancelRequest** | [**WorkflowCancelRequest**](WorkflowCancelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowGetPost

> WorkflowGetResponse ApiV1WorkflowGetPost(ctx).WorkflowGetRequest(workflowGetRequest).Execute()

get a workflow's status and results(if completed & requested)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowGetRequest := *openapiclient.NewWorkflowGetRequest("WorkflowId_example") // WorkflowGetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowGetPost(context.Background()).WorkflowGetRequest(workflowGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowGetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowGetPost`: WorkflowGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowGetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowGetRequest** | [**WorkflowGetRequest**](WorkflowGetRequest.md) |  | 

### Return type

[**WorkflowGetResponse**](WorkflowGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowGetWithWaitPost

> WorkflowGetResponse ApiV1WorkflowGetWithWaitPost(ctx).WorkflowGetRequest(workflowGetRequest).Execute()

get a workflow's status and results(if completed & requested), wait if the workflow is still running

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowGetRequest := *openapiclient.NewWorkflowGetRequest("WorkflowId_example") // WorkflowGetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowGetWithWaitPost(context.Background()).WorkflowGetRequest(workflowGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowGetWithWaitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowGetWithWaitPost`: WorkflowGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowGetWithWaitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowGetWithWaitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowGetRequest** | [**WorkflowGetRequest**](WorkflowGetRequest.md) |  | 

### Return type

[**WorkflowGetResponse**](WorkflowGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowQueryattributesGetPost

> WorkflowGetQueryAttributesResponse ApiV1WorkflowQueryattributesGetPost(ctx).WorkflowGetQueryAttributesRequest(workflowGetQueryAttributesRequest).Execute()

get workflow query attributes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowGetQueryAttributesRequest := *openapiclient.NewWorkflowGetQueryAttributesRequest("WorkflowId_example") // WorkflowGetQueryAttributesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowQueryattributesGetPost(context.Background()).WorkflowGetQueryAttributesRequest(workflowGetQueryAttributesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowQueryattributesGetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowQueryattributesGetPost`: WorkflowGetQueryAttributesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowQueryattributesGetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowQueryattributesGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowGetQueryAttributesRequest** | [**WorkflowGetQueryAttributesRequest**](WorkflowGetQueryAttributesRequest.md) |  | 

### Return type

[**WorkflowGetQueryAttributesResponse**](WorkflowGetQueryAttributesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowResetPost

> WorkflowResetResponse ApiV1WorkflowResetPost(ctx).WorkflowResetRequest(workflowResetRequest).Execute()

reset a workflow

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowResetRequest := *openapiclient.NewWorkflowResetRequest("WorkflowId_example", "ResetType_example") // WorkflowResetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowResetPost(context.Background()).WorkflowResetRequest(workflowResetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowResetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowResetPost`: WorkflowResetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowResetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowResetRequest** | [**WorkflowResetRequest**](WorkflowResetRequest.md) |  | 

### Return type

[**WorkflowResetResponse**](WorkflowResetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowSearchPost

> WorkflowSearchResponse ApiV1WorkflowSearchPost(ctx).WorkflowSearchRequest(workflowSearchRequest).Execute()

search for workflows by a search attribute query

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowSearchRequest := *openapiclient.NewWorkflowSearchRequest("Query_example") // WorkflowSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowSearchPost(context.Background()).WorkflowSearchRequest(workflowSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowSearchPost`: WorkflowSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowSearchRequest** | [**WorkflowSearchRequest**](WorkflowSearchRequest.md) |  | 

### Return type

[**WorkflowSearchResponse**](WorkflowSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowSearchattributesGetPost

> WorkflowGetSearchAttributesResponse ApiV1WorkflowSearchattributesGetPost(ctx).WorkflowGetSearchAttributesRequest(workflowGetSearchAttributesRequest).Execute()

get workflow search attributes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowGetSearchAttributesRequest := *openapiclient.NewWorkflowGetSearchAttributesRequest("WorkflowId_example") // WorkflowGetSearchAttributesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowSearchattributesGetPost(context.Background()).WorkflowGetSearchAttributesRequest(workflowGetSearchAttributesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowSearchattributesGetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowSearchattributesGetPost`: WorkflowGetSearchAttributesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowSearchattributesGetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowSearchattributesGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowGetSearchAttributesRequest** | [**WorkflowGetSearchAttributesRequest**](WorkflowGetSearchAttributesRequest.md) |  | 

### Return type

[**WorkflowGetSearchAttributesResponse**](WorkflowGetSearchAttributesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowSignalPost

> ApiV1WorkflowSignalPost(ctx).WorkflowSignalRequest(workflowSignalRequest).Execute()

signal a workflow

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowSignalRequest := *openapiclient.NewWorkflowSignalRequest("WorkflowId_example", "SignalChannelName_example") // WorkflowSignalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowSignalPost(context.Background()).WorkflowSignalRequest(workflowSignalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowSignalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowSignalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowSignalRequest** | [**WorkflowSignalRequest**](WorkflowSignalRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowStartPost

> WorkflowStartResponse ApiV1WorkflowStartPost(ctx).WorkflowStartRequest(workflowStartRequest).Execute()

start a workflow

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowStartRequest := *openapiclient.NewWorkflowStartRequest("WorkflowId_example", "IwfWorkflowType_example", int32(123), "IwfWorkerUrl_example", "StartStateId_example") // WorkflowStartRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowStartPost(context.Background()).WorkflowStartRequest(workflowStartRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowStartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowStartPost`: WorkflowStartResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowStartRequest** | [**WorkflowStartRequest**](WorkflowStartRequest.md) |  | 

### Return type

[**WorkflowStartResponse**](WorkflowStartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowStateDecidePost

> WorkflowStateDecideResponse ApiV1WorkflowStateDecidePost(ctx).WorkflowStateDecideRequest(workflowStateDecideRequest).Execute()

for invoking WorkflowState.decide API

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowStateDecideRequest := *openapiclient.NewWorkflowStateDecideRequest(*openapiclient.NewContext("WorkflowId_example", "WorkflowRunId_example", int64(123), "StateExecutionId_example"), "WorkflowType_example", "WorkflowStateId_example") // WorkflowStateDecideRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowStateDecidePost(context.Background()).WorkflowStateDecideRequest(workflowStateDecideRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowStateDecidePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowStateDecidePost`: WorkflowStateDecideResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowStateDecidePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowStateDecidePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowStateDecideRequest** | [**WorkflowStateDecideRequest**](WorkflowStateDecideRequest.md) |  | 

### Return type

[**WorkflowStateDecideResponse**](WorkflowStateDecideResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowStateStartPost

> WorkflowStateStartResponse ApiV1WorkflowStateStartPost(ctx).WorkflowStateStartRequest(workflowStateStartRequest).Execute()

for invoking WorkflowState.start API

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowStateStartRequest := *openapiclient.NewWorkflowStateStartRequest(*openapiclient.NewContext("WorkflowId_example", "WorkflowRunId_example", int64(123), "StateExecutionId_example"), "WorkflowType_example", "WorkflowStateId_example") // WorkflowStateStartRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowStateStartPost(context.Background()).WorkflowStateStartRequest(workflowStateStartRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowStateStartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowStateStartPost`: WorkflowStateStartResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowStateStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowStateStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowStateStartRequest** | [**WorkflowStateStartRequest**](WorkflowStateStartRequest.md) |  | 

### Return type

[**WorkflowStateStartResponse**](WorkflowStateStartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

