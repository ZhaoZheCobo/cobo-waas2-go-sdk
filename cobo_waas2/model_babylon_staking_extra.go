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

// checks if the BabylonStakingExtra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BabylonStakingExtra{}

// BabylonStakingExtra struct for BabylonStakingExtra
type BabylonStakingExtra struct {
	PoolType StakingPoolType `json:"pool_type"`
	// The Proof-of-Stake (PoS) chain.
	PosChain string `json:"pos_chain"`
	// The estimated time when the bitcoins will be unlocked, in Unix timestamp format, measured in milliseconds.
	UnlockTimestamp *int64 `json:"unlock_timestamp,omitempty"`
	// The block height at which the bitcoins will be unlocked.
	UnlockBlockHeight *int64 `json:"unlock_block_height,omitempty"`
	// The address receiving the staked bitcoins.
	StakeAddress *string `json:"stake_address,omitempty"`
	// The address receiving the unlocked bitcoins.
	UnbondAddress *string `json:"unbond_address,omitempty"`
}

type _BabylonStakingExtra BabylonStakingExtra

// NewBabylonStakingExtra instantiates a new BabylonStakingExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStakingExtra(poolType StakingPoolType, posChain string) *BabylonStakingExtra {
	this := BabylonStakingExtra{}
	this.PoolType = poolType
	this.PosChain = posChain
	return &this
}

// NewBabylonStakingExtraWithDefaults instantiates a new BabylonStakingExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStakingExtraWithDefaults() *BabylonStakingExtra {
	this := BabylonStakingExtra{}
	return &this
}

// GetPoolType returns the PoolType field value
func (o *BabylonStakingExtra) GetPoolType() StakingPoolType {
	if o == nil {
		var ret StakingPoolType
		return ret
	}

	return o.PoolType
}

// GetPoolTypeOk returns a tuple with the PoolType field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingExtra) GetPoolTypeOk() (*StakingPoolType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolType, true
}

// SetPoolType sets field value
func (o *BabylonStakingExtra) SetPoolType(v StakingPoolType) {
	o.PoolType = v
}

// GetPosChain returns the PosChain field value
func (o *BabylonStakingExtra) GetPosChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PosChain
}

// GetPosChainOk returns a tuple with the PosChain field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingExtra) GetPosChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PosChain, true
}

// SetPosChain sets field value
func (o *BabylonStakingExtra) SetPosChain(v string) {
	o.PosChain = v
}

// GetUnlockTimestamp returns the UnlockTimestamp field value if set, zero value otherwise.
func (o *BabylonStakingExtra) GetUnlockTimestamp() int64 {
	if o == nil || IsNil(o.UnlockTimestamp) {
		var ret int64
		return ret
	}
	return *o.UnlockTimestamp
}

// GetUnlockTimestampOk returns a tuple with the UnlockTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingExtra) GetUnlockTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.UnlockTimestamp) {
		return nil, false
	}
	return o.UnlockTimestamp, true
}

// HasUnlockTimestamp returns a boolean if a field has been set.
func (o *BabylonStakingExtra) HasUnlockTimestamp() bool {
	if o != nil && !IsNil(o.UnlockTimestamp) {
		return true
	}

	return false
}

// SetUnlockTimestamp gets a reference to the given int64 and assigns it to the UnlockTimestamp field.
func (o *BabylonStakingExtra) SetUnlockTimestamp(v int64) {
	o.UnlockTimestamp = &v
}

// GetUnlockBlockHeight returns the UnlockBlockHeight field value if set, zero value otherwise.
func (o *BabylonStakingExtra) GetUnlockBlockHeight() int64 {
	if o == nil || IsNil(o.UnlockBlockHeight) {
		var ret int64
		return ret
	}
	return *o.UnlockBlockHeight
}

// GetUnlockBlockHeightOk returns a tuple with the UnlockBlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingExtra) GetUnlockBlockHeightOk() (*int64, bool) {
	if o == nil || IsNil(o.UnlockBlockHeight) {
		return nil, false
	}
	return o.UnlockBlockHeight, true
}

