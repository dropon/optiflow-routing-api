# TollSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Costs** | [**[]TollSectionCost**](TollSectionCost.md) | The toll costs payable for this section. If different prices according to the payment method exist, this list contains one item for each price. The first item contains the price used to calculate the total toll costs for the route. Further items are alternative costs for the section with different payment methods or different subscriptions. | 
**TollRoadType** | [**TollRoadType**](TollRoadType.md) |  | 
**TollSystemIndex** | Pointer to **int32** | The index in the list of toll systems this toll section belongs to. | [optional] 
**CountryCode** | **string** | Countries are represented according to their [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) or [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) if referring to a subdivision. | 
**DisplayName** | Pointer to **string** | A name that describes this toll section. It can relate to the names of the road intersections, to the area in general or to the type of toll. | [optional] 
**OfficialDistance** | Pointer to **int32** | The official distance defined by the toll operator [m].  It may be different from the travel distance. Only present if the toll operator provides an official distance. | [optional] 
**CalculatedDistance** | Pointer to **int32** | The calculated distance of the toll section on the current route.  It may be different from the official distance. | [optional] 
**Approximated** | Pointer to **bool** | States whether the section is approximated because the start or destination waypoint is located inside a toll section and thus the exact toll price cannot be calculated. The closest toll location after the waypoint is used to approximate the toll price. | [optional] 

## Methods

### NewTollSection

`func NewTollSection(costs []TollSectionCost, tollRoadType TollRoadType, countryCode string, ) *TollSection`

NewTollSection instantiates a new TollSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTollSectionWithDefaults

`func NewTollSectionWithDefaults() *TollSection`

NewTollSectionWithDefaults instantiates a new TollSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCosts

`func (o *TollSection) GetCosts() []TollSectionCost`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *TollSection) GetCostsOk() (*[]TollSectionCost, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *TollSection) SetCosts(v []TollSectionCost)`

SetCosts sets Costs field to given value.


### GetTollRoadType

`func (o *TollSection) GetTollRoadType() TollRoadType`

GetTollRoadType returns the TollRoadType field if non-nil, zero value otherwise.

### GetTollRoadTypeOk

`func (o *TollSection) GetTollRoadTypeOk() (*TollRoadType, bool)`

GetTollRoadTypeOk returns a tuple with the TollRoadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollRoadType

`func (o *TollSection) SetTollRoadType(v TollRoadType)`

SetTollRoadType sets TollRoadType field to given value.


### GetTollSystemIndex

`func (o *TollSection) GetTollSystemIndex() int32`

GetTollSystemIndex returns the TollSystemIndex field if non-nil, zero value otherwise.

### GetTollSystemIndexOk

`func (o *TollSection) GetTollSystemIndexOk() (*int32, bool)`

GetTollSystemIndexOk returns a tuple with the TollSystemIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollSystemIndex

`func (o *TollSection) SetTollSystemIndex(v int32)`

SetTollSystemIndex sets TollSystemIndex field to given value.

### HasTollSystemIndex

`func (o *TollSection) HasTollSystemIndex() bool`

HasTollSystemIndex returns a boolean if a field has been set.

### GetCountryCode

`func (o *TollSection) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *TollSection) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *TollSection) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetDisplayName

`func (o *TollSection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TollSection) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TollSection) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TollSection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOfficialDistance

`func (o *TollSection) GetOfficialDistance() int32`

GetOfficialDistance returns the OfficialDistance field if non-nil, zero value otherwise.

### GetOfficialDistanceOk

`func (o *TollSection) GetOfficialDistanceOk() (*int32, bool)`

GetOfficialDistanceOk returns a tuple with the OfficialDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialDistance

`func (o *TollSection) SetOfficialDistance(v int32)`

SetOfficialDistance sets OfficialDistance field to given value.

### HasOfficialDistance

`func (o *TollSection) HasOfficialDistance() bool`

HasOfficialDistance returns a boolean if a field has been set.

### GetCalculatedDistance

`func (o *TollSection) GetCalculatedDistance() int32`

GetCalculatedDistance returns the CalculatedDistance field if non-nil, zero value otherwise.

### GetCalculatedDistanceOk

`func (o *TollSection) GetCalculatedDistanceOk() (*int32, bool)`

GetCalculatedDistanceOk returns a tuple with the CalculatedDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedDistance

`func (o *TollSection) SetCalculatedDistance(v int32)`

SetCalculatedDistance sets CalculatedDistance field to given value.

### HasCalculatedDistance

`func (o *TollSection) HasCalculatedDistance() bool`

HasCalculatedDistance returns a boolean if a field has been set.

### GetApproximated

`func (o *TollSection) GetApproximated() bool`

GetApproximated returns the Approximated field if non-nil, zero value otherwise.

### GetApproximatedOk

`func (o *TollSection) GetApproximatedOk() (*bool, bool)`

GetApproximatedOk returns a tuple with the Approximated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximated

`func (o *TollSection) SetApproximated(v bool)`

SetApproximated sets Approximated field to given value.

### HasApproximated

`func (o *TollSection) HasApproximated() bool`

HasApproximated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


