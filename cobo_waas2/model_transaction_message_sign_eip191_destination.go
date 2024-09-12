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

// checks if the TransactionMessageSignEIP191Destination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionMessageSignEIP191Destination{}

// TransactionMessageSignEIP191Destination Information about the transaction destination type `EVM_EIP_191_Signature`. Refer to [Transaction sources and destinations](/v2/guides/transactions/sources-and-destinations) for a detailed introduction about the supported sources and destinations for each transaction type.  Switch between the tabs to display the properties for different transaction destinations. 
type TransactionMessageSignEIP191Destination struct {
	DestinationType TransactionDestinationType `json:"destination_type"`
	// The raw data of the message to be signed, encoded in Base64 format.
	Message string `json:"message"`
}

type _TransactionMessageSignEIP191Destination TransactionMessageSignEIP191Destination

// NewTransactionMessageSignEIP191Destination instantiates a new TransactionMessageSignEIP191Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionMessageSignEIP191Destination(destinationType TransactionDestinationType, message string) *TransactionMessageSignEIP191Destination {
	this := TransactionMessageSignEIP191Destination{}
	this.DestinationType = destinationType
	this.Message = message
	return &this
}

// NewTransactionMessageSignEIP191DestinationWithDefaults instantiates a new TransactionMessageSignEIP191Destination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionMessageSignEIP191DestinationWithDefaults() *TransactionMessageSignEIP191Destination {
	this := TransactionMessageSignEIP191Destination{}
	return &this
}

// GetDestinationType returns the DestinationType field value
func (o *TransactionMessageSignEIP191Destination) GetDestinationType() TransactionDestinationType {
	if o == nil {
		var ret TransactionDestinationType
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *TransactionMessageSignEIP191Destination) GetDestinationTypeOk() (*TransactionDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *TransactionMessageSignEIP191Destination) SetDestinationType(v TransactionDestinationType) {
	o.DestinationType = v
}

// GetMessage returns the Message field value
func (o *TransactionMessageSignEIP191Destination) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TransactionMessageSignEIP191Destination) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TransactionMessageSignEIP191Destination) SetMessage(v string) {
	o.Message = v
}

func (o TransactionMessageSignEIP191Destination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionMessageSignEIP191Destination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *TransactionMessageSignEIP191Destination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_type",
		"message",
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

	varTransactionMessageSignEIP191Destination := _TransactionMessageSignEIP191Destination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionMessageSignEIP191Destination)

	if err != nil {
		return err
	}

	*o = TransactionMessageSignEIP191Destination(varTransactionMessageSignEIP191Destination)

	return err
}

type NullableTransactionMessageSignEIP191Destination struct {
	value *TransactionMessageSignEIP191Destination
	isSet bool
}

func (v NullableTransactionMessageSignEIP191Destination) Get() *TransactionMessageSignEIP191Destination {
	return v.value
}

func (v *NullableTransactionMessageSignEIP191Destination) Set(val *TransactionMessageSignEIP191Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionMessageSignEIP191Destination) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionMessageSignEIP191Destination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionMessageSignEIP191Destination(val *TransactionMessageSignEIP191Destination) *NullableTransactionMessageSignEIP191Destination {
	return &NullableTransactionMessageSignEIP191Destination{value: val, isSet: true}
}

func (v NullableTransactionMessageSignEIP191Destination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionMessageSignEIP191Destination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


