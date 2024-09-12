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

// checks if the MPCDelegate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MPCDelegate{}

// MPCDelegate The information about the MPC Wallet as the Delegate.
type MPCDelegate struct {
	DelegateType CoboSafeDelegateType `json:"delegate_type"`
	// The wallet ID of the Delegate. This is required when initiating a transfer from Smart Contract Wallets (Safe{Wallet}).
	WalletId string `json:"wallet_id"`
	// The wallet address of the Delegate. This is required when initiating a transfer from Smart Contract Wallets (Safe{Wallet}).
	Address string `json:"address"`
}

type _MPCDelegate MPCDelegate

// NewMPCDelegate instantiates a new MPCDelegate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMPCDelegate(delegateType CoboSafeDelegateType, walletId string, address string) *MPCDelegate {
	this := MPCDelegate{}
	this.DelegateType = delegateType
	this.WalletId = walletId
	this.Address = address
	return &this
}

// NewMPCDelegateWithDefaults instantiates a new MPCDelegate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMPCDelegateWithDefaults() *MPCDelegate {
	this := MPCDelegate{}
	var delegateType CoboSafeDelegateType = COBOSAFEDELEGATETYPE_ORG_CONTROLLED
	this.DelegateType = delegateType
	return &this
}

// GetDelegateType returns the DelegateType field value
func (o *MPCDelegate) GetDelegateType() CoboSafeDelegateType {
	if o == nil {
		var ret CoboSafeDelegateType
		return ret
	}

	return o.DelegateType
}

// GetDelegateTypeOk returns a tuple with the DelegateType field value
// and a boolean to check if the value has been set.
func (o *MPCDelegate) GetDelegateTypeOk() (*CoboSafeDelegateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DelegateType, true
}

// SetDelegateType sets field value
func (o *MPCDelegate) SetDelegateType(v CoboSafeDelegateType) {
	o.DelegateType = v
}

// GetWalletId returns the WalletId field value
func (o *MPCDelegate) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *MPCDelegate) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *MPCDelegate) SetWalletId(v string) {
	o.WalletId = v
}

// GetAddress returns the Address field value
func (o *MPCDelegate) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *MPCDelegate) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *MPCDelegate) SetAddress(v string) {
	o.Address = v
}

func (o MPCDelegate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MPCDelegate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["delegate_type"] = o.DelegateType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

func (o *MPCDelegate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"delegate_type",
		"wallet_id",
		"address",
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

	varMPCDelegate := _MPCDelegate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMPCDelegate)

	if err != nil {
		return err
	}

	*o = MPCDelegate(varMPCDelegate)

	return err
}

type NullableMPCDelegate struct {
	value *MPCDelegate
	isSet bool
}

func (v NullableMPCDelegate) Get() *MPCDelegate {
	return v.value
}

func (v *NullableMPCDelegate) Set(val *MPCDelegate) {
	v.value = val
	v.isSet = true
}

func (v NullableMPCDelegate) IsSet() bool {
	return v.isSet
}

func (v *NullableMPCDelegate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMPCDelegate(val *MPCDelegate) *NullableMPCDelegate {
	return &NullableMPCDelegate{value: val, isSet: true}
}

func (v NullableMPCDelegate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMPCDelegate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


