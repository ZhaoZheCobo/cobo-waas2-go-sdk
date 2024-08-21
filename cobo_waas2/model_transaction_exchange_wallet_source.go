/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionExchangeWalletSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionExchangeWalletSource{}

// TransactionExchangeWalletSource Information about the transaction source types `Main` and `Sub`. 
type TransactionExchangeWalletSource struct {
	SourceType TransactionSourceType `json:"source_type"`
	ExchangeId ExchangeId `json:"exchange_id"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	// The exchange trading account or a sub-wallet ID.
	TradingAccountType *string `json:"trading_account_type,omitempty"`
}

type _TransactionExchangeWalletSource TransactionExchangeWalletSource

// NewTransactionExchangeWalletSource instantiates a new TransactionExchangeWalletSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionExchangeWalletSource(sourceType TransactionSourceType, exchangeId ExchangeId, walletId string) *TransactionExchangeWalletSource {
	this := TransactionExchangeWalletSource{}
	this.SourceType = sourceType
	this.ExchangeId = exchangeId
	this.WalletId = walletId
	return &this
}

// NewTransactionExchangeWalletSourceWithDefaults instantiates a new TransactionExchangeWalletSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionExchangeWalletSourceWithDefaults() *TransactionExchangeWalletSource {
	this := TransactionExchangeWalletSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *TransactionExchangeWalletSource) GetSourceType() TransactionSourceType {
	if o == nil {
		var ret TransactionSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *TransactionExchangeWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *TransactionExchangeWalletSource) SetSourceType(v TransactionSourceType) {
	o.SourceType = v
}

// GetExchangeId returns the ExchangeId field value
func (o *TransactionExchangeWalletSource) GetExchangeId() ExchangeId {
	if o == nil {
		var ret ExchangeId
		return ret
	}

	return o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value
// and a boolean to check if the value has been set.
func (o *TransactionExchangeWalletSource) GetExchangeIdOk() (*ExchangeId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeId, true
}

// SetExchangeId sets field value
func (o *TransactionExchangeWalletSource) SetExchangeId(v ExchangeId) {
	o.ExchangeId = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionExchangeWalletSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionExchangeWalletSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionExchangeWalletSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetTradingAccountType returns the TradingAccountType field value if set, zero value otherwise.
func (o *TransactionExchangeWalletSource) GetTradingAccountType() string {
	if o == nil || IsNil(o.TradingAccountType) {
		var ret string
		return ret
	}
	return *o.TradingAccountType
}

// GetTradingAccountTypeOk returns a tuple with the TradingAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionExchangeWalletSource) GetTradingAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TradingAccountType) {
		return nil, false
	}
	return o.TradingAccountType, true
}

// HasTradingAccountType returns a boolean if a field has been set.
func (o *TransactionExchangeWalletSource) HasTradingAccountType() bool {
	if o != nil && !IsNil(o.TradingAccountType) {
		return true
	}

	return false
}

// SetTradingAccountType gets a reference to the given string and assigns it to the TradingAccountType field.
func (o *TransactionExchangeWalletSource) SetTradingAccountType(v string) {
	o.TradingAccountType = &v
}

func (o TransactionExchangeWalletSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionExchangeWalletSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["exchange_id"] = o.ExchangeId
	toSerialize["wallet_id"] = o.WalletId
	if !IsNil(o.TradingAccountType) {
		toSerialize["trading_account_type"] = o.TradingAccountType
	}
	return toSerialize, nil
}

func (o *TransactionExchangeWalletSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"exchange_id",
		"wallet_id",
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

	varTransactionExchangeWalletSource := _TransactionExchangeWalletSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionExchangeWalletSource)

	if err != nil {
		return err
	}

	*o = TransactionExchangeWalletSource(varTransactionExchangeWalletSource)

	return err
}

type NullableTransactionExchangeWalletSource struct {
	value *TransactionExchangeWalletSource
	isSet bool
}

func (v NullableTransactionExchangeWalletSource) Get() *TransactionExchangeWalletSource {
	return v.value
}

func (v *NullableTransactionExchangeWalletSource) Set(val *TransactionExchangeWalletSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionExchangeWalletSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionExchangeWalletSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionExchangeWalletSource(val *TransactionExchangeWalletSource) *NullableTransactionExchangeWalletSource {
	return &NullableTransactionExchangeWalletSource{value: val, isSet: true}
}

func (v NullableTransactionExchangeWalletSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionExchangeWalletSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


