/*
 * ShipBob Developer API
 *
 * ShipBob Developer API Documentation
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"time"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// ReturnsApiService ReturnsApi service
type ReturnsApiService service

type ApiReturnGetRequest struct {
	ctx _context.Context
	ApiService *ReturnsApiService
	page *int32
	limit *int32
	sortOrder *ReturnsSortOrder
	startDate *time.Time
	endDate *time.Time
	iDs *[]int32
	referenceIds *[]string
	status *[]ReturnsReturnStatus
	fulfillmentCenterIds *[]int32
	trackingNumbers *[]string
	originalShipmentIds *[]int32
	inventoryIds *[]int32
	shipbobChannelId *int32
}

func (r ApiReturnGetRequest) Page(page int32) ApiReturnGetRequest {
	r.page = &page
	return r
}
func (r ApiReturnGetRequest) Limit(limit int32) ApiReturnGetRequest {
	r.limit = &limit
	return r
}
func (r ApiReturnGetRequest) SortOrder(sortOrder ReturnsSortOrder) ApiReturnGetRequest {
	r.sortOrder = &sortOrder
	return r
}
func (r ApiReturnGetRequest) StartDate(startDate time.Time) ApiReturnGetRequest {
	r.startDate = &startDate
	return r
}
func (r ApiReturnGetRequest) EndDate(endDate time.Time) ApiReturnGetRequest {
	r.endDate = &endDate
	return r
}
func (r ApiReturnGetRequest) IDs(iDs []int32) ApiReturnGetRequest {
	r.iDs = &iDs
	return r
}
func (r ApiReturnGetRequest) ReferenceIds(referenceIds []string) ApiReturnGetRequest {
	r.referenceIds = &referenceIds
	return r
}
func (r ApiReturnGetRequest) Status(status []ReturnsReturnStatus) ApiReturnGetRequest {
	r.status = &status
	return r
}
func (r ApiReturnGetRequest) FulfillmentCenterIds(fulfillmentCenterIds []int32) ApiReturnGetRequest {
	r.fulfillmentCenterIds = &fulfillmentCenterIds
	return r
}
func (r ApiReturnGetRequest) TrackingNumbers(trackingNumbers []string) ApiReturnGetRequest {
	r.trackingNumbers = &trackingNumbers
	return r
}
func (r ApiReturnGetRequest) OriginalShipmentIds(originalShipmentIds []int32) ApiReturnGetRequest {
	r.originalShipmentIds = &originalShipmentIds
	return r
}
func (r ApiReturnGetRequest) InventoryIds(inventoryIds []int32) ApiReturnGetRequest {
	r.inventoryIds = &inventoryIds
	return r
}
func (r ApiReturnGetRequest) ShipbobChannelId(shipbobChannelId int32) ApiReturnGetRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

func (r ApiReturnGetRequest) Execute() ([]ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	return r.ApiService.ReturnGetExecute(r)
}

/*
 * ReturnGet Get Return Orders
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiReturnGetRequest
 */
