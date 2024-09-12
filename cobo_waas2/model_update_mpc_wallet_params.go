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

// checks if the UpdateMpcWalletParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMpcWalletParams{}

// UpdateMpcWalletParams The information of MPC Wallets.
type UpdateMpcWalletParams struct {
	WalletType WalletType `json:"wallet_type"`
	// The wallet name.
	Name *string `json:"name,omitempty"`
}

type _UpdateMpcWalletParams UpdateMpcWalletParams

// NewUpdateMpcWalletParams instantiates a new UpdateMpcWalletParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMpcWalletParams(walletType WalletType) *UpdateMpcWalletParams {
	this := UpdateMpcWalletParams{}
	this.WalletType = walletType
	return &this
}

// NewUpdateMpcWalletParamsWithDefaults instantiates a new UpdateMpcWalletParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMpcWalletParamsWithDefaults() *UpdateMpcWalletParams {
	this := UpdateMpcWalletParams{}
	return &this
}

// GetWalletType returns the WalletType field value
func (o *UpdateMpcWalletParams) GetWalletType() WalletType {
	if o == nil {
		var ret WalletType
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *UpdateMpcWalletParams) GetWalletTypeOk() (*WalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *UpdateMpcWalletParams) SetWalletType(v WalletType) {
	o.WalletType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateMpcWalletParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMpcWalletParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateMpcWalletParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateMpcWalletParams) SetName(v string) {
	o.Name = &v
}

func (o UpdateMpcWalletParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMpcWalletParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_type"] = o.WalletType
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

func (o *UpdateMpcWalletParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet_type",
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

	varUpdateMpcWalletParams := _UpdateMpcWalletParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMpcWalletParams)

	if err != nil {
		return err
	}

	*o = UpdateMpcWalletParams(varUpdateMpcWalletParams)

	return err
}

type NullableUpdateMpcWalletParams struct {
	value *UpdateMpcWalletParams
	isSet bool
}

func (v NullableUpdateMpcWalletParams) Get() *UpdateMpcWalletParams {
	return v.value
}

func (v *NullableUpdateMpcWalletParams) Set(val *UpdateMpcWalletParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMpcWalletParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMpcWalletParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMpcWalletParams(val *UpdateMpcWalletParams) *NullableUpdateMpcWalletParams {
	return &NullableUpdateMpcWalletParams{value: val, isSet: true}
}

func (v NullableUpdateMpcWalletParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMpcWalletParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


