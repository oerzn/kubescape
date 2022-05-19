package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	utilsmetav1 "github.com/armosec/opa-utils/httpserver/meta/v1"
	"github.com/armosec/utils-go/boolutils"
	"github.com/stretchr/testify/assert"
)

func TestGetScanParamsFromRequest(t *testing.T) {
	{
		body := utilsmetav1.PostScanRequest{
			Submit:      boolutils.BoolPointer(true),
			HostScanner: boolutils.BoolPointer(true),
			Account:     "aaaaaaaaaa",
		}

		jsonBytes, err := json.Marshal(body)
		assert.NoError(t, err)

		u := url.URL{
			Scheme:   "http",
			Host:     "bla",
			Path:     "bla",
			RawQuery: "wait=true&keep=true",
		}
		request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(jsonBytes))
		assert.NoError(t, err)

		scanID := "ccccccc"

		req, err := getScanParamsFromRequest(request, scanID)
		assert.NoError(t, err)
		assert.Equal(t, scanID, req.scanID)
		assert.True(t, req.scanQueryParams.KeepResults)
		assert.True(t, req.scanQueryParams.ReturnResults)
		assert.True(t, *req.scanRequest.HostScanner)
		assert.True(t, *req.scanRequest.Submit)
		assert.Equal(t, "aaaaaaaaaa", req.scanRequest.Account)
	}

	{
		body := utilsmetav1.PostScanRequest{
			Submit:      boolutils.BoolPointer(false),
			HostScanner: boolutils.BoolPointer(false),
			Account:     "aaaaaaaaaa",
		}

		jsonBytes, err := json.Marshal(body)
		assert.NoError(t, err)

		u := url.URL{
			Scheme: "http",
			Host:   "bla",
			Path:   "bla",
		}
		request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(jsonBytes))
		assert.NoError(t, err)

		scanID := "ccccccc"

		req, err := getScanParamsFromRequest(request, scanID)
		assert.NoError(t, err)
		assert.Equal(t, scanID, req.scanID)
		assert.False(t, req.scanQueryParams.KeepResults)
		assert.False(t, req.scanQueryParams.ReturnResults)
		assert.False(t, *req.scanRequest.HostScanner)
		assert.False(t, *req.scanRequest.Submit)
		assert.Equal(t, "aaaaaaaaaa", req.scanRequest.Account)
	}
}