// HasUnlockBlockHeight returns a boolean if a field has been set.
func (o *BabylonStakingExtra) HasUnlockBlockHeight() bool {
	if o != nil && !IsNil(o.UnlockBlockHeight) {
		return true
	}

	return false
}

// SetUnlockBlockHeight gets a reference to the given int64 and assigns it to the UnlockBlockHeight field.
func (o *BabylonStakingExtra) SetUnlockBlockHeight(v int64) {
	o.UnlockBlockHeight = &v
}

// GetStakeAddress returns the StakeAddress field value if set, zero value otherwise.
func (o *BabylonStakingExtra) GetStakeAddress() string {
	if o == nil || IsNil(o.StakeAddress) {
		var ret string
		return ret
	}
	return *o.StakeAddress
}

// GetStakeAddressOk returns a tuple with the StakeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingExtra) GetStakeAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StakeAddress) {
		return nil, false
	}
	return o.StakeAddress, true
}

// HasStakeAddress returns a boolean if a field has been set.
func (o *BabylonStakingExtra) HasStakeAddress() bool {
	if o != nil && !IsNil(o.StakeAddress) {
		return true
	}

	return false
}

// SetStakeAddress gets a reference to the given string and assigns it to the StakeAddress field.
func (o *BabylonStakingExtra) SetStakeAddress(v string) {
	o.StakeAddress = &v
}

// GetUnbondAddress returns the UnbondAddress field value if set, zero value otherwise.
func (o *BabylonStakingExtra) GetUnbondAddress() string {
	if o == nil || IsNil(o.UnbondAddress) {
		var ret string
		return ret
	}
	return *o.UnbondAddress
}

// GetUnbondAddressOk returns a tuple with the UnbondAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingExtra) GetUnbondAddressOk() (*string, bool) {
	if o == nil || IsNil(o.UnbondAddress) {
		return nil, false
	}
	return o.UnbondAddress, true
}

// HasUnbondAddress returns a boolean if a field has been set.
func (o *BabylonStakingExtra) HasUnbondAddress() bool {
	if o != nil && !IsNil(o.UnbondAddress) {
		return true
	}

	return false
}

// SetUnbondAddress gets a reference to the given string and assigns it to the UnbondAddress field.
func (o *BabylonStakingExtra) SetUnbondAddress(v string) {
	o.UnbondAddress = &v
}

func (o BabylonStakingExtra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BabylonStakingExtra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pool_type"] = o.PoolType
	toSerialize["pos_chain"] = o.PosChain
	if !IsNil(o.UnlockTimestamp) {
		toSerialize["unlock_timestamp"] = o.UnlockTimestamp
	}
	if !IsNil(o.UnlockBlockHeight) {
		toSerialize["unlock_block_height"] = o.UnlockBlockHeight
	}
	if !IsNil(o.StakeAddress) {
		toSerialize["stake_address"] = o.StakeAddress
	}
	if !IsNil(o.UnbondAddress) {
		toSerialize["unbond_address"] = o.UnbondAddress
	}
	return toSerialize, nil
}

func (o *BabylonStakingExtra) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pool_type",
		"pos_chain",
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

	varBabylonStakingExtra := _BabylonStakingExtra{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBabylonStakingExtra)

	if err != nil {
		return err
	}

	*o = BabylonStakingExtra(varBabylonStakingExtra)

	return err
}

type NullableBabylonStakingExtra struct {
	value *BabylonStakingExtra
	isSet bool
}

func (v NullableBabylonStakingExtra) Get() *BabylonStakingExtra {
	return v.value
}

func (v *NullableBabylonStakingExtra) Set(val *BabylonStakingExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStakingExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStakingExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStakingExtra(val *BabylonStakingExtra) *NullableBabylonStakingExtra {
	return &NullableBabylonStakingExtra{value: val, isSet: true}
}

func (v NullableBabylonStakingExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStakingExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


