/*
ShipBob Developer API

ShipBob Developer API Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"reflect"
)


// ReturnsApiService ReturnsApi service
type ReturnsApiService service

type ApiCancelReturnRequest struct {
	ctx context.Context
	ApiService *ReturnsApiService
	id int32
	shipbobChannelId *int32
}

// Channel Id for Operation
func (r ApiCancelReturnRequest) ShipbobChannelId(shipbobChannelId int32) ApiCancelReturnRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

func (r ApiCancelReturnRequest) Execute() (*ReturnOrder, *http.Response, error) {
	return r.ApiService.CancelReturnExecute(r)
}

/*
CancelReturn Cancel Return Order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id of the return order
 @return ApiCancelReturnRequest
*/
func (a *ReturnsApiService) CancelReturn(ctx context.Context, id int32) ApiCancelReturnRequest {
	return ApiCancelReturnRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ReturnOrder
func (a *ReturnsApiService) CancelReturnExecute(r ApiCancelReturnRequest) (*ReturnOrder, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnOrder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.CancelReturn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1.0/return/{id}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "shipbob_channel_id", r.shipbobChannelId, "")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
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

type ApiCreateReturnRequest struct {
	ctx context.Context
	ApiService *ReturnsApiService
	shipbobChannelId *int32
	createReturn *CreateReturn
}

// Channel Id for Operation
func (r ApiCreateReturnRequest) ShipbobChannelId(shipbobChannelId int32) ApiCreateReturnRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

// Model defining the return
func (r ApiCreateReturnRequest) CreateReturn(createReturn CreateReturn) ApiCreateReturnRequest {
	r.createReturn = &createReturn
	return r
}

func (r ApiCreateReturnRequest) Execute() (*ReturnOrder, *http.Response, error) {
	return r.ApiService.CreateReturnExecute(r)
}

/*
CreateReturn Create Return Order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateReturnRequest
*/
func (a *ReturnsApiService) CreateReturn(ctx context.Context) ApiCreateReturnRequest {
	return ApiCreateReturnRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnOrder
func (a *ReturnsApiService) CreateReturnExecute(r ApiCreateReturnRequest) (*ReturnOrder, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnOrder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.CreateReturn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1.0/return"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shipbobChannelId == nil {
		return localVarReturnValue, nil, reportError("shipbobChannelId is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "shipbob_channel_id", r.shipbobChannelId, "")
	// body params
	localVarPostBody = r.createReturn
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
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

type ApiGetReturnRequest struct {
	ctx context.Context
	ApiService *ReturnsApiService
	id int32
	shipbobChannelId *int32
}

// Channel Id for Operation
func (r ApiGetReturnRequest) ShipbobChannelId(shipbobChannelId int32) ApiGetReturnRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

func (r ApiGetReturnRequest) Execute() (*ReturnOrder, *http.Response, error) {
	return r.ApiService.GetReturnExecute(r)
}

/*
GetReturn Get Return Order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id of the return order
 @return ApiGetReturnRequest
*/
func (a *ReturnsApiService) GetReturn(ctx context.Context, id int32) ApiGetReturnRequest {
	return ApiGetReturnRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ReturnOrder
func (a *ReturnsApiService) GetReturnExecute(r ApiGetReturnRequest) (*ReturnOrder, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnOrder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.GetReturn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1.0/return/{id}"
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
	if r.shipbobChannelId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "shipbob_channel_id", r.shipbobChannelId, "")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
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

type ApiGetReturnStatushistoryRequest struct {
	ctx context.Context
	ApiService *ReturnsApiService
	id int32
	shipbobChannelId *int32
}

// Channel Id for Operation
func (r ApiGetReturnStatushistoryRequest) ShipbobChannelId(shipbobChannelId int32) ApiGetReturnStatushistoryRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

func (r ApiGetReturnStatushistoryRequest) Execute() ([]ReturnOrderStatusHistory, *http.Response, error) {
	return r.ApiService.GetReturnStatushistoryExecute(r)
}

/*
GetReturnStatushistory Get One Return's status history

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id of the return order
 @return ApiGetReturnStatushistoryRequest
*/
func (a *ReturnsApiService) GetReturnStatushistory(ctx context.Context, id int32) ApiGetReturnStatushistoryRequest {
	return ApiGetReturnStatushistoryRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []ReturnOrderStatusHistory
func (a *ReturnsApiService) GetReturnStatushistoryExecute(r ApiGetReturnStatushistoryRequest) ([]ReturnOrderStatusHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ReturnOrderStatusHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.GetReturnStatushistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1.0/return/{id}/statushistory"
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
	if r.shipbobChannelId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "shipbob_channel_id", r.shipbobChannelId, "")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
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

type ApiGetReturnsRequest struct {
	ctx context.Context
	ApiService *ReturnsApiService
	page *int32
	limit *int32
	sortOrder *SortOrder
	startDate *time.Time
	endDate *time.Time
	iDs *[]int32
	referenceIds *[]string
	status *[]ReturnStatus
	fulfillmentCenterIds *[]int32
	trackingNumbers *[]string
	originalShipmentIds *[]int32
	inventoryIds *[]int32
	shipbobChannelId *int32
}

// Page of return orders to get
func (r ApiGetReturnsRequest) Page(page int32) ApiGetReturnsRequest {
	r.page = &page
	return r
}

// Amount of return orders per page to request
func (r ApiGetReturnsRequest) Limit(limit int32) ApiGetReturnsRequest {
	r.limit = &limit
	return r
}

// Order to sort results by
func (r ApiGetReturnsRequest) SortOrder(sortOrder SortOrder) ApiGetReturnsRequest {
	r.sortOrder = &sortOrder
	return r
}

// Start date to filter orders inserted later than
func (r ApiGetReturnsRequest) StartDate(startDate time.Time) ApiGetReturnsRequest {
	r.startDate = &startDate
	return r
}

// End date to filter orders inserted earlier than
func (r ApiGetReturnsRequest) EndDate(endDate time.Time) ApiGetReturnsRequest {
	r.endDate = &endDate
	return r
}

// Comma separated list of return orders ids to filter by
func (r ApiGetReturnsRequest) IDs(iDs []int32) ApiGetReturnsRequest {
	r.iDs = &iDs
	return r
}

// Comma separated list of reference ids to filter by
func (r ApiGetReturnsRequest) ReferenceIds(referenceIds []string) ApiGetReturnsRequest {
	r.referenceIds = &referenceIds
	return r
}

// Comma separated list of statuses to filter by
func (r ApiGetReturnsRequest) Status(status []ReturnStatus) ApiGetReturnsRequest {
	r.status = &status
	return r
}

// Comma separated list of destination fulfillment center IDs to filter by
func (r ApiGetReturnsRequest) FulfillmentCenterIds(fulfillmentCenterIds []int32) ApiGetReturnsRequest {
	r.fulfillmentCenterIds = &fulfillmentCenterIds
	return r
}

// Comma separated list of tracking numbers to filter by
func (r ApiGetReturnsRequest) TrackingNumbers(trackingNumbers []string) ApiGetReturnsRequest {
	r.trackingNumbers = &trackingNumbers
	return r
}

// Comma separated list of original shipment IDs to filter by
func (r ApiGetReturnsRequest) OriginalShipmentIds(originalShipmentIds []int32) ApiGetReturnsRequest {
	r.originalShipmentIds = &originalShipmentIds
	return r
}

// Comma separated list of inventory IDs contained in return to filter by
func (r ApiGetReturnsRequest) InventoryIds(inventoryIds []int32) ApiGetReturnsRequest {
	r.inventoryIds = &inventoryIds
	return r
}

func (r ApiGetReturnsRequest) ShipbobChannelId(shipbobChannelId int32) ApiGetReturnsRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

func (r ApiGetReturnsRequest) Execute() ([]ReturnOrder, *http.Response, error) {
	return r.ApiService.GetReturnsExecute(r)
}

/*
GetReturns Get Return Orders

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetReturnsRequest
*/
func (a *ReturnsApiService) GetReturns(ctx context.Context) ApiGetReturnsRequest {
	return ApiGetReturnsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ReturnOrder
func (a *ReturnsApiService) GetReturnsExecute(r ApiGetReturnsRequest) ([]ReturnOrder, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ReturnOrder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.GetReturns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1.0/return"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Page", r.page, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Limit", r.limit, "")
	}
	if r.sortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SortOrder", r.sortOrder, "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "StartDate", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EndDate", r.endDate, "")
	}
	if r.iDs != nil {
		t := *r.iDs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "IDs", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "IDs", t, "multi")
		}
	}
	if r.referenceIds != nil {
		t := *r.referenceIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "ReferenceIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "ReferenceIds", t, "multi")
		}
	}
	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "Status", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "Status", t, "multi")
		}
	}
	if r.fulfillmentCenterIds != nil {
		t := *r.fulfillmentCenterIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "FulfillmentCenterIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "FulfillmentCenterIds", t, "multi")
		}
	}
	if r.trackingNumbers != nil {
		t := *r.trackingNumbers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "TrackingNumbers", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "TrackingNumbers", t, "multi")
		}
	}
	if r.originalShipmentIds != nil {
		t := *r.originalShipmentIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "OriginalShipmentIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "OriginalShipmentIds", t, "multi")
		}
	}
	if r.inventoryIds != nil {
		t := *r.inventoryIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "InventoryIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "InventoryIds", t, "multi")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "shipbob_channel_id", r.shipbobChannelId, "")
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

type ApiUpdateReturnRequest struct {
	ctx context.Context
	ApiService *ReturnsApiService
	shipbobChannelId *int32
	id int32
	createReturn *CreateReturn
}

// Channel Id for Operation
func (r ApiUpdateReturnRequest) ShipbobChannelId(shipbobChannelId int32) ApiUpdateReturnRequest {
	r.shipbobChannelId = &shipbobChannelId
	return r
}

// Model defining the return
func (r ApiUpdateReturnRequest) CreateReturn(createReturn CreateReturn) ApiUpdateReturnRequest {
	r.createReturn = &createReturn
	return r
}

func (r ApiUpdateReturnRequest) Execute() (*ReturnOrder, *http.Response, error) {
	return r.ApiService.UpdateReturnExecute(r)
}

/*
UpdateReturn Modify Return Order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id of the return order
 @return ApiUpdateReturnRequest
*/
func (a *ReturnsApiService) UpdateReturn(ctx context.Context, id int32) ApiUpdateReturnRequest {
	return ApiUpdateReturnRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ReturnOrder
func (a *ReturnsApiService) UpdateReturnExecute(r ApiUpdateReturnRequest) (*ReturnOrder, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnOrder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.UpdateReturn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1.0/return/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shipbobChannelId == nil {
		return localVarReturnValue, nil, reportError("shipbobChannelId is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "shipbob_channel_id", r.shipbobChannelId, "")
	// body params
	localVarPostBody = r.createReturn
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
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
