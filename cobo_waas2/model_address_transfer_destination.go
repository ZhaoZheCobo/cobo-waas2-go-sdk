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

// checks if the AddressTransferDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressTransferDestination{}

// AddressTransferDestination The information about the transaction destination type `Address`.   Specify either the `account_output` property or the `utxo_outputs` property. You can transfer tokens to multiple addresses only if you use MPC Wallets as the transaction source. You should use the `utxo_outputs` property to specify the destination addresses. 
type AddressTransferDestination struct {
	DestinationType TransferDestinationType `json:"destination_type"`
	AccountOutput *AddressTransferDestinationAccountOutput `json:"account_output,omitempty"`
	UtxoOutputs []AddressTransferDestinationUtxoOutputsInner `json:"utxo_outputs,omitempty"`
	// The address used to receive the remaining funds or change from the transaction.
	ChangeAddress *string `json:"change_address,omitempty"`
	// Whether the transaction request must be executed as a Loop transfer. For more information about Loop, see [Loop's website](https://loop.top/).   - `true`: The transaction request must be executed as a Loop transfer.   - `false`: The transaction request may not be executed as a Loop transfer. <Note>Please do not set both `force_internal` and `force_internal` as `true`.</Note> 
	ForceInternal *bool `json:"force_internal,omitempty"`
	// Whether the transaction request must not be executed as a Loop transfer. For more information about Loop, see [Loop's website](https://loop.top/).   - `true`: The transaction request must not be executed as a Loop transfer.   - `false`: The transaction request can be executed as a Loop transfer. <Note>Please do not set both `force_internal` and `force_internal` as `true`.</Note> 
	ForceExternal *bool `json:"force_external,omitempty"`
}

type _AddressTransferDestination AddressTransferDestination

// NewAddressTransferDestination instantiates a new AddressTransferDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransferDestination(destinationType TransferDestinationType) *AddressTransferDestination {
	this := AddressTransferDestination{}
	this.DestinationType = destinationType
	return &this
}

// NewAddressTransferDestinationWithDefaults instantiates a new AddressTransferDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransferDestinationWithDefaults() *AddressTransferDestination {
	this := AddressTransferDestination{}
	return &this
}

// GetDestinationType returns the DestinationType field value
func (o *AddressTransferDestination) GetDestinationType() TransferDestinationType {
	if o == nil {
		var ret TransferDestinationType
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *AddressTransferDestination) GetDestinationTypeOk() (*TransferDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *AddressTransferDestination) SetDestinationType(v TransferDestinationType) {
	o.DestinationType = v
}

// GetAccountOutput returns the AccountOutput field value if set, zero value otherwise.
func (o *AddressTransferDestination) GetAccountOutput() AddressTransferDestinationAccountOutput {
	if o == nil || IsNil(o.AccountOutput) {
		var ret AddressTransferDestinationAccountOutput
		return ret
	}
	return *o.AccountOutput
}

// GetAccountOutputOk returns a tuple with the AccountOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestination) GetAccountOutputOk() (*AddressTransferDestinationAccountOutput, bool) {
	if o == nil || IsNil(o.AccountOutput) {
		return nil, false
	}
	return o.AccountOutput, true
}

// HasAccountOutput returns a boolean if a field has been set.
func (o *AddressTransferDestination) HasAccountOutput() bool {
	if o != nil && !IsNil(o.AccountOutput) {
		return true
	}

	return false
}

// SetAccountOutput gets a reference to the given AddressTransferDestinationAccountOutput and assigns it to the AccountOutput field.
func (o *AddressTransferDestination) SetAccountOutput(v AddressTransferDestinationAccountOutput) {
	o.AccountOutput = &v
}

// GetUtxoOutputs returns the UtxoOutputs field value if set, zero value otherwise.
func (o *AddressTransferDestination) GetUtxoOutputs() []AddressTransferDestinationUtxoOutputsInner {
	if o == nil || IsNil(o.UtxoOutputs) {
		var ret []AddressTransferDestinationUtxoOutputsInner
		return ret
	}
	return o.UtxoOutputs
}

// GetUtxoOutputsOk returns a tuple with the UtxoOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestination) GetUtxoOutputsOk() ([]AddressTransferDestinationUtxoOutputsInner, bool) {
	if o == nil || IsNil(o.UtxoOutputs) {
		return nil, false
	}
	return o.UtxoOutputs, true
}

