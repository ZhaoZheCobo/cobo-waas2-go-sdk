/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the ListTokenBalancesForAddress200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTokenBalancesForAddress200Response{}

// ListTokenBalancesForAddress200Response struct for ListTokenBalancesForAddress200Response
type ListTokenBalancesForAddress200Response struct {
	Data []TokenBalance `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListTokenBalancesForAddress200Response instantiates a new ListTokenBalancesForAddress200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTokenBalancesForAddress200Response() *ListTokenBalancesForAddress200Response {
	this := ListTokenBalancesForAddress200Response{}
	return &this
}

// NewListTokenBalancesForAddress200ResponseWithDefaults instantiates a new ListTokenBalancesForAddress200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTokenBalancesForAddress200ResponseWithDefaults() *ListTokenBalancesForAddress200Response {
	this := ListTokenBalancesForAddress200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListTokenBalancesForAddress200Response) GetData() []TokenBalance {
	if o == nil || IsNil(o.Data) {
		var ret []TokenBalance
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTokenBalancesForAddress200Response) GetDataOk() ([]TokenBalance, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListTokenBalancesForAddress200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TokenBalance and assigns it to the Data field.
func (o *ListTokenBalancesForAddress200Response) SetData(v []TokenBalance) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListTokenBalancesForAddress200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTokenBalancesForAddress200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListTokenBalancesForAddress200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListTokenBalancesForAddress200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListTokenBalancesForAddress200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTokenBalancesForAddress200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListTokenBalancesForAddress200Response struct {
	value *ListTokenBalancesForAddress200Response
	isSet bool
}

func (v NullableListTokenBalancesForAddress200Response) Get() *ListTokenBalancesForAddress200Response {
	return v.value
}

func (v *NullableListTokenBalancesForAddress200Response) Set(val *ListTokenBalancesForAddress200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListTokenBalancesForAddress200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListTokenBalancesForAddress200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTokenBalancesForAddress200Response(val *ListTokenBalancesForAddress200Response) *NullableListTokenBalancesForAddress200Response {
	return &NullableListTokenBalancesForAddress200Response{value: val, isSet: true}
}

func (v NullableListTokenBalancesForAddress200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTokenBalancesForAddress200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


