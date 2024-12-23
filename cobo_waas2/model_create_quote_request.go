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

// checks if the CreateQuoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateQuoteRequest{}

// CreateQuoteRequest struct for CreateQuoteRequest
type CreateQuoteRequest struct {
	// Unique id of the token to pay.
	PayTokenId string `json:"pay_token_id"`
	// Unique id of the token to receive.
	ReceiveTokenId string `json:"receive_token_id"`
	// The amount of token to swap.
	PayAmount *string `json:"pay_amount,omitempty"`
	// The amount of token to receive.
	ReceiveAmount *string `json:"receive_amount,omitempty"`
}

type _CreateQuoteRequest CreateQuoteRequest

// NewCreateQuoteRequest instantiates a new CreateQuoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateQuoteRequest(payTokenId string, receiveTokenId string) *CreateQuoteRequest {
	this := CreateQuoteRequest{}
	this.PayTokenId = payTokenId
	this.ReceiveTokenId = receiveTokenId
	return &this
}

// NewCreateQuoteRequestWithDefaults instantiates a new CreateQuoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateQuoteRequestWithDefaults() *CreateQuoteRequest {
	this := CreateQuoteRequest{}
	return &this
}

// GetPayTokenId returns the PayTokenId field value
func (o *CreateQuoteRequest) GetPayTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayTokenId
}

// GetPayTokenIdOk returns a tuple with the PayTokenId field value
// and a boolean to check if the value has been set.
func (o *CreateQuoteRequest) GetPayTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayTokenId, true
}

// SetPayTokenId sets field value
func (o *CreateQuoteRequest) SetPayTokenId(v string) {
	o.PayTokenId = v
}

// GetReceiveTokenId returns the ReceiveTokenId field value
func (o *CreateQuoteRequest) GetReceiveTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiveTokenId
}

// GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field value
// and a boolean to check if the value has been set.
func (o *CreateQuoteRequest) GetReceiveTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiveTokenId, true
}

// SetReceiveTokenId sets field value
func (o *CreateQuoteRequest) SetReceiveTokenId(v string) {
	o.ReceiveTokenId = v
}

// GetPayAmount returns the PayAmount field value if set, zero value otherwise.
func (o *CreateQuoteRequest) GetPayAmount() string {
	if o == nil || IsNil(o.PayAmount) {
		var ret string
		return ret
	}
	return *o.PayAmount
}

// GetPayAmountOk returns a tuple with the PayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateQuoteRequest) GetPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PayAmount) {
		return nil, false
	}
	return o.PayAmount, true
}

// HasPayAmount returns a boolean if a field has been set.
func (o *CreateQuoteRequest) HasPayAmount() bool {
	if o != nil && !IsNil(o.PayAmount) {
		return true
	}

	return false
}

// SetPayAmount gets a reference to the given string and assigns it to the PayAmount field.
func (o *CreateQuoteRequest) SetPayAmount(v string) {
	o.PayAmount = &v
}

// GetReceiveAmount returns the ReceiveAmount field value if set, zero value otherwise.
func (o *CreateQuoteRequest) GetReceiveAmount() string {
	if o == nil || IsNil(o.ReceiveAmount) {
		var ret string
		return ret
	}
	return *o.ReceiveAmount
}

// GetReceiveAmountOk returns a tuple with the ReceiveAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateQuoteRequest) GetReceiveAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveAmount) {
		return nil, false
	}
	return o.ReceiveAmount, true
}

// HasReceiveAmount returns a boolean if a field has been set.
func (o *CreateQuoteRequest) HasReceiveAmount() bool {
	if o != nil && !IsNil(o.ReceiveAmount) {
		return true
	}

	return false
}

// SetReceiveAmount gets a reference to the given string and assigns it to the ReceiveAmount field.
func (o *CreateQuoteRequest) SetReceiveAmount(v string) {
	o.ReceiveAmount = &v
}

func (o CreateQuoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateQuoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pay_token_id"] = o.PayTokenId
	toSerialize["receive_token_id"] = o.ReceiveTokenId
	if !IsNil(o.PayAmount) {
		toSerialize["pay_amount"] = o.PayAmount
	}
	if !IsNil(o.ReceiveAmount) {
		toSerialize["receive_amount"] = o.ReceiveAmount
	}
	return toSerialize, nil
}

func (o *CreateQuoteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pay_token_id",
		"receive_token_id",
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

	varCreateQuoteRequest := _CreateQuoteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateQuoteRequest)

	if err != nil {
		return err
	}

	*o = CreateQuoteRequest(varCreateQuoteRequest)

	return err
}

type NullableCreateQuoteRequest struct {
	value *CreateQuoteRequest
	isSet bool
}

func (v NullableCreateQuoteRequest) Get() *CreateQuoteRequest {
	return v.value
}

func (v *NullableCreateQuoteRequest) Set(val *CreateQuoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateQuoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateQuoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateQuoteRequest(val *CreateQuoteRequest) *NullableCreateQuoteRequest {
	return &NullableCreateQuoteRequest{value: val, isSet: true}
}

func (v NullableCreateQuoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateQuoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

