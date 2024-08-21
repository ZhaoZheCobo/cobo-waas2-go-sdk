/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// WebhookEventData - struct for WebhookEventData
type WebhookEventData struct {
	TSSRequestWebhookEventData *TSSRequestWebhookEventData
	TransactionWebhookEventData *TransactionWebhookEventData
}

// TSSRequestWebhookEventDataAsWebhookEventData is a convenience function that returns TSSRequestWebhookEventData wrapped in WebhookEventData
func TSSRequestWebhookEventDataAsWebhookEventData(v *TSSRequestWebhookEventData) WebhookEventData {
	return WebhookEventData{
		TSSRequestWebhookEventData: v,
	}
}

// TransactionWebhookEventDataAsWebhookEventData is a convenience function that returns TransactionWebhookEventData wrapped in WebhookEventData
func TransactionWebhookEventDataAsWebhookEventData(v *TransactionWebhookEventData) WebhookEventData {
	return WebhookEventData{
		TransactionWebhookEventData: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WebhookEventData) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'TSSRequest'
	if jsonDict["data_type"] == "TSSRequest" {
		// try to unmarshal JSON data into TSSRequestWebhookEventData
		err = json.Unmarshal(data, &dst.TSSRequestWebhookEventData)
		if err == nil {
			return nil // data stored in dst.TSSRequestWebhookEventData, return on the first match
		} else {
			dst.TSSRequestWebhookEventData = nil
			return fmt.Errorf("failed to unmarshal WebhookEventData as TSSRequestWebhookEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Transaction'
	if jsonDict["data_type"] == "Transaction" {
		// try to unmarshal JSON data into TransactionWebhookEventData
		err = json.Unmarshal(data, &dst.TransactionWebhookEventData)
		if err == nil {
			return nil // data stored in dst.TransactionWebhookEventData, return on the first match
		} else {
			dst.TransactionWebhookEventData = nil
			return fmt.Errorf("failed to unmarshal WebhookEventData as TransactionWebhookEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TSSRequestWebhookEventData'
	if jsonDict["data_type"] == "TSSRequestWebhookEventData" {
		// try to unmarshal JSON data into TSSRequestWebhookEventData
		err = json.Unmarshal(data, &dst.TSSRequestWebhookEventData)
		if err == nil {
			return nil // data stored in dst.TSSRequestWebhookEventData, return on the first match
		} else {
			dst.TSSRequestWebhookEventData = nil
			return fmt.Errorf("failed to unmarshal WebhookEventData as TSSRequestWebhookEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionWebhookEventData'
	if jsonDict["data_type"] == "TransactionWebhookEventData" {
		// try to unmarshal JSON data into TransactionWebhookEventData
		err = json.Unmarshal(data, &dst.TransactionWebhookEventData)
		if err == nil {
			return nil // data stored in dst.TransactionWebhookEventData, return on the first match
		} else {
			dst.TransactionWebhookEventData = nil
			return fmt.Errorf("failed to unmarshal WebhookEventData as TransactionWebhookEventData: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WebhookEventData) MarshalJSON() ([]byte, error) {
	if src.TSSRequestWebhookEventData != nil {
		return json.Marshal(&src.TSSRequestWebhookEventData)
	}

	if src.TransactionWebhookEventData != nil {
		return json.Marshal(&src.TransactionWebhookEventData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WebhookEventData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TSSRequestWebhookEventData != nil {
		return obj.TSSRequestWebhookEventData
	}

	if obj.TransactionWebhookEventData != nil {
		return obj.TransactionWebhookEventData
	}

	// all schemas are nil
	return nil
}

type NullableWebhookEventData struct {
	value *WebhookEventData
	isSet bool
}

func (v NullableWebhookEventData) Get() *WebhookEventData {
	return v.value
}

func (v *NullableWebhookEventData) Set(val *WebhookEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventData(val *WebhookEventData) *NullableWebhookEventData {
	return &NullableWebhookEventData{value: val, isSet: true}
}

func (v NullableWebhookEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


