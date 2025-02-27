# Go API client for iwfidl

This APIs for iwf SDKs to operate workflows

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import iwfidl "github.com/indeedeng/iwf-idl"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), iwfidl.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), iwfidl.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), iwfidl.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), iwfidl.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://petstore.swagger.io/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ApiV1WorkflowCancelPost**](docs/DefaultApi.md#apiv1workflowcancelpost) | **Post** /api/v1/workflow/cancel | cancel a workflow
*DefaultApi* | [**ApiV1WorkflowGetPost**](docs/DefaultApi.md#apiv1workflowgetpost) | **Post** /api/v1/workflow/get | get a workflow&#39;s status and results(if completed &amp; requested)
*DefaultApi* | [**ApiV1WorkflowGetWithWaitPost**](docs/DefaultApi.md#apiv1workflowgetwithwaitpost) | **Post** /api/v1/workflow/getWithWait | get a workflow&#39;s status and results(if completed &amp; requested), wait if the workflow is still running
*DefaultApi* | [**ApiV1WorkflowQueryattributesGetPost**](docs/DefaultApi.md#apiv1workflowqueryattributesgetpost) | **Post** /api/v1/workflow/queryattributes/get | get workflow query attributes
*DefaultApi* | [**ApiV1WorkflowResetPost**](docs/DefaultApi.md#apiv1workflowresetpost) | **Post** /api/v1/workflow/reset | reset a workflow
*DefaultApi* | [**ApiV1WorkflowSearchPost**](docs/DefaultApi.md#apiv1workflowsearchpost) | **Post** /api/v1/workflow/search | search for workflows by a search attribute query
*DefaultApi* | [**ApiV1WorkflowSearchattributesGetPost**](docs/DefaultApi.md#apiv1workflowsearchattributesgetpost) | **Post** /api/v1/workflow/searchattributes/get | get workflow search attributes
*DefaultApi* | [**ApiV1WorkflowSignalPost**](docs/DefaultApi.md#apiv1workflowsignalpost) | **Post** /api/v1/workflow/signal | signal a workflow
*DefaultApi* | [**ApiV1WorkflowStartPost**](docs/DefaultApi.md#apiv1workflowstartpost) | **Post** /api/v1/workflow/start | start a workflow
*DefaultApi* | [**ApiV1WorkflowStateDecidePost**](docs/DefaultApi.md#apiv1workflowstatedecidepost) | **Post** /api/v1/workflowState/decide | for invoking WorkflowState.decide API
*DefaultApi* | [**ApiV1WorkflowStateStartPost**](docs/DefaultApi.md#apiv1workflowstatestartpost) | **Post** /api/v1/workflowState/start | for invoking WorkflowState.start API


## Documentation For Models

 - [AttributesLoadingPolicy](docs/AttributesLoadingPolicy.md)
 - [CommandCarryOverPolicy](docs/CommandCarryOverPolicy.md)
 - [CommandRequest](docs/CommandRequest.md)
 - [CommandResults](docs/CommandResults.md)
 - [Context](docs/Context.md)
 - [EncodedObject](docs/EncodedObject.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [InterStateChannelCommand](docs/InterStateChannelCommand.md)
 - [InterStateChannelPublishing](docs/InterStateChannelPublishing.md)
 - [InterStateChannelResult](docs/InterStateChannelResult.md)
 - [KeyValue](docs/KeyValue.md)
 - [SearchAttribute](docs/SearchAttribute.md)
 - [SearchAttributeKeyAndType](docs/SearchAttributeKeyAndType.md)
 - [SignalCommand](docs/SignalCommand.md)
 - [SignalResult](docs/SignalResult.md)
 - [StateCompletionOutput](docs/StateCompletionOutput.md)
 - [StateDecision](docs/StateDecision.md)
 - [StateMovement](docs/StateMovement.md)
 - [TimerCommand](docs/TimerCommand.md)
 - [TimerResult](docs/TimerResult.md)
 - [WorkflowCancelRequest](docs/WorkflowCancelRequest.md)
 - [WorkflowGetQueryAttributesRequest](docs/WorkflowGetQueryAttributesRequest.md)
 - [WorkflowGetQueryAttributesResponse](docs/WorkflowGetQueryAttributesResponse.md)
 - [WorkflowGetRequest](docs/WorkflowGetRequest.md)
 - [WorkflowGetResponse](docs/WorkflowGetResponse.md)
 - [WorkflowGetSearchAttributesRequest](docs/WorkflowGetSearchAttributesRequest.md)
 - [WorkflowGetSearchAttributesResponse](docs/WorkflowGetSearchAttributesResponse.md)
 - [WorkflowResetRequest](docs/WorkflowResetRequest.md)
 - [WorkflowResetResponse](docs/WorkflowResetResponse.md)
 - [WorkflowSearchRequest](docs/WorkflowSearchRequest.md)
 - [WorkflowSearchResponse](docs/WorkflowSearchResponse.md)
 - [WorkflowSearchResponseEntry](docs/WorkflowSearchResponseEntry.md)
 - [WorkflowSignalRequest](docs/WorkflowSignalRequest.md)
 - [WorkflowStartRequest](docs/WorkflowStartRequest.md)
 - [WorkflowStartResponse](docs/WorkflowStartResponse.md)
 - [WorkflowStateDecideRequest](docs/WorkflowStateDecideRequest.md)
 - [WorkflowStateDecideResponse](docs/WorkflowStateDecideResponse.md)
 - [WorkflowStateOptions](docs/WorkflowStateOptions.md)
 - [WorkflowStateStartRequest](docs/WorkflowStateStartRequest.md)
 - [WorkflowStateStartResponse](docs/WorkflowStateStartResponse.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



