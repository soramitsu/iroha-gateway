// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "Iroha-Gateway Server": signatories TestHelpers
//
// Command:
// $ goagen
// --design=github.com/soramitsu/iroha-gateway/design
// --out=$(GOPATH)/src/github.com/soramitsu/iroha-gateway
// --version=v1.2.0-dirty

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/soramitsu/iroha-gateway/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// AddSignatoriesBadRequest runs the method Add of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddSignatoriesBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SignatoriesController, uuid string, payload *app.SignatoryRequest) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/accounts/%v/signatories", uuid),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["uuid"] = []string{fmt.Sprintf("%v", uuid)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SignatoriesTest"), rw, req, prms)
	addCtx, __err := app.NewAddSignatoriesContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	addCtx.Payload = payload

	// Perform action
	__err = ctrl.Add(addCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Message)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// AddSignatoriesInternalServerError runs the method Add of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddSignatoriesInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SignatoriesController, uuid string, payload *app.SignatoryRequest) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/accounts/%v/signatories", uuid),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["uuid"] = []string{fmt.Sprintf("%v", uuid)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SignatoriesTest"), rw, req, prms)
	addCtx, __err := app.NewAddSignatoriesContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	addCtx.Payload = payload

	// Perform action
	__err = ctrl.Add(addCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Message)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// AddSignatoriesOK runs the method Add of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddSignatoriesOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SignatoriesController, uuid string, payload *app.SignatoryRequest) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/accounts/%v/signatories", uuid),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["uuid"] = []string{fmt.Sprintf("%v", uuid)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SignatoriesTest"), rw, req, prms)
	addCtx, __err := app.NewAddSignatoriesContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	addCtx.Payload = payload

	// Perform action
	__err = ctrl.Add(addCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Message)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// DeleteSignatoriesBadRequest runs the method Delete of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteSignatoriesBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SignatoriesController, uuid string, sig string, payload *app.DeleteSignatoryRequest) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/accounts/%v/signatories/%v", uuid, sig),
	}
	req, _err := http.NewRequest("DELETE", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["uuid"] = []string{fmt.Sprintf("%v", uuid)}
	prms["sig"] = []string{fmt.Sprintf("%v", sig)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SignatoriesTest"), rw, req, prms)
	deleteCtx, __err := app.NewDeleteSignatoriesContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	deleteCtx.Payload = payload

	// Perform action
	__err = ctrl.Delete(deleteCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Message)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// DeleteSignatoriesInternalServerError runs the method Delete of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteSignatoriesInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SignatoriesController, uuid string, sig string, payload *app.DeleteSignatoryRequest) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/accounts/%v/signatories/%v", uuid, sig),
	}
	req, _err := http.NewRequest("DELETE", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["uuid"] = []string{fmt.Sprintf("%v", uuid)}
	prms["sig"] = []string{fmt.Sprintf("%v", sig)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SignatoriesTest"), rw, req, prms)
	deleteCtx, __err := app.NewDeleteSignatoriesContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	deleteCtx.Payload = payload

	// Perform action
	__err = ctrl.Delete(deleteCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Message)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// DeleteSignatoriesOK runs the method Delete of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteSignatoriesOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SignatoriesController, uuid string, sig string, payload *app.DeleteSignatoryRequest) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/accounts/%v/signatories/%v", uuid, sig),
	}
	req, _err := http.NewRequest("DELETE", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["uuid"] = []string{fmt.Sprintf("%v", uuid)}
	prms["sig"] = []string{fmt.Sprintf("%v", sig)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SignatoriesTest"), rw, req, prms)
	deleteCtx, __err := app.NewDeleteSignatoriesContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	deleteCtx.Payload = payload

	// Perform action
	__err = ctrl.Delete(deleteCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Message)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// GetAllSignatoriesBadRequest runs the method GetAll of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetAllSignatoriesBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SignatoriesController, uuid string, creatorPubkey *string, isCommitted *bool) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if creatorPubkey != nil {
		sliceVal := []string{*creatorPubkey}
		query["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		query["is_committed"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/accounts/%v/signatories", uuid),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["uuid"] = []string{fmt.Sprintf("%v", uuid)}
	if creatorPubkey != nil {
		sliceVal := []string{*creatorPubkey}
		prms["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		prms["is_committed"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SignatoriesTest"), rw, req, prms)
	getAllCtx, _err := app.NewGetAllSignatoriesContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetAll(getAllCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Message)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetAllSignatoriesInternalServerError runs the method GetAll of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetAllSignatoriesInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SignatoriesController, uuid string, creatorPubkey *string, isCommitted *bool) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if creatorPubkey != nil {
		sliceVal := []string{*creatorPubkey}
		query["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		query["is_committed"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/accounts/%v/signatories", uuid),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["uuid"] = []string{fmt.Sprintf("%v", uuid)}
	if creatorPubkey != nil {
		sliceVal := []string{*creatorPubkey}
		prms["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		prms["is_committed"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SignatoriesTest"), rw, req, prms)
	getAllCtx, _err := app.NewGetAllSignatoriesContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetAll(getAllCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Message)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetAllSignatoriesOK runs the method GetAll of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetAllSignatoriesOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SignatoriesController, uuid string, creatorPubkey *string, isCommitted *bool) (http.ResponseWriter, *app.Signatories) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if creatorPubkey != nil {
		sliceVal := []string{*creatorPubkey}
		query["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		query["is_committed"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/accounts/%v/signatories", uuid),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["uuid"] = []string{fmt.Sprintf("%v", uuid)}
	if creatorPubkey != nil {
		sliceVal := []string{*creatorPubkey}
		prms["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		prms["is_committed"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SignatoriesTest"), rw, req, prms)
	getAllCtx, _err := app.NewGetAllSignatoriesContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetAll(getAllCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Signatories
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Signatories)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Signatories", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
