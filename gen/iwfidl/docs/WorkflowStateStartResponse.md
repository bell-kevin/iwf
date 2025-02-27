# WorkflowStateStartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpsertSearchAttributes** | Pointer to [**[]SearchAttribute**](SearchAttribute.md) |  | [optional] 
**UpsertQueryAttributes** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**CommandRequest** | Pointer to [**CommandRequest**](CommandRequest.md) |  | [optional] 
**UpsertStateLocalAttributes** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**RecordEvents** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**PublishToInterStateChannel** | Pointer to [**[]InterStateChannelPublishing**](InterStateChannelPublishing.md) |  | [optional] 

## Methods

### NewWorkflowStateStartResponse

`func NewWorkflowStateStartResponse() *WorkflowStateStartResponse`

NewWorkflowStateStartResponse instantiates a new WorkflowStateStartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStateStartResponseWithDefaults

`func NewWorkflowStateStartResponseWithDefaults() *WorkflowStateStartResponse`

NewWorkflowStateStartResponseWithDefaults instantiates a new WorkflowStateStartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpsertSearchAttributes

`func (o *WorkflowStateStartResponse) GetUpsertSearchAttributes() []SearchAttribute`

GetUpsertSearchAttributes returns the UpsertSearchAttributes field if non-nil, zero value otherwise.

### GetUpsertSearchAttributesOk

`func (o *WorkflowStateStartResponse) GetUpsertSearchAttributesOk() (*[]SearchAttribute, bool)`

GetUpsertSearchAttributesOk returns a tuple with the UpsertSearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertSearchAttributes

`func (o *WorkflowStateStartResponse) SetUpsertSearchAttributes(v []SearchAttribute)`

SetUpsertSearchAttributes sets UpsertSearchAttributes field to given value.

### HasUpsertSearchAttributes

`func (o *WorkflowStateStartResponse) HasUpsertSearchAttributes() bool`

HasUpsertSearchAttributes returns a boolean if a field has been set.

### GetUpsertQueryAttributes

`func (o *WorkflowStateStartResponse) GetUpsertQueryAttributes() []KeyValue`

GetUpsertQueryAttributes returns the UpsertQueryAttributes field if non-nil, zero value otherwise.

### GetUpsertQueryAttributesOk

`func (o *WorkflowStateStartResponse) GetUpsertQueryAttributesOk() (*[]KeyValue, bool)`

GetUpsertQueryAttributesOk returns a tuple with the UpsertQueryAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertQueryAttributes

`func (o *WorkflowStateStartResponse) SetUpsertQueryAttributes(v []KeyValue)`

SetUpsertQueryAttributes sets UpsertQueryAttributes field to given value.

### HasUpsertQueryAttributes

`func (o *WorkflowStateStartResponse) HasUpsertQueryAttributes() bool`

HasUpsertQueryAttributes returns a boolean if a field has been set.

### GetCommandRequest

`func (o *WorkflowStateStartResponse) GetCommandRequest() CommandRequest`

GetCommandRequest returns the CommandRequest field if non-nil, zero value otherwise.

### GetCommandRequestOk

`func (o *WorkflowStateStartResponse) GetCommandRequestOk() (*CommandRequest, bool)`

GetCommandRequestOk returns a tuple with the CommandRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandRequest

`func (o *WorkflowStateStartResponse) SetCommandRequest(v CommandRequest)`

SetCommandRequest sets CommandRequest field to given value.

### HasCommandRequest

`func (o *WorkflowStateStartResponse) HasCommandRequest() bool`

HasCommandRequest returns a boolean if a field has been set.

### GetUpsertStateLocalAttributes

`func (o *WorkflowStateStartResponse) GetUpsertStateLocalAttributes() []KeyValue`

GetUpsertStateLocalAttributes returns the UpsertStateLocalAttributes field if non-nil, zero value otherwise.

### GetUpsertStateLocalAttributesOk

`func (o *WorkflowStateStartResponse) GetUpsertStateLocalAttributesOk() (*[]KeyValue, bool)`

GetUpsertStateLocalAttributesOk returns a tuple with the UpsertStateLocalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertStateLocalAttributes

`func (o *WorkflowStateStartResponse) SetUpsertStateLocalAttributes(v []KeyValue)`

SetUpsertStateLocalAttributes sets UpsertStateLocalAttributes field to given value.

### HasUpsertStateLocalAttributes

`func (o *WorkflowStateStartResponse) HasUpsertStateLocalAttributes() bool`

HasUpsertStateLocalAttributes returns a boolean if a field has been set.

### GetRecordEvents

`func (o *WorkflowStateStartResponse) GetRecordEvents() []KeyValue`

GetRecordEvents returns the RecordEvents field if non-nil, zero value otherwise.

### GetRecordEventsOk

`func (o *WorkflowStateStartResponse) GetRecordEventsOk() (*[]KeyValue, bool)`

GetRecordEventsOk returns a tuple with the RecordEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordEvents

`func (o *WorkflowStateStartResponse) SetRecordEvents(v []KeyValue)`

SetRecordEvents sets RecordEvents field to given value.

### HasRecordEvents

`func (o *WorkflowStateStartResponse) HasRecordEvents() bool`

HasRecordEvents returns a boolean if a field has been set.

### GetPublishToInterStateChannel

`func (o *WorkflowStateStartResponse) GetPublishToInterStateChannel() []InterStateChannelPublishing`

GetPublishToInterStateChannel returns the PublishToInterStateChannel field if non-nil, zero value otherwise.

### GetPublishToInterStateChannelOk

`func (o *WorkflowStateStartResponse) GetPublishToInterStateChannelOk() (*[]InterStateChannelPublishing, bool)`

GetPublishToInterStateChannelOk returns a tuple with the PublishToInterStateChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishToInterStateChannel

`func (o *WorkflowStateStartResponse) SetPublishToInterStateChannel(v []InterStateChannelPublishing)`

SetPublishToInterStateChannel sets PublishToInterStateChannel field to given value.

### HasPublishToInterStateChannel

`func (o *WorkflowStateStartResponse) HasPublishToInterStateChannel() bool`

HasPublishToInterStateChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


