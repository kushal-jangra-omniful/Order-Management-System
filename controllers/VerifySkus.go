package controllers

import (
	"fmt"
	"net/http"
)

func VerifySKU(sku string) (bool, error) {
	wmsAPIURL := fmt.Sprintf("http://localhost:8082/skus/%s", sku)
	resp, err := http.Get(wmsAPIURL)
	if err != nil {
		return false, fmt.Errorf("failed to reach WMS: %v", err)
	}
	defer resp.Body.Close()

	// If WMS returns 200 OK, SKU exists
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	// If SKU is not found (404 or other errors), return false
	return false, nil
}
