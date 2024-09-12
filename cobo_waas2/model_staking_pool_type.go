/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// StakingPoolType The type of the staking pool.
type StakingPoolType string

// List of StakingPoolType
const (
	STAKINGPOOLTYPE_BABYLON StakingPoolType = "Babylon"
)

// All allowed values of StakingPoolType enum
var AllowedStakingPoolTypeEnumValues = []StakingPoolType{
	"Babylon",
}

func (v *StakingPoolType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StakingPoolType(value)
	for _, existing := range AllowedStakingPoolTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = StakingPoolType("unknown")
    return nil
}

// NewStakingPoolTypeFromValue returns a pointer to a valid StakingPoolType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStakingPoolTypeFromValue(v string) (*StakingPoolType, error) {
	ev := StakingPoolType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StakingPoolType: valid values are %v", v, AllowedStakingPoolTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StakingPoolType) IsValid() bool {
	for _, existing := range AllowedStakingPoolTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StakingPoolType value
func (v StakingPoolType) Ptr() *StakingPoolType {
	return &v
}

type NullableStakingPoolType struct {
	value *StakingPoolType
	isSet bool
}

func (v NullableStakingPoolType) Get() *StakingPoolType {
	return v.value
}

func (v *NullableStakingPoolType) Set(val *StakingPoolType) {
	v.value = val
	v.isSet = true
}

func (v NullableStakingPoolType) IsSet() bool {
	return v.isSet
}

func (v *NullableStakingPoolType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakingPoolType(val *StakingPoolType) *NullableStakingPoolType {
	return &NullableStakingPoolType{value: val, isSet: true}
}

func (v NullableStakingPoolType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakingPoolType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

