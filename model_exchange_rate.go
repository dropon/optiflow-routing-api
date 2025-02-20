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

// checks if the ExchangeRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeRate{}

// ExchangeRate struct for ExchangeRate
type ExchangeRate struct {
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the currency from which a price has been converted.
	Currency string `json:"currency" validate:"regexp=[A-Z]{3}"`
	// The exchange rate to convert a price from the base currency to this currency.
	Rate float64 `json:"rate"`
}

type _ExchangeRate ExchangeRate

// NewExchangeRate instantiates a new ExchangeRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeRate(currency string, rate float64) *ExchangeRate {
	this := ExchangeRate{}
	this.Currency = currency
	this.Rate = rate
	return &this
}

// NewExchangeRateWithDefaults instantiates a new ExchangeRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeRateWithDefaults() *ExchangeRate {
	this := ExchangeRate{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *ExchangeRate) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ExchangeRate) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ExchangeRate) SetCurrency(v string) {
	o.Currency = v
}

// GetRate returns the Rate field value
func (o *ExchangeRate) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *ExchangeRate) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *ExchangeRate) SetRate(v float64) {
	o.Rate = v
}

func (o ExchangeRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["rate"] = o.Rate
	return toSerialize, nil
}

func (o *ExchangeRate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"rate",
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

	varExchangeRate := _ExchangeRate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeRate)

	if err != nil {
		return err
	}

	*o = ExchangeRate(varExchangeRate)

	return err
}

type NullableExchangeRate struct {
	value *ExchangeRate
	isSet bool
}

func (v NullableExchangeRate) Get() *ExchangeRate {
	return v.value
}

func (v *NullableExchangeRate) Set(val *ExchangeRate) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeRate) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeRate(val *ExchangeRate) *NullableExchangeRate {
	return &NullableExchangeRate{value: val, isSet: true}
}

func (v NullableExchangeRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


