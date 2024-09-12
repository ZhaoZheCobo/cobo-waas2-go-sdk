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

// KeyShareHolderGroupType The type of key share holder group. Possible values include:  - `MainGroup`: A [Main Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups).  - `SigningGroup`: A [Signing Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups).  - `RecoveryGroup`: A [Recovery Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups).  **Note:** For `MainGroup` and `SigningGroup`, a Cobo key share holder will be added automatically. 
type KeyShareHolderGroupType string

// List of KeyShareHolderGroupType
const (
	KEYSHAREHOLDERGROUPTYPE_MAIN_GROUP KeyShareHolderGroupType = "MainGroup"
	KEYSHAREHOLDERGROUPTYPE_SIGNING_GROUP KeyShareHolderGroupType = "SigningGroup"
	KEYSHAREHOLDERGROUPTYPE_RECOVERY_GROUP KeyShareHolderGroupType = "RecoveryGroup"
)

// All allowed values of KeyShareHolderGroupType enum
var AllowedKeyShareHolderGroupTypeEnumValues = []KeyShareHolderGroupType{
	"MainGroup",
	"SigningGroup",
	"RecoveryGroup",
}

func (v *KeyShareHolderGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyShareHolderGroupType(value)
	for _, existing := range AllowedKeyShareHolderGroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = KeyShareHolderGroupType("unknown")
    return nil
}

// NewKeyShareHolderGroupTypeFromValue returns a pointer to a valid KeyShareHolderGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyShareHolderGroupTypeFromValue(v string) (*KeyShareHolderGroupType, error) {
	ev := KeyShareHolderGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyShareHolderGroupType: valid values are %v", v, AllowedKeyShareHolderGroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyShareHolderGroupType) IsValid() bool {
	for _, existing := range AllowedKeyShareHolderGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeyShareHolderGroupType value
func (v KeyShareHolderGroupType) Ptr() *KeyShareHolderGroupType {
	return &v
}

type NullableKeyShareHolderGroupType struct {
	value *KeyShareHolderGroupType
	isSet bool
}

func (v NullableKeyShareHolderGroupType) Get() *KeyShareHolderGroupType {
	return v.value
}

func (v *NullableKeyShareHolderGroupType) Set(val *KeyShareHolderGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyShareHolderGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyShareHolderGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyShareHolderGroupType(val *KeyShareHolderGroupType) *NullableKeyShareHolderGroupType {
	return &NullableKeyShareHolderGroupType{value: val, isSet: true}
}

func (v NullableKeyShareHolderGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyShareHolderGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

