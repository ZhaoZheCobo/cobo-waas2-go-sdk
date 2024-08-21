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

// checks if the TransferParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferParams{}

// TransferParams The information about a token transfer.
type TransferParams struct {
	// The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization.
	RequestId string `json:"request_id"`
	Source TransferSource `json:"source"`
	// The token ID of the transferred token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens). For transfers from Exchange Wallets, this property value represents the asset ID.
	TokenId string `json:"token_id"`
	Destination TransferDestination `json:"destination"`
	// The custom category for you to identify your transactions.
	CategoryNames []string `json:"category_names,omitempty"`
	// The description of the transfer.
	Description *string `json:"description,omitempty"`
	Fee *TransactionRequestFee `json:"fee,omitempty"`
}

type _TransferParams TransferParams

// NewTransferParams instantiates a new TransferParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferParams(requestId string, source TransferSource, tokenId string, destination TransferDestination) *TransferParams {
	this := TransferParams{}
	this.RequestId = requestId
	this.Source = source
	this.TokenId = tokenId
	this.Destination = destination
	return &this
}

// NewTransferParamsWithDefaults instantiates a new TransferParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferParamsWithDefaults() *TransferParams {
	this := TransferParams{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransferParams) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferParams) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferParams) SetRequestId(v string) {
	o.RequestId = v
}

// GetSource returns the Source field value
func (o *TransferParams) GetSource() TransferSource {
	if o == nil {
		var ret TransferSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TransferParams) GetSourceOk() (*TransferSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *TransferParams) SetSource(v TransferSource) {
	o.Source = v
}

// GetTokenId returns the TokenId field value
func (o *TransferParams) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TransferParams) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TransferParams) SetTokenId(v string) {
	o.TokenId = v
}

// GetDestination returns the Destination field value
func (o *TransferParams) GetDestination() TransferDestination {
	if o == nil {
		var ret TransferDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *TransferParams) GetDestinationOk() (*TransferDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *TransferParams) SetDestination(v TransferDestination) {
	o.Destination = v
}

// GetCategoryNames returns the CategoryNames field value if set, zero value otherwise.
func (o *TransferParams) GetCategoryNames() []string {
	if o == nil || IsNil(o.CategoryNames) {
		var ret []string
		return ret
	}
	return o.CategoryNames
}

// GetCategoryNamesOk returns a tuple with the CategoryNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferParams) GetCategoryNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CategoryNames) {
		return nil, false
	}
	return o.CategoryNames, true
}

// HasCategoryNames returns a boolean if a field has been set.
func (o *TransferParams) HasCategoryNames() bool {
	if o != nil && !IsNil(o.CategoryNames) {
		return true
	}

	return false
}

// SetCategoryNames gets a reference to the given []string and assigns it to the CategoryNames field.
func (o *TransferParams) SetCategoryNames(v []string) {
	o.CategoryNames = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransferParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransferParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransferParams) SetDescription(v string) {
	o.Description = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *TransferParams) GetFee() TransactionRequestFee {
	if o == nil || IsNil(o.Fee) {
		var ret TransactionRequestFee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferParams) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *TransferParams) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given TransactionRequestFee and assigns it to the Fee field.
func (o *TransferParams) SetFee(v TransactionRequestFee) {
	o.Fee = &v
}

func (o TransferParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["source"] = o.Source
	toSerialize["token_id"] = o.TokenId
	toSerialize["destination"] = o.Destination
	if !IsNil(o.CategoryNames) {
		toSerialize["category_names"] = o.CategoryNames
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	return toSerialize, nil
}

func (o *TransferParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
		"source",
		"token_id",
		"destination",
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

	varTransferParams := _TransferParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransferParams)

	if err != nil {
		return err
	}

	*o = TransferParams(varTransferParams)

	return err
}

type NullableTransferParams struct {
	value *TransferParams
	isSet bool
}

func (v NullableTransferParams) Get() *TransferParams {
	return v.value
}

func (v *NullableTransferParams) Set(val *TransferParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferParams(val *TransferParams) *NullableTransferParams {
	return &NullableTransferParams{value: val, isSet: true}
}

func (v NullableTransferParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


