package main

import (
	"fmt"
	"ic-indexer-service/app/initializer"
	"ic-indexer-service/client"
	"ic-indexer-service/client/operations"
)

//swagger generate spec -b ic-indexer-service -o ./swagger.json --scan-models
//swagger generate client -f swagger.json
//use to generate client
//install swagger-go from https://goswagger.io/install.html

func main() {

	initializer.Initialize()
	param := operations.NewGetIcecreamParams()
	name := "Vanilla"
	param.Name = &name

	client := client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("localhost:5010"))
	client.Operations.SetTransport(client.Transport)

	icecream, err := client.Operations.GetIcecream(param)
	if err == nil {
		fmt.Print(icecream.Payload.Data)
	} else {
		data, ok := err.(*operations.SaveOrUpdateIcecreamDefault)
		if ok {
			fmt.Print(data.Payload.Data)
		}
	}
}
