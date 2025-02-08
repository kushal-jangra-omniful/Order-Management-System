package interservice

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/omniful/go_commons/http"
	interservice_client "github.com/omniful/go_commons/interservice-client"
)

type SKUResponse struct {
	IsSuccess bool        `json:"is_success"`
	StatusCode int        `json:"status_code"`
	Data      struct {
		ID             string         `json:"id"`
		ProductID      string         `json:"product_id"`
		Price          float64        `json:"price"`
		Fragile        string         `json:"fragile"`
		Specifications string         `json:"specifications"`
		CreatedAt      time.Time      `json:"created_at"`
		UpdatedAt      time.Time      `json:"updated_at"`
		DeletedAt      *time.Time     `json:"deleted_at"`
	} `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

var Client *interservice_client.Client



func GetReq(ctx context.Context, userData interface{}, Url string) (interface{}, *interservice_client.Error) {
	request := &http.Request{
		Url: Url,
	}

	_, err := Client.Get(request, &userData)
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(userData)
	fmt.Println(string(jsonData))

	return &userData, nil
}
func InitInterSrvClient() {
	config := interservice_client.Config{
		ServiceName: "user-service",
		BaseURL:     "http://localhost:8082",
		Timeout:     5 * time.Second,
	}

	client, err := interservice_client.NewClientWithConfig(config)
	if err != nil {
		panic(err)
	}
	Client = client
	// var data []Hub
	// GetReq(context.Background(), &data, "/hub/view")
}

func PostReq(ctx context.Context, userData interface{}, Url string, body interface{}) (interface{}, *interservice_client.Error) {
	request := &http.Request{
		Url:  Url,
		Body: body,
	}

	_, err := Client.Post(request, &userData)
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(userData)
	fmt.Println(string(jsonData))

	return &userData, nil
}


func VerifySKU(ctx context.Context, sku string) (bool, error) {
	url := fmt.Sprintf("http://localhost:8082/skus/%s", sku)

	var response SKUResponse
	_, err := GetReq(ctx, &response, url)
	if err != nil {
		return false, fmt.Errorf("failed to verify SKU: %v", err)
	}

	return true, nil
}