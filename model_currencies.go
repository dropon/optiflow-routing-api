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

// checks if the Currencies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Currencies{}

// Currencies Information about the currencies that are listed in the toll costs and/or toll sections objects.
type Currencies struct {
	// The date of the exchange rates formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339).
	Date string `json:"date"`
	// The provider of the exchange rates.
	Provider string `json:"provider"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code as provided in the request.
	BaseCurrency string `json:"baseCurrency" validate:"regexp=[A-Z]{3}"`
	// The exchange rates that were used to determine the converted prices.
	ExchangeRates []ExchangeRate `json:"exchangeRates"`
}

type _Currencies Currencies

// NewCurrencies instantiates a new Currencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencies(date string, provider string, baseCurrency string, exchangeRates []ExchangeRate) *Currencies {
	this := Currencies{}
	this.Date = date
	this.Provider = provider
	this.BaseCurrency = baseCurrency
	this.ExchangeRates = exchangeRates
	return &this
}

// NewCurrenciesWithDefaults instantiates a new Currencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrenciesWithDefaults() *Currencies {
	this := Currencies{}
	return &this
}

// GetDate returns the Date field value
func (o *Currencies) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Currencies) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Currencies) SetDate(v string) {
	o.Date = v
}

// GetProvider returns the Provider field value
func (o *Currencies) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Currencies) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *Currencies) SetProvider(v string) {
	o.Provider = v
}

// GetBaseCurrency returns the BaseCurrency field value
func (o *Currencies) GetBaseCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseCurrency
}

// GetBaseCurrencyOk returns a tuple with the BaseCurrency field value
// and a boolean to check if the value has been set.
func (o *Currencies) GetBaseCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCurrency, true
}

// SetBaseCurrency sets field value
func (o *Currencies) SetBaseCurrency(v string) {
	o.BaseCurrency = v
}

// GetExchangeRates returns the ExchangeRates field value
func (o *Currencies) GetExchangeRates() []ExchangeRate {
	if o == nil {
		var ret []ExchangeRate
		return ret
	}

	return o.ExchangeRates
}

// GetExchangeRatesOk returns a tuple with the ExchangeRates field value
// and a boolean to check if the value has been set.
func (o *Currencies) GetExchangeRatesOk() ([]ExchangeRate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExchangeRates, true
}

// SetExchangeRates sets field value
func (o *Currencies) SetExchangeRates(v []ExchangeRate) {
	o.ExchangeRates = v
}

func (o Currencies) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Currencies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["provider"] = o.Provider
	toSerialize["baseCurrency"] = o.BaseCurrency
	toSerialize["exchangeRates"] = o.ExchangeRates
	return toSerialize, nil
}

func (o *Currencies) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"provider",
		"baseCurrency",
		"exchangeRates",
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

	varCurrencies := _Currencies{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCurrencies)

	if err != nil {
		return err
	}

	*o = Currencies(varCurrencies)

	return err
}

type NullableCurrencies struct {
	value *Currencies
	isSet bool
}

func (v NullableCurrencies) Get() *Currencies {
	return v.value
}

func (v *NullableCurrencies) Set(val *Currencies) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencies) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencies(val *Currencies) *NullableCurrencies {
	return &NullableCurrencies{value: val, isSet: true}
}

func (v NullableCurrencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


