/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// WorkflowStateDecideResponse struct for WorkflowStateDecideResponse
type WorkflowStateDecideResponse struct {
	StateDecision *StateDecision `json:"stateDecision,omitempty"`
	UpsertSearchAttributes []SearchAttribute `json:"upsertSearchAttributes,omitempty"`
	UpsertQueryAttributes []KeyValue `json:"upsertQueryAttributes,omitempty"`
	RecordEvents []KeyValue `json:"recordEvents,omitempty"`
	UpsertStateLocalAttributes []KeyValue `json:"upsertStateLocalAttributes,omitempty"`
	PublishToInterStateChannel []InterStateChannelPublishing `json:"publishToInterStateChannel,omitempty"`
}

// NewWorkflowStateDecideResponse instantiates a new WorkflowStateDecideResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateDecideResponse() *WorkflowStateDecideResponse {
	this := WorkflowStateDecideResponse{}
	return &this
}

// NewWorkflowStateDecideResponseWithDefaults instantiates a new WorkflowStateDecideResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateDecideResponseWithDefaults() *WorkflowStateDecideResponse {
	this := WorkflowStateDecideResponse{}
	return &this
}

// GetStateDecision returns the StateDecision field value if set, zero value otherwise.
func (o *WorkflowStateDecideResponse) GetStateDecision() StateDecision {
	if o == nil || o.StateDecision == nil {
		var ret StateDecision
		return ret
	}
	return *o.StateDecision
}

// GetStateDecisionOk returns a tuple with the StateDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateDecideResponse) GetStateDecisionOk() (*StateDecision, bool) {
	if o == nil || o.StateDecision == nil {
		return nil, false
	}
	return o.StateDecision, true
}

// HasStateDecision returns a boolean if a field has been set.
func (o *WorkflowStateDecideResponse) HasStateDecision() bool {
	if o != nil && o.StateDecision != nil {
		return true
	}

	return false
}

// SetStateDecision gets a reference to the given StateDecision and assigns it to the StateDecision field.
func (o *WorkflowStateDecideResponse) SetStateDecision(v StateDecision) {
	o.StateDecision = &v
}

// GetUpsertSearchAttributes returns the UpsertSearchAttributes field value if set, zero value otherwise.
func (o *WorkflowStateDecideResponse) GetUpsertSearchAttributes() []SearchAttribute {
	if o == nil || o.UpsertSearchAttributes == nil {
		var ret []SearchAttribute
		return ret
	}
	return o.UpsertSearchAttributes
}

// GetUpsertSearchAttributesOk returns a tuple with the UpsertSearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateDecideResponse) GetUpsertSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || o.UpsertSearchAttributes == nil {
		return nil, false
	}
	return o.UpsertSearchAttributes, true
}

// HasUpsertSearchAttributes returns a boolean if a field has been set.
func (o *WorkflowStateDecideResponse) HasUpsertSearchAttributes() bool {
	if o != nil && o.UpsertSearchAttributes != nil {
		return true
	}

	return false
}

// SetUpsertSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the UpsertSearchAttributes field.
func (o *WorkflowStateDecideResponse) SetUpsertSearchAttributes(v []SearchAttribute) {
	o.UpsertSearchAttributes = v
}

// GetUpsertQueryAttributes returns the UpsertQueryAttributes field value if set, zero value otherwise.
func (o *WorkflowStateDecideResponse) GetUpsertQueryAttributes() []KeyValue {
	if o == nil || o.UpsertQueryAttributes == nil {
		var ret []KeyValue
		return ret
	}
	return o.UpsertQueryAttributes
}

// GetUpsertQueryAttributesOk returns a tuple with the UpsertQueryAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateDecideResponse) GetUpsertQueryAttributesOk() ([]KeyValue, bool) {
	if o == nil || o.UpsertQueryAttributes == nil {
		return nil, false
	}
	return o.UpsertQueryAttributes, true
}

// HasUpsertQueryAttributes returns a boolean if a field has been set.
func (o *WorkflowStateDecideResponse) HasUpsertQueryAttributes() bool {
	if o != nil && o.UpsertQueryAttributes != nil {
		return true
	}

	return false
}

// SetUpsertQueryAttributes gets a reference to the given []KeyValue and assigns it to the UpsertQueryAttributes field.
func (o *WorkflowStateDecideResponse) SetUpsertQueryAttributes(v []KeyValue) {
	o.UpsertQueryAttributes = v
}

// GetRecordEvents returns the RecordEvents field value if set, zero value otherwise.
func (o *WorkflowStateDecideResponse) GetRecordEvents() []KeyValue {
	if o == nil || o.RecordEvents == nil {
		var ret []KeyValue
		return ret
	}
	return o.RecordEvents
}