func (a *ReturnsApiService) ReturnGet(ctx _context.Context) ApiReturnGetRequest {
	return ApiReturnGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []ReturnsReturnOrderViewModel
 */
func (a *ReturnsApiService) ReturnGetExecute(r ApiReturnGetRequest) ([]ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ReturnsReturnOrderViewModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.ReturnGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("Page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("Limit", parameterToString(*r.limit, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("SortOrder", parameterToString(*r.sortOrder, ""))
	}
	if r.startDate != nil {
		localVarQueryParams.Add("StartDate", parameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("EndDate", parameterToString(*r.endDate, ""))
	}
	if r.iDs != nil {
		t := *r.iDs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("IDs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("IDs", parameterToString(t, "multi"))
		}
	}
	if r.referenceIds != nil {
		t := *r.referenceIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("ReferenceIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("ReferenceIds", parameterToString(t, "multi"))
		}
	}
	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("Status", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("Status", parameterToString(t, "multi"))
		}
	}
	if r.fulfillmentCenterIds != nil {
		t := *r.fulfillmentCenterIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("FulfillmentCenterIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("FulfillmentCenterIds", parameterToString(t, "multi"))
		}
	}
	if r.trackingNumbers != nil {
		t := *r.trackingNumbers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("TrackingNumbers", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("TrackingNumbers", parameterToString(t, "multi"))
		}
	}
	if r.originalShipmentIds != nil {
		t := *r.originalShipmentIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("OriginalShipmentIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("OriginalShipmentIds", parameterToString(t, "multi"))
		}
	}
	if r.inventoryIds != nil {
		t := *r.inventoryIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("InventoryIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("InventoryIds", parameterToString(t, "multi"))
		}
	}
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
	if r.shipbobChannelId != nil {
		localVarHeaderParams["shipbob_channel_id"] = parameterToString(*r.shipbobChannelId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnIdCancelPostRequest struct {
	ctx _context.Context
	ApiService *ReturnsApiService
	id int32
	shipbobChannelId *int32
}

func (r ApiReturnIdCancelPostRequest) ShipbobChannelId(shipbobChannelId int32) ApiReturnIdCancelPostRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

func (r ApiReturnIdCancelPostRequest) Execute() (ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	return r.ApiService.ReturnIdCancelPostExecute(r)
}

/*
 * ReturnIdCancelPost Cancel Return Order
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Id of the return order
 * @return ApiReturnIdCancelPostRequest
 */
func (a *ReturnsApiService) ReturnIdCancelPost(ctx _context.Context, id int32) ApiReturnIdCancelPostRequest {
	return ApiReturnIdCancelPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return ReturnsReturnOrderViewModel
 */
func (a *ReturnsApiService) ReturnIdCancelPostExecute(r ApiReturnIdCancelPostRequest) (ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReturnsReturnOrderViewModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.ReturnIdCancelPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return/{id}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.shipbobChannelId == nil {
		return localVarReturnValue, nil, reportError("shipbobChannelId is required and must be specified")
	}

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
	localVarHeaderParams["shipbob_channel_id"] = parameterToString(*r.shipbobChannelId, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnIdGetRequest struct {
	ctx _context.Context
	ApiService *ReturnsApiService
	id int32
	shipbobChannelId *int32
}

func (r ApiReturnIdGetRequest) ShipbobChannelId(shipbobChannelId int32) ApiReturnIdGetRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

func (r ApiReturnIdGetRequest) Execute() (ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	return r.ApiService.ReturnIdGetExecute(r)
}

/*
 * ReturnIdGet Get Return Order
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Id of the return order
 * @return ApiReturnIdGetRequest
 */
func (a *ReturnsApiService) ReturnIdGet(ctx _context.Context, id int32) ApiReturnIdGetRequest {
	return ApiReturnIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return ReturnsReturnOrderViewModel
 */
func (a *ReturnsApiService) ReturnIdGetExecute(r ApiReturnIdGetRequest) (ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReturnsReturnOrderViewModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.ReturnIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if r.shipbobChannelId != nil {
		localVarHeaderParams["shipbob_channel_id"] = parameterToString(*r.shipbobChannelId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnIdPutRequest struct {
	ctx _context.Context
	ApiService *ReturnsApiService
	shipbobChannelId *int32
	id int32
	returnsCreateReturnViewModel *ReturnsCreateReturnViewModel
}

func (r ApiReturnIdPutRequest) ShipbobChannelId(shipbobChannelId int32) ApiReturnIdPutRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}
func (r ApiReturnIdPutRequest) ReturnsCreateReturnViewModel(returnsCreateReturnViewModel ReturnsCreateReturnViewModel) ApiReturnIdPutRequest {
	r.returnsCreateReturnViewModel = &returnsCreateReturnViewModel
	return r
}

func (r ApiReturnIdPutRequest) Execute() (ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	return r.ApiService.ReturnIdPutExecute(r)
}

/*
 * ReturnIdPut Modify Return Order
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Id of the return order
 * @return ApiReturnIdPutRequest
 */
func (a *ReturnsApiService) ReturnIdPut(ctx _context.Context, id int32) ApiReturnIdPutRequest {
	return ApiReturnIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return ReturnsReturnOrderViewModel
 */
func (a *ReturnsApiService) ReturnIdPutExecute(r ApiReturnIdPutRequest) (ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReturnsReturnOrderViewModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.ReturnIdPut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.shipbobChannelId == nil {
		return localVarReturnValue, nil, reportError("shipbobChannelId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/_*+json", "application/json", "application/json-patch+json", "text/json"}

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
	localVarHeaderParams["shipbob_channel_id"] = parameterToString(*r.shipbobChannelId, "")
	// body params
	localVarPostBody = r.returnsCreateReturnViewModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnIdStatushistoryGetRequest struct {
	ctx _context.Context
	ApiService *ReturnsApiService
	id int32
	shipbobChannelId *int32
}

func (r ApiReturnIdStatushistoryGetRequest) ShipbobChannelId(shipbobChannelId int32) ApiReturnIdStatushistoryGetRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

func (r ApiReturnIdStatushistoryGetRequest) Execute() (ReturnsReturnOrderStatusHistoryViewModel, *_nethttp.Response, error) {
	return r.ApiService.ReturnIdStatushistoryGetExecute(r)
}

/*
 * ReturnIdStatushistoryGet Get One Return's status history
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Id of the return order
 * @return ApiReturnIdStatushistoryGetRequest
 */
func (a *ReturnsApiService) ReturnIdStatushistoryGet(ctx _context.Context, id int32) ApiReturnIdStatushistoryGetRequest {
	return ApiReturnIdStatushistoryGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return ReturnsReturnOrderStatusHistoryViewModel
 */
func (a *ReturnsApiService) ReturnIdStatushistoryGetExecute(r ApiReturnIdStatushistoryGetRequest) (ReturnsReturnOrderStatusHistoryViewModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReturnsReturnOrderStatusHistoryViewModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.ReturnIdStatushistoryGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return/{id}/statushistory"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if r.shipbobChannelId != nil {
		localVarHeaderParams["shipbob_channel_id"] = parameterToString(*r.shipbobChannelId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnPostRequest struct {
	ctx _context.Context
	ApiService *ReturnsApiService
	shipbobChannelId *int32
	returnsCreateReturnViewModel *ReturnsCreateReturnViewModel
}

func (r ApiReturnPostRequest) ShipbobChannelId(shipbobChannelId int32) ApiReturnPostRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}
func (r ApiReturnPostRequest) ReturnsCreateReturnViewModel(returnsCreateReturnViewModel ReturnsCreateReturnViewModel) ApiReturnPostRequest {
	r.returnsCreateReturnViewModel = &returnsCreateReturnViewModel
	return r
}

func (r ApiReturnPostRequest) Execute() (ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	return r.ApiService.ReturnPostExecute(r)
}

/*
 * ReturnPost Create Return Order
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiReturnPostRequest
 */
func (a *ReturnsApiService) ReturnPost(ctx _context.Context) ApiReturnPostRequest {
	return ApiReturnPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ReturnsReturnOrderViewModel
 */
func (a *ReturnsApiService) ReturnPostExecute(r ApiReturnPostRequest) (ReturnsReturnOrderViewModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReturnsReturnOrderViewModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.ReturnPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.shipbobChannelId == nil {
		return localVarReturnValue, nil, reportError("shipbobChannelId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/_*+json", "application/json", "application/json-patch+json", "text/json"}

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
	localVarHeaderParams["shipbob_channel_id"] = parameterToString(*r.shipbobChannelId, "")
	// body params
	localVarPostBody = r.returnsCreateReturnViewModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
