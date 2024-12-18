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

// checks if the Activity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity{}

// Activity The staking activity.
type Activity struct {
	// The activity ID.
	Id *string `json:"id,omitempty"`
	// The initiator of the activity.
	Initiator *string `json:"initiator,omitempty"`
	InitiatorType *TransactionInitiatorType `json:"initiator_type,omitempty"`
	Type *ActivityType `json:"type,omitempty"`
	// The staker's wallet ID.
	WalletId *string `json:"wallet_id,omitempty"`
	// The staker's wallet address.
	Address *string `json:"address,omitempty"`
	PoolId StakingPoolId `json:"pool_id"`
	// The token ID.
	TokenId string `json:"token_id"`
	// The ID of the corresponding staking position.
	StakingId *string `json:"staking_id,omitempty"`
	// The staking amount.
	Amount string `json:"amount"`
	// The IDs of the corresponding transactions of the activity.
	TransactionIds []string `json:"transaction_ids,omitempty"`
	// The timeline of the activity.
	Timeline []ActivityTimeline `json:"timeline,omitempty"`
	Fee *TransactionRequestFee `json:"fee,omitempty"`
	Status ActivityStatus `json:"status"`
	// The time when the activity was created.
	CreatedTimestamp *int64 `json:"created_timestamp,omitempty"`
	// The time when the activity was last updated.
	UpdatedTimestamp *int64 `json:"updated_timestamp,omitempty"`
}

type _Activity Activity

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity(poolId StakingPoolId, tokenId string, amount string, status ActivityStatus) *Activity {
	this := Activity{}
	this.PoolId = poolId
	this.TokenId = tokenId
	this.Amount = amount
	this.Status = status
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Activity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Activity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Activity) SetId(v string) {
	o.Id = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *Activity) GetInitiator() string {
	if o == nil || IsNil(o.Initiator) {
		var ret string
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetInitiatorOk() (*string, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *Activity) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given string and assigns it to the Initiator field.
func (o *Activity) SetInitiator(v string) {
	o.Initiator = &v
}

// GetInitiatorType returns the InitiatorType field value if set, zero value otherwise.
func (o *Activity) GetInitiatorType() TransactionInitiatorType {
	if o == nil || IsNil(o.InitiatorType) {
		var ret TransactionInitiatorType
		return ret
	}
	return *o.InitiatorType
}

// GetInitiatorTypeOk returns a tuple with the InitiatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetInitiatorTypeOk() (*TransactionInitiatorType, bool) {
	if o == nil || IsNil(o.InitiatorType) {
		return nil, false
	}
	return o.InitiatorType, true
}

// HasInitiatorType returns a boolean if a field has been set.
func (o *Activity) HasInitiatorType() bool {
	if o != nil && !IsNil(o.InitiatorType) {
		return true
	}

	return false
}

// SetInitiatorType gets a reference to the given TransactionInitiatorType and assigns it to the InitiatorType field.
func (o *Activity) SetInitiatorType(v TransactionInitiatorType) {
	o.InitiatorType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Activity) GetType() ActivityType {
	if o == nil || IsNil(o.Type) {
		var ret ActivityType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetTypeOk() (*ActivityType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Activity) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ActivityType and assigns it to the Type field.
func (o *Activity) SetType(v ActivityType) {
	o.Type = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *Activity) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *Activity) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *Activity) SetWalletId(v string) {
	o.WalletId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Activity) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Activity) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Activity) SetAddress(v string) {
	o.Address = &v
}

// GetPoolId returns the PoolId field value
func (o *Activity) GetPoolId() StakingPoolId {
	if o == nil {
		var ret StakingPoolId
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *Activity) GetPoolIdOk() (*StakingPoolId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *Activity) SetPoolId(v StakingPoolId) {
	o.PoolId = v
}

// GetTokenId returns the TokenId field value
func (o *Activity) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *Activity) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *Activity) SetTokenId(v string) {
	o.TokenId = v
}

// GetStakingId returns the StakingId field value if set, zero value otherwise.
func (o *Activity) GetStakingId() string {
	if o == nil || IsNil(o.StakingId) {
		var ret string
		return ret
	}
	return *o.StakingId
}

// GetStakingIdOk returns a tuple with the StakingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetStakingIdOk() (*string, bool) {
	if o == nil || IsNil(o.StakingId) {
		return nil, false
	}
	return o.StakingId, true
}

