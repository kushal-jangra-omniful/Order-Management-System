package utils

import (
	"github.com/omniful/go_commons/sqs"

)

func init(){
	config:=sqs.Config{
		Account: "539247490249",
		Endpoint: "https://sqs.eu-north-1.amazonaws.com/539247490249/Myqueue.fifo",
		Region: "eu-north-1",
	}
	
}

