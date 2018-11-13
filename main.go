package main

import (
	"encoding/json"
	"fmt"

	"k8s.io/client-go/pkg/apis/clientauthentication/v1beta1"
)

func main() {
	creds := v1beta1.ExecCredential{}
	jCreds, _ := json.Marshal(creds)
	fmt.Println(string(jCreds))
}
