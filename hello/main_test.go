package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	mockUserResp := `{"msg":"hello world"}`
	ts := httptest.NewServer(SetUpSever())
	defer ts.Close()
	resp, err := http.Get(fmt.Sprintf("%s/", ts.URL))
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	responseData, _ := io.ReadAll(resp.Body)
	assert.Equal(t, mockUserResp, string(responseData))
}