// HasUtxoOutputs returns a boolean if a field has been set.
func (o *AddressTransferDestination) HasUtxoOutputs() bool {
	if o != nil && !IsNil(o.UtxoOutputs) {
		return true
	}

	return false
}

// SetUtxoOutputs gets a reference to the given []AddressTransferDestinationUtxoOutputsInner and assigns it to the UtxoOutputs field.
func (o *AddressTransferDestination) SetUtxoOutputs(v []AddressTransferDestinationUtxoOutputsInner) {
	o.UtxoOutputs = v
}

// GetChangeAddress returns the ChangeAddress field value if set, zero value otherwise.
func (o *AddressTransferDestination) GetChangeAddress() string {
	if o == nil || IsNil(o.ChangeAddress) {
		var ret string
		return ret
	}
	return *o.ChangeAddress
}

// GetChangeAddressOk returns a tuple with the ChangeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestination) GetChangeAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeAddress) {
		return nil, false
	}
	return o.ChangeAddress, true
}

// HasChangeAddress returns a boolean if a field has been set.
func (o *AddressTransferDestination) HasChangeAddress() bool {
	if o != nil && !IsNil(o.ChangeAddress) {
		return true
	}

	return false
}

// SetChangeAddress gets a reference to the given string and assigns it to the ChangeAddress field.
func (o *AddressTransferDestination) SetChangeAddress(v string) {
	o.ChangeAddress = &v
}

// GetForceInternal returns the ForceInternal field value if set, zero value otherwise.
func (o *AddressTransferDestination) GetForceInternal() bool {
	if o == nil || IsNil(o.ForceInternal) {
		var ret bool
		return ret
	}
	return *o.ForceInternal
}

// GetForceInternalOk returns a tuple with the ForceInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestination) GetForceInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceInternal) {
		return nil, false
	}
	return o.ForceInternal, true
}

// HasForceInternal returns a boolean if a field has been set.
func (o *AddressTransferDestination) HasForceInternal() bool {
	if o != nil && !IsNil(o.ForceInternal) {
		return true
	}

	return false
}

// SetForceInternal gets a reference to the given bool and assigns it to the ForceInternal field.
func (o *AddressTransferDestination) SetForceInternal(v bool) {
	o.ForceInternal = &v
}

// GetForceExternal returns the ForceExternal field value if set, zero value otherwise.
func (o *AddressTransferDestination) GetForceExternal() bool {
	if o == nil || IsNil(o.ForceExternal) {
		var ret bool
		return ret
	}
	return *o.ForceExternal
}

// GetForceExternalOk returns a tuple with the ForceExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestination) GetForceExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceExternal) {
		return nil, false
	}
	return o.ForceExternal, true
}

// HasForceExternal returns a boolean if a field has been set.
func (o *AddressTransferDestination) HasForceExternal() bool {
	if o != nil && !IsNil(o.ForceExternal) {
		return true
	}

	return false
}

// SetForceExternal gets a reference to the given bool and assigns it to the ForceExternal field.
func (o *AddressTransferDestination) SetForceExternal(v bool) {
	o.ForceExternal = &v
}

func (o AddressTransferDestination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransferDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination_type"] = o.DestinationType
	if !IsNil(o.AccountOutput) {
		toSerialize["account_output"] = o.AccountOutput
	}
	if !IsNil(o.UtxoOutputs) {
		toSerialize["utxo_outputs"] = o.UtxoOutputs
	}
	if !IsNil(o.ChangeAddress) {
		toSerialize["change_address"] = o.ChangeAddress
	}
	if !IsNil(o.ForceInternal) {
		toSerialize["force_internal"] = o.ForceInternal
	}
	if !IsNil(o.ForceExternal) {
		toSerialize["force_external"] = o.ForceExternal
	}
	return toSerialize, nil
}

func (o *AddressTransferDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_type",
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

	varAddressTransferDestination := _AddressTransferDestination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransferDestination)

	if err != nil {
		return err
	}

	*o = AddressTransferDestination(varAddressTransferDestination)

	return err
}

type NullableAddressTransferDestination struct {
	value *AddressTransferDestination
	isSet bool
}

func (v NullableAddressTransferDestination) Get() *AddressTransferDestination {
	return v.value
}

func (v *NullableAddressTransferDestination) Set(val *AddressTransferDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransferDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransferDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransferDestination(val *AddressTransferDestination) *NullableAddressTransferDestination {
	return &NullableAddressTransferDestination{value: val, isSet: true}
}

func (v NullableAddressTransferDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransferDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


