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
	resp, err := http.Get(testurl)

	if err != nil {
		return err
	}

	// Make sure we close the body so we're not leaving anything hanging. 
	// We don't need to read the body since we're only checking status.
	defer resp.Body.Close()

	// Either we succeeded or we didn't
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("file %s not found, TCP status code: %d", testurl, resp.StatusCode)
	}

	return nil
}