// GetRecordEventsOk returns a tuple with the RecordEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateDecideResponse) GetRecordEventsOk() ([]KeyValue, bool) {
	if o == nil || o.RecordEvents == nil {
		return nil, false
	}
	return o.RecordEvents, true
}

// HasRecordEvents returns a boolean if a field has been set.
func (o *WorkflowStateDecideResponse) HasRecordEvents() bool {
	if o != nil && o.RecordEvents != nil {
		return true
	}

	return false
}

// SetRecordEvents gets a reference to the given []KeyValue and assigns it to the RecordEvents field.
func (o *WorkflowStateDecideResponse) SetRecordEvents(v []KeyValue) {
	o.RecordEvents = v
}

// GetUpsertStateLocalAttributes returns the UpsertStateLocalAttributes field value if set, zero value otherwise.
func (o *WorkflowStateDecideResponse) GetUpsertStateLocalAttributes() []KeyValue {
	if o == nil || o.UpsertStateLocalAttributes == nil {
		var ret []KeyValue
		return ret
	}
	return o.UpsertStateLocalAttributes
}

// GetUpsertStateLocalAttributesOk returns a tuple with the UpsertStateLocalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateDecideResponse) GetUpsertStateLocalAttributesOk() ([]KeyValue, bool) {
	if o == nil || o.UpsertStateLocalAttributes == nil {
		return nil, false
	}
	return o.UpsertStateLocalAttributes, true
}

// HasUpsertStateLocalAttributes returns a boolean if a field has been set.
func (o *WorkflowStateDecideResponse) HasUpsertStateLocalAttributes() bool {
	if o != nil && o.UpsertStateLocalAttributes != nil {
		return true
	}

	return false
}

// SetUpsertStateLocalAttributes gets a reference to the given []KeyValue and assigns it to the UpsertStateLocalAttributes field.
func (o *WorkflowStateDecideResponse) SetUpsertStateLocalAttributes(v []KeyValue) {
	o.UpsertStateLocalAttributes = v
}

// GetPublishToInterStateChannel returns the PublishToInterStateChannel field value if set, zero value otherwise.
func (o *WorkflowStateDecideResponse) GetPublishToInterStateChannel() []InterStateChannelPublishing {
	if o == nil || o.PublishToInterStateChannel == nil {
		var ret []InterStateChannelPublishing
		return ret
	}
	return o.PublishToInterStateChannel
}

// GetPublishToInterStateChannelOk returns a tuple with the PublishToInterStateChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateDecideResponse) GetPublishToInterStateChannelOk() ([]InterStateChannelPublishing, bool) {
	if o == nil || o.PublishToInterStateChannel == nil {
		return nil, false
	}
	return o.PublishToInterStateChannel, true
}

// HasPublishToInterStateChannel returns a boolean if a field has been set.
func (o *WorkflowStateDecideResponse) HasPublishToInterStateChannel() bool {
	if o != nil && o.PublishToInterStateChannel != nil {
		return true
	}

	return false
}

// SetPublishToInterStateChannel gets a reference to the given []InterStateChannelPublishing and assigns it to the PublishToInterStateChannel field.
func (o *WorkflowStateDecideResponse) SetPublishToInterStateChannel(v []InterStateChannelPublishing) {
	o.PublishToInterStateChannel = v
}

func (o WorkflowStateDecideResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StateDecision != nil {
		toSerialize["stateDecision"] = o.StateDecision
	}
	if o.UpsertSearchAttributes != nil {
		toSerialize["upsertSearchAttributes"] = o.UpsertSearchAttributes
	}
	if o.UpsertQueryAttributes != nil {
		toSerialize["upsertQueryAttributes"] = o.UpsertQueryAttributes
	}
	if o.RecordEvents != nil {
		toSerialize["recordEvents"] = o.RecordEvents
	}
	if o.UpsertStateLocalAttributes != nil {
		toSerialize["upsertStateLocalAttributes"] = o.UpsertStateLocalAttributes
	}
	if o.PublishToInterStateChannel != nil {
		toSerialize["publishToInterStateChannel"] = o.PublishToInterStateChannel
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowStateDecideResponse struct {
	value *WorkflowStateDecideResponse
	isSet bool
}

func (v NullableWorkflowStateDecideResponse) Get() *WorkflowStateDecideResponse {
	return v.value
}

func (v *NullableWorkflowStateDecideResponse) Set(val *WorkflowStateDecideResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateDecideResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateDecideResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateDecideResponse(val *WorkflowStateDecideResponse) *NullableWorkflowStateDecideResponse {
	return &NullableWorkflowStateDecideResponse{value: val, isSet: true}
}

func (v NullableWorkflowStateDecideResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateDecideResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


