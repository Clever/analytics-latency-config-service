package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
	"github.com/go-errors/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/xerrors"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

var _ = strconv.ParseInt
var _ = strfmt.Default
var _ = swag.ConvertInt32
var _ = errors.New
var _ = mux.Vars
var _ = bytes.Compare
var _ = ioutil.ReadAll
var _ = log.String

var formats = strfmt.Default
var _ = formats

// convertBase64 takes in a string and returns a strfmt.Base64 if the input
// is valid base64 and an error otherwise.
func convertBase64(input string) (strfmt.Base64, error) {
	temp, err := formats.Parse("byte", input)
	if err != nil {
		return strfmt.Base64{}, err
	}
	return *temp.(*strfmt.Base64), nil
}

// convertDateTime takes in a string and returns a strfmt.DateTime if the input
// is a valid DateTime and an error otherwise.
func convertDateTime(input string) (strfmt.DateTime, error) {
	temp, err := formats.Parse("date-time", input)
	if err != nil {
		return strfmt.DateTime{}, err
	}
	return *temp.(*strfmt.DateTime), nil
}

// convertDate takes in a string and returns a strfmt.Date if the input
// is a valid Date and an error otherwise.
func convertDate(input string) (strfmt.Date, error) {
	temp, err := formats.Parse("date", input)
	if err != nil {
		return strfmt.Date{}, err
	}
	return *temp.(*strfmt.Date), nil
}

func jsonMarshalNoError(i interface{}) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		// This should never happen
		return ""
	}
	return string(bytes)
}

// statusCodeForHealthCheck returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForHealthCheck(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	default:
		return -1
	}
}

func (h handler) HealthCheckHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	err := h.HealthCheck(ctx)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		} else if xerr, ok := err.(xerrors.Formatter); ok {
			logger.FromContext(ctx).AddContext("frames", fmt.Sprintf("%+v", xerr))
		}
		statusCode := statusCodeForHealthCheck(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(""))

}

// newHealthCheckInput takes in an http.Request an returns the input struct.
func newHealthCheckInput(r *http.Request) (*models.HealthCheckInput, error) {
	var input models.HealthCheckInput

	sp := opentracing.SpanFromContext(r.Context())
	_ = sp

	var err error
	_ = err

	return &input, nil
}

// statusCodeForGetTableLatency returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForGetTableLatency(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.GetTableLatencyResponse:
		return 200

	case *models.InternalError:
		return 500

	case *models.NotFound:
		return 404

	case models.BadRequest:
		return 400

	case models.GetTableLatencyResponse:
		return 200

	case models.InternalError:
		return 500

	case models.NotFound:
		return 404

	default:
		return -1
	}
}

func (h handler) GetTableLatencyHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	sp := opentracing.SpanFromContext(ctx)

	input, err := newGetTableLatencyInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	if input != nil {
		err = input.Validate(nil)
	}

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.GetTableLatency(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		} else if xerr, ok := err.(xerrors.Formatter); ok {
			logger.FromContext(ctx).AddContext("frames", fmt.Sprintf("%+v", xerr))
		}
		statusCode := statusCodeForGetTableLatency(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	jsonSpan, _ := opentracing.StartSpanFromContext(ctx, "json-response-marshaling")
	defer jsonSpan.Finish()

	respBytes, err := json.Marshal(resp)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	sp.LogFields(log.Int("response-size-bytes", len(respBytes)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForGetTableLatency(resp))
	w.Write(respBytes)

}

// newGetTableLatencyInput takes in an http.Request an returns the input struct.
func newGetTableLatencyInput(r *http.Request) (*models.GetTableLatencyRequest, error) {
	sp := opentracing.SpanFromContext(r.Context())
	_ = sp

	var err error
	_ = err

	data, err := ioutil.ReadAll(r.Body)
	if len(data) == 0 {
		return nil, errors.New("request body is required, but was empty")
	}
	sp.LogFields(log.Int("request-size-bytes", len(data)))

	if len(data) > 0 {
		jsonSpan, _ := opentracing.StartSpanFromContext(r.Context(), "json-request-marshaling")
		defer jsonSpan.Finish()

		var input models.GetTableLatencyRequest
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(&input); err != nil {
			return nil, err
		}
		return &input, nil

	}

	return nil, nil
}

// statusCodeForGetAllLegacyConfigs returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForGetAllLegacyConfigs(obj interface{}) int {

	switch obj.(type) {

	case *models.AnalyticsLatencyConfigs:
		return 200

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case models.AnalyticsLatencyConfigs:
		return 200

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	default:
		return -1
	}
}

func (h handler) GetAllLegacyConfigsHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	sp := opentracing.SpanFromContext(ctx)

	resp, err := h.GetAllLegacyConfigs(ctx)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		} else if xerr, ok := err.(xerrors.Formatter); ok {
			logger.FromContext(ctx).AddContext("frames", fmt.Sprintf("%+v", xerr))
		}
		statusCode := statusCodeForGetAllLegacyConfigs(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	jsonSpan, _ := opentracing.StartSpanFromContext(ctx, "json-response-marshaling")
	defer jsonSpan.Finish()

	respBytes, err := json.Marshal(resp)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	sp.LogFields(log.Int("response-size-bytes", len(respBytes)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForGetAllLegacyConfigs(resp))
	w.Write(respBytes)

}

// newGetAllLegacyConfigsInput takes in an http.Request an returns the input struct.
func newGetAllLegacyConfigsInput(r *http.Request) (*models.GetAllLegacyConfigsInput, error) {
	var input models.GetAllLegacyConfigsInput

	sp := opentracing.SpanFromContext(r.Context())
	_ = sp

	var err error
	_ = err

	return &input, nil
}
