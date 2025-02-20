# ReachableAreas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Polygons** | **[]string** | The list of polygons calculated for the specified horizons in GeoJson format. For each horizon there is a separate polygon at the same index. | 
**Warnings** | Pointer to [**[]Warning**](Warning.md) | A list of warnings concerning the validity of the result. | [optional] 

## Methods

### NewReachableAreas

`func NewReachableAreas(polygons []string, ) *ReachableAreas`

NewReachableAreas instantiates a new ReachableAreas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachableAreasWithDefaults

`func NewReachableAreasWithDefaults() *ReachableAreas`

NewReachableAreasWithDefaults instantiates a new ReachableAreas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolygons

`func (o *ReachableAreas) GetPolygons() []string`

GetPolygons returns the Polygons field if non-nil, zero value otherwise.

### GetPolygonsOk

`func (o *ReachableAreas) GetPolygonsOk() (*[]string, bool)`

GetPolygonsOk returns a tuple with the Polygons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolygons

`func (o *ReachableAreas) SetPolygons(v []string)`

SetPolygons sets Polygons field to given value.


### GetWarnings

`func (o *ReachableAreas) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ReachableAreas) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ReachableAreas) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ReachableAreas) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


