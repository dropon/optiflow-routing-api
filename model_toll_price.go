/*
Routing

With the Routing service you can calculate routes from A to B taking into account vehicle-specific restrictions, traffic situations, toll, emissions, drivers' working hours, service times and opening intervals.

API version: 1.33
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow_routing

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TollPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TollPrice{}

// TollPrice The toll price payable in this country.
type TollPrice struct {
	// The toll price in the specified currency.
	Price float64 `json:"price"`
	// The currency code according to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Currency string `json:"currency"`
}

type _TollPrice TollPrice

// NewTollPrice instantiates a new TollPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTollPrice(price float64, currency string) *TollPrice {
	this := TollPrice{}
	this.Price = price
	this.Currency = currency
	return &this
}

// NewTollPriceWithDefaults instantiates a new TollPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTollPriceWithDefaults() *TollPrice {
	this := TollPrice{}
	return &this
}

// GetPrice returns the Price field value
func (o *TollPrice) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *TollPrice) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *TollPrice) SetPrice(v float64) {
	o.Price = v
}

// GetCurrency returns the Currency field value
func (o *TollPrice) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TollPrice) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TollPrice) SetCurrency(v string) {
	o.Currency = v
}

func (o TollPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TollPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["price"] = o.Price
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *TollPrice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"price",
		"currency",
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

	varTollPrice := _TollPrice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTollPrice)

	if err != nil {
		return err
	}

	*o = TollPrice(varTollPrice)

	return err
}

type NullableTollPrice struct {
	value *TollPrice
	isSet bool
}

func (v NullableTollPrice) Get() *TollPrice {
	return v.value
}

func (v *NullableTollPrice) Set(val *TollPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableTollPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableTollPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTollPrice(val *TollPrice) *NullableTollPrice {
	return &NullableTollPrice{value: val, isSet: true}
}

func (v NullableTollPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTollPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


