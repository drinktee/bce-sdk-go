package cce

import (
	"fmt"
	"net/http"
	"path"

	"github.com/drinktee/bce-sdk-go/bce"
)

var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

//var bceConfig = bce.NewConfig(credentials)
var bceConfig = &bce.Config{
	Credentials: credentials,
	Checksum:    true,
}
var cceConfig = NewConfig(bceConfig)
var cceClient = NewClient(cceConfig)

func InstancesHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/instance", func(w http.ResponseWriter, r *http.Request) {
		handleInstanceList(w, r)
	})
	mux.HandleFunc("/v2/instance/", func(w http.ResponseWriter, r *http.Request) {
		handleDescribeInstance(w, r)
	})
	return mux
}

func handleInstanceList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		response := `{
    "instances": [
        {
            "id": "i-IyWRtII7",
            "name": "instance-j93wzbn1",
            "internalIp": "192.168.6.15",
            "zoneName": "cn-bj-a",
            "vpcId": "vpc-9999"
        },
        {
        "id": "i-YufwpQAe",
        "name": "instance-luz2ef4l-1",
        "internalIp": "192.168.0.25",
        "zoneName": "cn-bj-a",
        "vpcId": "vpc-i80sab3o"
    }        
    ]
}`
		fmt.Fprint(w, response)
	}
}

func handleDescribeInstance(w http.ResponseWriter, r *http.Request) {
	_, id := path.Split(r.URL.Path)
	if id != "i-YufwpQAe" {
		return
	}
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		response := `{
    "instance": {
        "id": "i-YufwpQAe",
        "createTime": "2015-07-09T10:27:15Z",
        "name": "instance-luz2ef4l-1",
        "status": "Stopped",  
        "desc": "console",
        "paymentTiming":"Postpaid",
        "expireTime": null,
        "internalIp": "192.168.0.25",
        "publicIp": "-",
        "cpuCount": 1,
        "memoryCapacityInGB": 1,
        "localDiskSizeInGB": 0,
        "networkCapacityInMbps": 5,
        "imageId": "m-nky7qeom",
        "placementPolicy": "default",
        "zoneName": "cn-bj-a",
        "subnetId": "sbn-oioiadda",
        "vpcId": "vpc-i80sab3o"
    }
}`
		fmt.Fprint(w, response)
	}
}
