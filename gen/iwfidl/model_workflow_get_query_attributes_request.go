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

// WorkflowGetQueryAttributesRequest struct for WorkflowGetQueryAttributesRequest
type WorkflowGetQueryAttributesRequest struct {
	WorkflowId string `json:"workflowId"`
	WorkflowRunId *string `json:"workflowRunId,omitempty"`
	AttributeKeys []string `json:"attributeKeys,omitempty"`
}

// NewWorkflowGetQueryAttributesRequest instantiates a new WorkflowGetQueryAttributesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowGetQueryAttributesRequest(workflowId string) *WorkflowGetQueryAttributesRequest {
	this := WorkflowGetQueryAttributesRequest{}
	this.WorkflowId = workflowId
	return &this
}

// NewWorkflowGetQueryAttributesRequestWithDefaults instantiates a new WorkflowGetQueryAttributesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowGetQueryAttributesRequestWithDefaults() *WorkflowGetQueryAttributesRequest {
	this := WorkflowGetQueryAttributesRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowGetQueryAttributesRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowGetQueryAttributesRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowGetQueryAttributesRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *WorkflowGetQueryAttributesRequest) GetWorkflowRunId() string {
	if o == nil || o.WorkflowRunId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetQueryAttributesRequest) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || o.WorkflowRunId == nil {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *WorkflowGetQueryAttributesRequest) HasWorkflowRunId() bool {
	if o != nil && o.WorkflowRunId != nil {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *WorkflowGetQueryAttributesRequest) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

// GetAttributeKeys returns the AttributeKeys field value if set, zero value otherwise.
func (o *WorkflowGetQueryAttributesRequest) GetAttributeKeys() []string {
	if o == nil || o.AttributeKeys == nil {
		var ret []string
		return ret
	}
	return o.AttributeKeys
}

// GetAttributeKeysOk returns a tuple with the AttributeKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetQueryAttributesRequest) GetAttributeKeysOk() ([]string, bool) {
	if o == nil || o.AttributeKeys == nil {
		return nil, false
	}
	return o.AttributeKeys, true
}

// HasAttributeKeys returns a boolean if a field has been set.
func (o *WorkflowGetQueryAttributesRequest) HasAttributeKeys() bool {
	if o != nil && o.AttributeKeys != nil {
		return true
	}

	return false
}

// SetAttributeKeys gets a reference to the given []string and assigns it to the AttributeKeys field.
func (o *WorkflowGetQueryAttributesRequest) SetAttributeKeys(v []string) {
	o.AttributeKeys = v
}

func (o WorkflowGetQueryAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if o.WorkflowRunId != nil {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	if o.AttributeKeys != nil {
		toSerialize["attributeKeys"] = o.AttributeKeys
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowGetQueryAttributesRequest struct {
	value *WorkflowGetQueryAttributesRequest
	isSet bool
}

func (v NullableWorkflowGetQueryAttributesRequest) Get() *WorkflowGetQueryAttributesRequest {
	return v.value
}

func (v *NullableWorkflowGetQueryAttributesRequest) Set(val *WorkflowGetQueryAttributesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowGetQueryAttributesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowGetQueryAttributesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowGetQueryAttributesRequest(val *WorkflowGetQueryAttributesRequest) *NullableWorkflowGetQueryAttributesRequest {
	return &NullableWorkflowGetQueryAttributesRequest{value: val, isSet: true}
}

func (v NullableWorkflowGetQueryAttributesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowGetQueryAttributesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


