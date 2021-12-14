package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

func makeRequest(ctx context.Context, request, response interface{}, method string, endpoint string) error {
	atomic.AddInt64(&totalRequests, 1)
	var reqBodyBuffer io.Reader
	if request != nil {
		reqBodyBytes, err := json.Marshal(request)
		if err != nil {
			return fmt.Errorf("json.Marshal: %w", err)
		}
		reqBodyBuffer = bytes.NewBuffer(reqBodyBytes)
	}
	req, err := http.NewRequestWithContext(ctx, method, baseUrl+endpoint, reqBodyBuffer)
	if err != nil {
		return fmt.Errorf("http.NewRequest: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http.DefaultClient.Do: %w", err)
	}
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("got status code %d. Body: %s", resp.StatusCode, string(respBody))
	}
	if response != nil {
		err = json.NewDecoder(resp.Body).Decode(response)
		if err != nil {
			return fmt.Errorf("json.DewDecoder().Decode: %w", err)
		}
	}
	return nil
}
