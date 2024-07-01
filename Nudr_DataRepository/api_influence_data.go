/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudr_DataRepository

import (
	"github.com/Noodlesalat/openapi"
	"github.com/Noodlesalat/openapi/models"

	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type InfluenceDataApiService service

/*
InfluenceDataApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ApplicationDataInfluenceDataGetParamOpts - Optional Parameters:
 * @param "InfluenceIds" (optional.Interface of []string) -  Each element identifies a service.
 * @param "Dnns" (optional.Interface of []string) -  Each element identifies a DNN.
 * @param "Snssais" (optional.Interface of []models.Snssai) -  Each element identifies a slice.
 * @param "InternalGroupIds" (optional.Interface of []string) -  Each element identifies a group of users.
 * @param "Supis" (optional.Interface of []string) -  Each element identifies the user.
@return []models.TrafficInfluData
*/

type ApplicationDataInfluenceDataGetParamOpts struct {
	InfluenceIds     optional.Interface
	Dnns             optional.Interface
	Snssais          optional.Interface
	InternalGroupIds optional.Interface
	Supis            optional.Interface
}

func (a *InfluenceDataApiService) ApplicationDataInfluenceDataGet(ctx context.Context, localVarOptionals *ApplicationDataInfluenceDataGetParamOpts) ([]models.TrafficInfluData, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []models.TrafficInfluData
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/application-data/influenceData"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.InfluenceIds.IsSet() && localVarOptionals.InfluenceIds.Value() != "" {
		localVarQueryParams.Add("influence-Ids", openapi.ParameterToString(localVarOptionals.InfluenceIds.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Dnns.IsSet() && localVarOptionals.Dnns.Value() != "" {
		localVarQueryParams.Add("dnns", openapi.ParameterToString(localVarOptionals.Dnns.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Snssais.IsSet() && localVarOptionals.Snssais.Value() != "" {
		localVarQueryParams.Add("snssais", openapi.ParameterToString(localVarOptionals.Snssais.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.InternalGroupIds.IsSet() && localVarOptionals.InternalGroupIds.Value() != "" {
		localVarQueryParams.Add("internal-Group-Ids", openapi.ParameterToString(localVarOptionals.InternalGroupIds.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Supis.IsSet() && localVarOptionals.Supis.Value() != "" {
		localVarQueryParams.Add("supis", openapi.ParameterToString(localVarOptionals.Supis.Value(), "csv"))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	case 400:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 401:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 403:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 404:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 406:
		return localVarReturnValue, localVarHTTPResponse, nil
	case 414:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 429:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	default:
		return localVarReturnValue, localVarHTTPResponse, nil
	}
}
