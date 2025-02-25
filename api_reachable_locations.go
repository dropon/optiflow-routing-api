/*
Routing

With the Routing service you can calculate routes from A to B taking into account vehicle-specific restrictions, traffic situations, toll, emissions, drivers' working hours, service times and opening intervals.

API version: 1.33
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow_routing

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ReachableLocationsAPIService ReachableLocationsAPI service
type ReachableLocationsAPIService service

type ApiDeleteReachableLocationsRequest struct {
	ctx context.Context
	ApiService *ReachableLocationsAPIService
	id string
}

func (r ApiDeleteReachableLocationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteReachableLocationsExecute(r)
}

/*
DeleteReachableLocations Method for DeleteReachableLocations

Cancels a reachable locations calculation and deletes the calculated results specified by its ID. Results already calculated cannot be requested by its ID, anymore.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the calculated reachable locations.
 @return ApiDeleteReachableLocationsRequest
*/
func (a *ReachableLocationsAPIService) DeleteReachableLocations(ctx context.Context, id string) ApiDeleteReachableLocationsRequest {
	return ApiDeleteReachableLocationsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ReachableLocationsAPIService) DeleteReachableLocationsExecute(r ApiDeleteReachableLocationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachableLocationsAPIService.DeleteReachableLocations")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reachable-locations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["apiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetReachableLocationsRequest struct {
	ctx context.Context
	ApiService *ReachableLocationsAPIService
	id string
}

func (r ApiGetReachableLocationsRequest) Execute() (*ReachableLocationsResponse, *http.Response, error) {
	return r.ApiService.GetReachableLocationsExecute(r)
}

/*
GetReachableLocations Method for GetReachableLocations

Gets the results of a reachable locations calculation specified by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the calculated reachable locations.
 @return ApiGetReachableLocationsRequest
*/
func (a *ReachableLocationsAPIService) GetReachableLocations(ctx context.Context, id string) ApiGetReachableLocationsRequest {
	return ApiGetReachableLocationsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ReachableLocationsResponse
func (a *ReachableLocationsAPIService) GetReachableLocationsExecute(r ApiGetReachableLocationsRequest) (*ReachableLocationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReachableLocationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachableLocationsAPIService.GetReachableLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reachable-locations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["apiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStartAndCreateReachableLocationsRequest struct {
	ctx context.Context
	ApiService *ReachableLocationsAPIService
	horizon *int32
	locations *Locations
	waypoint *string
	routeId *string
	profile *string
	horizonType *HorizonType
	options *ReachableOptions
}

// The distance [m] or travel time [s] of the horizons, depending of the **horizonType** (limited to 100 km or 1 hours).
func (r ApiStartAndCreateReachableLocationsRequest) Horizon(horizon int32) ApiStartAndCreateReachableLocationsRequest {
	r.horizon = &horizon
	return r
}

func (r ApiStartAndCreateReachableLocationsRequest) Locations(locations Locations) ApiStartAndCreateReachableLocationsRequest {
	r.locations = &locations
	return r
}

// The start or destination waypoint. The format of the waypoint is &#x60;&lt;lat&gt;,&lt;lon&gt;[;&lt;attribute&gt;;&lt;attribute&gt;;...]&#x60; representing a point with the latitude value in degrees from south to north and the longitude value in degrees (WGS84/EPSG:4326) from west to east. This point will be matched to the nearest possible road. By default the air-line connection between given and matched coordinates is not included in the distance or duration. We will refer to this type of waypoint as an _on-road waypoint_.  The behaviour of a waypoint can be changed by appending the following attributes:   * &#x60;includeLastMeters&#x60; to include the air-line connection between given and matched coordinates in the distance or duration.   We will refer to this type of waypoint as an _off-road waypoint_.   * &#x60;roadAccess&#x3D;&lt;lat&gt;,&lt;lon&gt;&#x60;, to use these coordinates for matching to the nearest road. Implies **includeLastMeters**, i.e.   the air-line connection between the waypoint coordinates and the matched coordinates   is included in the distance or duration. This is useful if the waypoint should not be matched to the nearest possible road but to some road further away,   e.g. garage exit at a different road.  See [here](./concepts/waypoints) for more information.
func (r ApiStartAndCreateReachableLocationsRequest) Waypoint(waypoint string) ApiStartAndCreateReachableLocationsRequest {
	r.waypoint = &waypoint
	return r
}

// Instead of the waypoint mentioned above, a **routeId** from a previously calculated route or a matched track can be entered. More information and applying restrictions can be found [here](./concepts/waypoints).
func (r ApiStartAndCreateReachableLocationsRequest) RouteId(routeId string) ApiStartAndCreateReachableLocationsRequest {
	r.routeId = &routeId
	return r
}

// A profile defines a vehicle by a set of attributes, matching typical transport situations. It must be the name of one of the [predefined profiles](../data-api/concepts/profiles) such as _EUR_TRAILER_TRUCK_.  If this parameter is not specified and the first waypoint or the routeId is located in the Americas, _USA_8_SEMITRAILER_5AXLE_ is used as the default instead of _EUR_TRAILER_TRUCK_.  If the first waypoint or the routeId is located in the Americas but a non-American profile is specified or vice-versa, a warning is returned (routing only). Always use a profile which matches the region of the waypoints to obtain best results.  If the attributes of the profile do not fit to your vehicle, the values can be changed by the corresponding attributes in the **vehicle** parameter (routing only).  The values of the predefined profiles may be adapted to reflect current vehicle standards. To obtain the same results when values change, it is recommended to  always send with the request the **vehicle** parameters that are important for your use case.
func (r ApiStartAndCreateReachableLocationsRequest) Profile(profile string) ApiStartAndCreateReachableLocationsRequest {
	r.profile = &profile
	return r
}

func (r ApiStartAndCreateReachableLocationsRequest) HorizonType(horizonType HorizonType) ApiStartAndCreateReachableLocationsRequest {
	r.horizonType = &horizonType
	return r
}

// Routing-relevant options like driving direction or the use of additional data. Use array notation like &#x60;options[trafficMode]&#x3D;AVERAGE&#x60; to set options.
func (r ApiStartAndCreateReachableLocationsRequest) Options(options ReachableOptions) ApiStartAndCreateReachableLocationsRequest {
	r.options = &options
	return r
}

func (r ApiStartAndCreateReachableLocationsRequest) Execute() (*ReachableLocationsId, *http.Response, error) {
	return r.ApiService.StartAndCreateReachableLocationsExecute(r)
}

/*
StartAndCreateReachableLocations Method for StartAndCreateReachableLocations

Starts the calculation of the sets of reachable and unreachable locations from the given ones and creates them as the result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStartAndCreateReachableLocationsRequest
*/
func (a *ReachableLocationsAPIService) StartAndCreateReachableLocations(ctx context.Context) ApiStartAndCreateReachableLocationsRequest {
	return ApiStartAndCreateReachableLocationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReachableLocationsId
func (a *ReachableLocationsAPIService) StartAndCreateReachableLocationsExecute(r ApiStartAndCreateReachableLocationsRequest) (*ReachableLocationsId, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReachableLocationsId
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachableLocationsAPIService.StartAndCreateReachableLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reachable-locations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.horizon == nil {
		return localVarReturnValue, nil, reportError("horizon is required and must be specified")
	}
	if *r.horizon < 1 {
		return localVarReturnValue, nil, reportError("horizon must be greater than 1")
	}
	if r.locations == nil {
		return localVarReturnValue, nil, reportError("locations is required and must be specified")
	}

	if r.waypoint != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "waypoint", r.waypoint, "form", "")
	}
	if r.routeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "routeId", r.routeId, "form", "")
	}
	if r.profile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "profile", r.profile, "form", "")
	} else {
		var defaultValue string = "EUR_TRAILER_TRUCK"
		r.profile = &defaultValue
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "horizon", r.horizon, "form", "")
	if r.horizonType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "horizonType", r.horizonType, "form", "")
	} else {
		var defaultValue HorizonType = "TRAVEL_TIME"
		r.horizonType = &defaultValue
	}
	if r.options != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "options", r.options, "deepObject", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.locations
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["apiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
