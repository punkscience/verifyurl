package verifyurl

import (
	"fmt"
	"net/http"
	"net/url"
)

func Verify(testurl string) error {
	// First, ensure it's formed properly
	_, err := url.ParseRequestURI(testurl)

	if err != nil {
		return err
	}

	// Try to get it and ensure it comes back with 200
	resp, _ := http.Get(testurl)

	if resp.StatusCode != 200 {
		return fmt.Errorf("file %s not found, TCP status code: %d", testurl, resp.StatusCode)
	}

	return nil
}
