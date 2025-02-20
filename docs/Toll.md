# Toll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Costs** | Pointer to [**TollCosts**](TollCosts.md) |  | [optional] 
**Sections** | Pointer to [**[]TollSection**](TollSection.md) | The list of toll sections defined by the toll operators. | [optional] 
**Systems** | Pointer to [**[]TollSystem**](TollSystem.md) | The list of toll systems defined by the toll operators. | [optional] 
**Currencies** | Pointer to [**Currencies**](Currencies.md) |  | [optional] 

## Methods

### NewToll

`func NewToll() *Toll`

NewToll instantiates a new Toll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTollWithDefaults

`func NewTollWithDefaults() *Toll`

NewTollWithDefaults instantiates a new Toll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCosts

`func (o *Toll) GetCosts() TollCosts`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *Toll) GetCostsOk() (*TollCosts, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *Toll) SetCosts(v TollCosts)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *Toll) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetSections

`func (o *Toll) GetSections() []TollSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *Toll) GetSectionsOk() (*[]TollSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *Toll) SetSections(v []TollSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *Toll) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSystems

`func (o *Toll) GetSystems() []TollSystem`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *Toll) GetSystemsOk() (*[]TollSystem, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *Toll) SetSystems(v []TollSystem)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *Toll) HasSystems() bool`

HasSystems returns a boolean if a field has been set.

### GetCurrencies

`func (o *Toll) GetCurrencies() Currencies`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *Toll) GetCurrenciesOk() (*Currencies, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *Toll) SetCurrencies(v Currencies)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *Toll) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


