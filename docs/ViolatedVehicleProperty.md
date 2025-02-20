# ViolatedVehicleProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | The name of the violated vehicle property.  The following values represent a vehicle property which is greater than the given **limit**: \&quot;WEIGHT\&quot;, \&quot;HEIGHT\&quot;, \&quot;LENGTH\&quot;, \&quot;AXLE_WEIGHT\&quot;, \&quot;WIDTH\&quot;, \&quot;KPRA_LENGTH\&quot;, \&quot;NUMBER_OF_AXLES\&quot;, \&quot;TOTAL_PERMITTED_WEIGHT\&quot;.  The following values represent a vehicle property for which **value** contains additional information: \&quot;LOW_EMISSION_ZONE\&quot;, \&quot;TRUCK_ROUTE\&quot;, \&quot;HAZARDOUS_MATERIALS\&quot;, \&quot;TUNNEL_RESTRICTION\&quot;, \&quot;TRAILER\&quot;.  This list can be extended at any time, clients should handle unknown values properly.  See [here](./concepts/violations) for more information. | [optional] 
**Limit** | Pointer to **int32** | If the **property** represents a vehicle dimension such as _WEIGHT_ or _HEIGHT_, this value contains the limit which was exceeded by the vehicle. | [optional] 
**Value** | Pointer to **string** | If the **property** is:  * _LOW_EMISSION_ZONE_, this value contains the name of the low-emission zone, see the corresponding low-emission zone event for details on what is required to enter this zone, * _TRUCK_ROUTE_, this value contains the name of the truck route prohibited for the selected vehicle, * _HAZARDOUS_MATERIALS_, this value contains a comma-separated list of the prohibited hazardous materials. * _TUNNEL_RESTRICTION_, this value represents the allowed tunnel restriction code. * _TRAILER_, this value represents that a trailer is not allowed. | [optional] 

## Methods

### NewViolatedVehicleProperty

`func NewViolatedVehicleProperty() *ViolatedVehicleProperty`

NewViolatedVehicleProperty instantiates a new ViolatedVehicleProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViolatedVehiclePropertyWithDefaults

`func NewViolatedVehiclePropertyWithDefaults() *ViolatedVehicleProperty`

NewViolatedVehiclePropertyWithDefaults instantiates a new ViolatedVehicleProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ViolatedVehicleProperty) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ViolatedVehicleProperty) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ViolatedVehicleProperty) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ViolatedVehicleProperty) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetLimit

`func (o *ViolatedVehicleProperty) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ViolatedVehicleProperty) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ViolatedVehicleProperty) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ViolatedVehicleProperty) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetValue

`func (o *ViolatedVehicleProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ViolatedVehicleProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ViolatedVehicleProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ViolatedVehicleProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


