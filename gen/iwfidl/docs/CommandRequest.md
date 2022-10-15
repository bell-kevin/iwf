# CommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeciderTriggerType** | **string** |  | 
**ActivityCommands** | Pointer to [**[]ActivityCommand**](ActivityCommand.md) |  | [optional] 
**TimerCommands** | Pointer to [**[]TimerCommand**](TimerCommand.md) |  | [optional] 
**SignalCommands** | Pointer to [**[]SignalCommand**](SignalCommand.md) |  | [optional] 
**OnQueryAttributeChangeCommands** | Pointer to [**[]OnQueryAttributeChangeCommand**](OnQueryAttributeChangeCommand.md) |  | [optional] 

## Methods

### NewCommandRequest

`func NewCommandRequest(deciderTriggerType string, ) *CommandRequest`

NewCommandRequest instantiates a new CommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandRequestWithDefaults

`func NewCommandRequestWithDefaults() *CommandRequest`

NewCommandRequestWithDefaults instantiates a new CommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeciderTriggerType

`func (o *CommandRequest) GetDeciderTriggerType() string`

GetDeciderTriggerType returns the DeciderTriggerType field if non-nil, zero value otherwise.

### GetDeciderTriggerTypeOk

`func (o *CommandRequest) GetDeciderTriggerTypeOk() (*string, bool)`

GetDeciderTriggerTypeOk returns a tuple with the DeciderTriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeciderTriggerType

`func (o *CommandRequest) SetDeciderTriggerType(v string)`

SetDeciderTriggerType sets DeciderTriggerType field to given value.


### GetActivityCommands

`func (o *CommandRequest) GetActivityCommands() []ActivityCommand`

GetActivityCommands returns the ActivityCommands field if non-nil, zero value otherwise.

### GetActivityCommandsOk

`func (o *CommandRequest) GetActivityCommandsOk() (*[]ActivityCommand, bool)`

GetActivityCommandsOk returns a tuple with the ActivityCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCommands

`func (o *CommandRequest) SetActivityCommands(v []ActivityCommand)`

SetActivityCommands sets ActivityCommands field to given value.

### HasActivityCommands

`func (o *CommandRequest) HasActivityCommands() bool`

HasActivityCommands returns a boolean if a field has been set.

### GetTimerCommands

`func (o *CommandRequest) GetTimerCommands() []TimerCommand`

GetTimerCommands returns the TimerCommands field if non-nil, zero value otherwise.

### GetTimerCommandsOk

`func (o *CommandRequest) GetTimerCommandsOk() (*[]TimerCommand, bool)`

GetTimerCommandsOk returns a tuple with the TimerCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerCommands

`func (o *CommandRequest) SetTimerCommands(v []TimerCommand)`

SetTimerCommands sets TimerCommands field to given value.

### HasTimerCommands

`func (o *CommandRequest) HasTimerCommands() bool`

HasTimerCommands returns a boolean if a field has been set.

### GetSignalCommands

`func (o *CommandRequest) GetSignalCommands() []SignalCommand`

GetSignalCommands returns the SignalCommands field if non-nil, zero value otherwise.

### GetSignalCommandsOk

`func (o *CommandRequest) GetSignalCommandsOk() (*[]SignalCommand, bool)`

GetSignalCommandsOk returns a tuple with the SignalCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalCommands

`func (o *CommandRequest) SetSignalCommands(v []SignalCommand)`

SetSignalCommands sets SignalCommands field to given value.

### HasSignalCommands

`func (o *CommandRequest) HasSignalCommands() bool`

HasSignalCommands returns a boolean if a field has been set.

### GetOnQueryAttributeChangeCommands

`func (o *CommandRequest) GetOnQueryAttributeChangeCommands() []OnQueryAttributeChangeCommand`

GetOnQueryAttributeChangeCommands returns the OnQueryAttributeChangeCommands field if non-nil, zero value otherwise.

### GetOnQueryAttributeChangeCommandsOk

`func (o *CommandRequest) GetOnQueryAttributeChangeCommandsOk() (*[]OnQueryAttributeChangeCommand, bool)`

GetOnQueryAttributeChangeCommandsOk returns a tuple with the OnQueryAttributeChangeCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnQueryAttributeChangeCommands

`func (o *CommandRequest) SetOnQueryAttributeChangeCommands(v []OnQueryAttributeChangeCommand)`

SetOnQueryAttributeChangeCommands sets OnQueryAttributeChangeCommands field to given value.

### HasOnQueryAttributeChangeCommands

`func (o *CommandRequest) HasOnQueryAttributeChangeCommands() bool`

HasOnQueryAttributeChangeCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


