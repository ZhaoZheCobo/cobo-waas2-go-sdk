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

// checks if the TransactionDepositFromAddressSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDepositFromAddressSource{}

// TransactionDepositFromAddressSource Information about the transaction source type `DepositFromAddress`. Refer to [Transaction sources and destinations](/v2/guides/transactions/sources-and-destinations) for a detailed introduction about the supported sources and destinations for each transaction type.  Switch between the tabs to display the properties for different transaction sources. 
type TransactionDepositFromAddressSource struct {
	SourceType TransactionSourceType `json:"source_type"`
	// The wallet ID.
	WalletId *string `json:"wallet_id,omitempty"`
	WalletType *WalletType `json:"wallet_type,omitempty"`
	WalletSubtype *WalletSubtype `json:"wallet_subtype,omitempty"`
	// A list of addresses.
	Addresses []string `json:"addresses"`
}

type _TransactionDepositFromAddressSource TransactionDepositFromAddressSource

// NewTransactionDepositFromAddressSource instantiates a new TransactionDepositFromAddressSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDepositFromAddressSource(sourceType TransactionSourceType, addresses []string) *TransactionDepositFromAddressSource {
	this := TransactionDepositFromAddressSource{}
	this.SourceType = sourceType
	this.Addresses = addresses
	return &this
}

// NewTransactionDepositFromAddressSourceWithDefaults instantiates a new TransactionDepositFromAddressSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDepositFromAddressSourceWithDefaults() *TransactionDepositFromAddressSource {
	this := TransactionDepositFromAddressSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *TransactionDepositFromAddressSource) GetSourceType() TransactionSourceType {
	if o == nil {
		var ret TransactionSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromAddressSource) GetSourceTypeOk() (*TransactionSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *TransactionDepositFromAddressSource) SetSourceType(v TransactionSourceType) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *TransactionDepositFromAddressSource) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromAddressSource) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *TransactionDepositFromAddressSource) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *TransactionDepositFromAddressSource) SetWalletId(v string) {
	o.WalletId = &v
}

// GetWalletType returns the WalletType field value if set, zero value otherwise.
func (o *TransactionDepositFromAddressSource) GetWalletType() WalletType {
	if o == nil || IsNil(o.WalletType) {
		var ret WalletType
		return ret
	}
	return *o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromAddressSource) GetWalletTypeOk() (*WalletType, bool) {
	if o == nil || IsNil(o.WalletType) {
		return nil, false
	}
	return o.WalletType, true
}

// HasWalletType returns a boolean if a field has been set.
func (o *TransactionDepositFromAddressSource) HasWalletType() bool {
	if o != nil && !IsNil(o.WalletType) {
		return true
	}

	return false
}

// SetWalletType gets a reference to the given WalletType and assigns it to the WalletType field.
func (o *TransactionDepositFromAddressSource) SetWalletType(v WalletType) {
	o.WalletType = &v
}

// GetWalletSubtype returns the WalletSubtype field value if set, zero value otherwise.
func (o *TransactionDepositFromAddressSource) GetWalletSubtype() WalletSubtype {
	if o == nil || IsNil(o.WalletSubtype) {
		var ret WalletSubtype
		return ret
	}
	return *o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromAddressSource) GetWalletSubtypeOk() (*WalletSubtype, bool) {
	if o == nil || IsNil(o.WalletSubtype) {
		return nil, false
	}
	return o.WalletSubtype, true
}

// HasWalletSubtype returns a boolean if a field has been set.
func (o *TransactionDepositFromAddressSource) HasWalletSubtype() bool {
	if o != nil && !IsNil(o.WalletSubtype) {
		return true
	}

	return false
}

// SetWalletSubtype gets a reference to the given WalletSubtype and assigns it to the WalletSubtype field.
func (o *TransactionDepositFromAddressSource) SetWalletSubtype(v WalletSubtype) {
	o.WalletSubtype = &v
}

// GetAddresses returns the Addresses field value
func (o *TransactionDepositFromAddressSource) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromAddressSource) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *TransactionDepositFromAddressSource) SetAddresses(v []string) {
	o.Addresses = v
}

func (o TransactionDepositFromAddressSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDepositFromAddressSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	if !IsNil(o.WalletType) {
		toSerialize["wallet_type"] = o.WalletType
	}
	if !IsNil(o.WalletSubtype) {
		toSerialize["wallet_subtype"] = o.WalletSubtype
	}
	toSerialize["addresses"] = o.Addresses
	return toSerialize, nil
}

func (o *TransactionDepositFromAddressSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"addresses",
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

	varTransactionDepositFromAddressSource := _TransactionDepositFromAddressSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionDepositFromAddressSource)

	if err != nil {
		return err
	}

	*o = TransactionDepositFromAddressSource(varTransactionDepositFromAddressSource)

	return err
}

type NullableTransactionDepositFromAddressSource struct {
	value *TransactionDepositFromAddressSource
	isSet bool
}

func (v NullableTransactionDepositFromAddressSource) Get() *TransactionDepositFromAddressSource {
	return v.value
}

func (v *NullableTransactionDepositFromAddressSource) Set(val *TransactionDepositFromAddressSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDepositFromAddressSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDepositFromAddressSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDepositFromAddressSource(val *TransactionDepositFromAddressSource) *NullableTransactionDepositFromAddressSource {
	return &NullableTransactionDepositFromAddressSource{value: val, isSet: true}
}

func (v NullableTransactionDepositFromAddressSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDepositFromAddressSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