// HasStakingId returns a boolean if a field has been set.
func (o *Activity) HasStakingId() bool {
	if o != nil && !IsNil(o.StakingId) {
		return true
	}

	return false
}

// SetStakingId gets a reference to the given string and assigns it to the StakingId field.
func (o *Activity) SetStakingId(v string) {
	o.StakingId = &v
}

// GetAmount returns the Amount field value
func (o *Activity) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Activity) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Activity) SetAmount(v string) {
	o.Amount = v
}

// GetTransactionIds returns the TransactionIds field value if set, zero value otherwise.
func (o *Activity) GetTransactionIds() []string {
	if o == nil || IsNil(o.TransactionIds) {
		var ret []string
		return ret
	}
	return o.TransactionIds
}

// GetTransactionIdsOk returns a tuple with the TransactionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetTransactionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TransactionIds) {
		return nil, false
	}
	return o.TransactionIds, true
}

// HasTransactionIds returns a boolean if a field has been set.
func (o *Activity) HasTransactionIds() bool {
	if o != nil && !IsNil(o.TransactionIds) {
		return true
	}

	return false
}

// SetTransactionIds gets a reference to the given []string and assigns it to the TransactionIds field.
func (o *Activity) SetTransactionIds(v []string) {
	o.TransactionIds = v
}

// GetTimeline returns the Timeline field value if set, zero value otherwise.
func (o *Activity) GetTimeline() []ActivityTimeline {
	if o == nil || IsNil(o.Timeline) {
		var ret []ActivityTimeline
		return ret
	}
	return o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetTimelineOk() ([]ActivityTimeline, bool) {
	if o == nil || IsNil(o.Timeline) {
		return nil, false
	}
	return o.Timeline, true
}

// HasTimeline returns a boolean if a field has been set.
func (o *Activity) HasTimeline() bool {
	if o != nil && !IsNil(o.Timeline) {
		return true
	}

	return false
}

// SetTimeline gets a reference to the given []ActivityTimeline and assigns it to the Timeline field.
func (o *Activity) SetTimeline(v []ActivityTimeline) {
	o.Timeline = v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *Activity) GetFee() TransactionRequestFee {
	if o == nil || IsNil(o.Fee) {
		var ret TransactionRequestFee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *Activity) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given TransactionRequestFee and assigns it to the Fee field.
func (o *Activity) SetFee(v TransactionRequestFee) {
	o.Fee = &v
}

// GetStatus returns the Status field value
func (o *Activity) GetStatus() ActivityStatus {
	if o == nil {
		var ret ActivityStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Activity) GetStatusOk() (*ActivityStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Activity) SetStatus(v ActivityStatus) {
	o.Status = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *Activity) GetCreatedTimestamp() int64 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *Activity) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int64 and assigns it to the CreatedTimestamp field.
func (o *Activity) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value if set, zero value otherwise.
func (o *Activity) GetUpdatedTimestamp() int64 {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetUpdatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		return nil, false
	}
	return o.UpdatedTimestamp, true
}

// HasUpdatedTimestamp returns a boolean if a field has been set.
func (o *Activity) HasUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UpdatedTimestamp) {
		return true
	}

	return false
}

// SetUpdatedTimestamp gets a reference to the given int64 and assigns it to the UpdatedTimestamp field.
func (o *Activity) SetUpdatedTimestamp(v int64) {
	o.UpdatedTimestamp = &v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	if !IsNil(o.InitiatorType) {
		toSerialize["initiator_type"] = o.InitiatorType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["pool_id"] = o.PoolId
	toSerialize["token_id"] = o.TokenId
	if !IsNil(o.StakingId) {
		toSerialize["staking_id"] = o.StakingId
	}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.TransactionIds) {
		toSerialize["transaction_ids"] = o.TransactionIds
	}
	if !IsNil(o.Timeline) {
		toSerialize["timeline"] = o.Timeline
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.UpdatedTimestamp) {
		toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	}
	return toSerialize, nil
}

func (o *Activity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pool_id",
		"token_id",
		"amount",
		"status",
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

	varActivity := _Activity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivity)

	if err != nil {
		return err
	}

	*o = Activity(varActivity)

	return err
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


