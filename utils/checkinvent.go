package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// InventoryResponse defines the expected response structure from WMS
type InventoryResponse struct {
	ID               string `json:"id"`
	HubID           string `json:"hubid"`
	TenantID        string `json:"tenantid"`
	SellerID        string `json:"sellerid"`
	SKU_ID         string `json:"sku_id"`
	AvailableQuantity int  `json:"available_quantity"`
}

func CheckInventory(sellerID, hubID, skuID string, requiredQuantity int) (bool, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	url := fmt.Sprintf("http://localhost:8081/inventories?sellerid=%s&hubid=%s", sellerID, hubID)

	resp, err := client.Get(url)
	if err != nil {
		return false, fmt.Errorf("error contacting WMS: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected response from WMS: %d", resp.StatusCode)
	}

	var inventories []InventoryResponse
	err = json.NewDecoder(resp.Body).Decode(&inventories)
	if err != nil {
		return false, fmt.Errorf("error decoding WMS response: %v", err)
	}

	for _, inventory := range inventories {
		if inventory.SKU_ID == skuID && inventory.AvailableQuantity >= requiredQuantity {
			return true, nil
		}
	}

	return false, nil
}
