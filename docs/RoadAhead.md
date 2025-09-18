# RoadAhead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the road. It does not depend on the selected language, instead it is presented as it can be found on local signs. | [optional] 
**Number** | Pointer to **string** | The number of the road which may consist of several numbers separated by \&quot;/\&quot;. | [optional] 

## Methods

### NewRoadAhead

`func NewRoadAhead() *RoadAhead`

NewRoadAhead instantiates a new RoadAhead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoadAheadWithDefaults

`func NewRoadAheadWithDefaults() *RoadAhead`

NewRoadAheadWithDefaults instantiates a new RoadAhead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoadAhead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoadAhead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoadAhead) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoadAhead) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *RoadAhead) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *RoadAhead) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *RoadAhead) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *RoadAhead) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


