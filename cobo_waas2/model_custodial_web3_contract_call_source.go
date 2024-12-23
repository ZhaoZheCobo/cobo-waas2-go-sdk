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

// checks if the CustodialWeb3ContractCallSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustodialWeb3ContractCallSource{}

// CustodialWeb3ContractCallSource struct for CustodialWeb3ContractCallSource
type CustodialWeb3ContractCallSource struct {
	SourceType ContractCallSourceType `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	// The wallet address.
	Address string `json:"address"`
}

type _CustodialWeb3ContractCallSource CustodialWeb3ContractCallSource

// NewCustodialWeb3ContractCallSource instantiates a new CustodialWeb3ContractCallSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustodialWeb3ContractCallSource(sourceType ContractCallSourceType, walletId string, address string) *CustodialWeb3ContractCallSource {
	this := CustodialWeb3ContractCallSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	this.Address = address
	return &this
}

// NewCustodialWeb3ContractCallSourceWithDefaults instantiates a new CustodialWeb3ContractCallSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustodialWeb3ContractCallSourceWithDefaults() *CustodialWeb3ContractCallSource {
	this := CustodialWeb3ContractCallSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *CustodialWeb3ContractCallSource) GetSourceType() ContractCallSourceType {
	if o == nil {
		var ret ContractCallSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *CustodialWeb3ContractCallSource) GetSourceTypeOk() (*ContractCallSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *CustodialWeb3ContractCallSource) SetSourceType(v ContractCallSourceType) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *CustodialWeb3ContractCallSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *CustodialWeb3ContractCallSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *CustodialWeb3ContractCallSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetAddress returns the Address field value
func (o *CustodialWeb3ContractCallSource) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CustodialWeb3ContractCallSource) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CustodialWeb3ContractCallSource) SetAddress(v string) {
	o.Address = v
}

func (o CustodialWeb3ContractCallSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustodialWeb3ContractCallSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

func (o *CustodialWeb3ContractCallSource) UnmarshalJSON(data []byte) (err error) {
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

	varCustodialWeb3ContractCallSource := _CustodialWeb3ContractCallSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustodialWeb3ContractCallSource)

	if err != nil {
		return err
	}

	*o = CustodialWeb3ContractCallSource(varCustodialWeb3ContractCallSource)

	return err
}

type NullableCustodialWeb3ContractCallSource struct {
	value *CustodialWeb3ContractCallSource
	isSet bool
}

func (v NullableCustodialWeb3ContractCallSource) Get() *CustodialWeb3ContractCallSource {
	return v.value
}

func (v *NullableCustodialWeb3ContractCallSource) Set(val *CustodialWeb3ContractCallSource) {
	v.value = val
	v.isSet = true
}

func (v NullableCustodialWeb3ContractCallSource) IsSet() bool {
	return v.isSet
}

func (v *NullableCustodialWeb3ContractCallSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustodialWeb3ContractCallSource(val *CustodialWeb3ContractCallSource) *NullableCustodialWeb3ContractCallSource {
	return &NullableCustodialWeb3ContractCallSource{value: val, isSet: true}
}

func (v NullableCustodialWeb3ContractCallSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustodialWeb3ContractCallSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

