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

// checks if the MpcContractCallSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MpcContractCallSource{}

// MpcContractCallSource struct for MpcContractCallSource
type MpcContractCallSource struct {
	SourceType ContractCallSourceType `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	// The wallet address.
	Address string `json:"address"`
}

type _MpcContractCallSource MpcContractCallSource

// NewMpcContractCallSource instantiates a new MpcContractCallSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMpcContractCallSource(sourceType ContractCallSourceType, walletId string, address string) *MpcContractCallSource {
	this := MpcContractCallSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	this.Address = address
	return &this
}

// NewMpcContractCallSourceWithDefaults instantiates a new MpcContractCallSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMpcContractCallSourceWithDefaults() *MpcContractCallSource {
	this := MpcContractCallSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *MpcContractCallSource) GetSourceType() ContractCallSourceType {
	if o == nil {
		var ret ContractCallSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *MpcContractCallSource) GetSourceTypeOk() (*ContractCallSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *MpcContractCallSource) SetSourceType(v ContractCallSourceType) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *MpcContractCallSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *MpcContractCallSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *MpcContractCallSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetAddress returns the Address field value
func (o *MpcContractCallSource) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *MpcContractCallSource) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *MpcContractCallSource) SetAddress(v string) {
	o.Address = v
}

func (o MpcContractCallSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MpcContractCallSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

func (o *MpcContractCallSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
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

	varMpcContractCallSource := _MpcContractCallSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMpcContractCallSource)

	if err != nil {
		return err
	}

	*o = MpcContractCallSource(varMpcContractCallSource)

	return err
}

type NullableMpcContractCallSource struct {
	value *MpcContractCallSource
	isSet bool
}

func (v NullableMpcContractCallSource) Get() *MpcContractCallSource {
	return v.value
}

func (v *NullableMpcContractCallSource) Set(val *MpcContractCallSource) {
	v.value = val
	v.isSet = true
}

func (v NullableMpcContractCallSource) IsSet() bool {
	return v.isSet
}

func (v *NullableMpcContractCallSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMpcContractCallSource(val *MpcContractCallSource) *NullableMpcContractCallSource {
	return &NullableMpcContractCallSource{value: val, isSet: true}
}

func (v NullableMpcContractCallSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMpcContractCallSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


