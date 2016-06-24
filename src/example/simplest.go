package main

import (
	"sdk"
	"fmt"
)

func main() {
	var config *sdk.Config = sdk.NewConfig()
	var credential *sdk.Credential = sdk.NewCredential()
	var client *sdk.Client = sdk.NewClient()

	//init config
	config.AppVersionValue = "1.0.1"
	config.AppMarketIdValue = "678"
	config.DisableSSL = true
	config.GatewayHost = "api.OpsKitchen.com"

	//init credential
	credential.AppKey = "101"
	credential.Secret = "your secret"

	//init client
	client.RequestBuilder.Config = config
	client.RequestBuilder.Credential = credential

	//init user
	resp, err := client.CallApi("ops.meta.os.list", "1.0", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
}
