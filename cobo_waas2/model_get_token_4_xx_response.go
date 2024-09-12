/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetToken4XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetToken4XXResponse{}

// GetToken4XXResponse The response of a failed request.
type GetToken4XXResponse struct {
	// The error name.
	Error string `json:"error"`
	// The error description.
	ErrorMessage *string `json:"error_message,omitempty"`
}

type _GetToken4XXResponse GetToken4XXResponse

// NewGetToken4XXResponse instantiates a new GetToken4XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetToken4XXResponse(error_ string) *GetToken4XXResponse {
	this := GetToken4XXResponse{}
	this.Error = error_
	return &this
}

// NewGetToken4XXResponseWithDefaults instantiates a new GetToken4XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetToken4XXResponseWithDefaults() *GetToken4XXResponse {
	this := GetToken4XXResponse{}
	return &this
}

// GetError returns the Error field value
func (o *GetToken4XXResponse) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *GetToken4XXResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *GetToken4XXResponse) SetError(v string) {
	o.Error = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GetToken4XXResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken4XXResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GetToken4XXResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GetToken4XXResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o GetToken4XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetToken4XXResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return toSerialize, nil
}

func (o *GetToken4XXResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetToken4XXResponse := _GetToken4XXResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetToken4XXResponse)

	if err != nil {
		return err
	}

	*o = GetToken4XXResponse(varGetToken4XXResponse)

	return err
}

type NullableGetToken4XXResponse struct {
	value *GetToken4XXResponse
	isSet bool
}

func (v NullableGetToken4XXResponse) Get() *GetToken4XXResponse {
	return v.value
}

func (v *NullableGetToken4XXResponse) Set(val *GetToken4XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetToken4XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetToken4XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetToken4XXResponse(val *GetToken4XXResponse) *NullableGetToken4XXResponse {
	return &NullableGetToken4XXResponse{value: val, isSet: true}
}

func (v NullableGetToken4XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetToken4XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


