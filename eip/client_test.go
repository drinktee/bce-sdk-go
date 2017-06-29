package eip

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"io/ioutil"

	"time"

	"github.com/drinktee/bce-sdk-go/bce"
)

var eipClient *Client

func init() {
	var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

	//var bceConfig = bce.NewConfig(credentials)
	var bceConfig = &bce.Config{
		Credentials: credentials,
		Checksum:    true,
		Timeout:     5 * time.Second,
		Region:      os.Getenv("BOS_REGION"),
	}
	var bccConfig = NewConfig(bceConfig)
	eipClient = NewEIPClient(bccConfig)
	eipClient.SetDebug(true)
}

func EipHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/eip", func(w http.ResponseWriter, r *http.Request) {
		handleCreateGetEip(w, r)
	})
	mux.HandleFunc("/v1/eip/", func(w http.ResponseWriter, r *http.Request) {
		handleUpdateEip(w, r)
	})
	return mux
}

func handleCreateGetEip(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		response := `{
    "eipList": [
        {
            "name":"eip-xrllt5M-1",
            "eip": "180.181.3.133",
            "status":"binded",
            "instanceType": "BCC",
            "instanceId": "i-IyWRtII7",
            "shareGroupId": "eg-0c31c93a",
            "eipInstanceType": "shared",
            "bandwidthInMbps": 5,
            "paymentTiming":"Prepaid",
            "billingMethod":null,
            "createTime":"2016-03-08T08:13:09Z",
            "expireTime":"2016-04-08T08:13:09Z"
        },
        {
            "name":"eip-scewa1M-1",
            "eip": "180.181.3.134",
            "status":"binded",
            "instanceType": "BCC",
            "instanceId": "i-KjdgweC4",
            "shareGroupId": null,
            "eipInstanceType": "normal",
            "bandwidthInMbps": 1,
            "paymentTiming":"Postpaid",
            "billingMethod":"ByTraffic",
            "createTime":"2016-03-08T08:13:09Z",
            "expireTime":null
        }
    ],
    "marker":"eip-DCB50385",
    "isTruncated": true,
    "nextMarker": "eip-DCB50387",
    "maxKeys": 2
}`
		fmt.Fprint(w, response)
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		args := &CreateEipArgs{}
		json.Unmarshal(body, args)
		if args.Billing.BillingMethod != "ByTraffic" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		response := `{
    "eip":"180.181.3.133"
}`
		fmt.Fprint(w, response)
	}
}

func handleUpdateEip(w http.ResponseWriter, r *http.Request) {
	_, eip := path.Split(r.URL.Path)
	if r.Method == http.MethodPut {
		query := r.URL.Query()
		_, ok := query["resize"]
		if ok {
			if eip != expectResizeEip.Ip {
				w.WriteHeader(400)
			}
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			args := &ResizeEipArgs{}
			json.Unmarshal(body, args)
			if args.BandwidthInMbps != expectResizeEip.BandwidthInMbps {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
		_, ok = query["bind"]
		if ok {
			if eip != expectBindEip.Ip {
				w.WriteHeader(400)
			}
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			args := &BindEipArgs{}
			json.Unmarshal(body, args)
			if args.InstanceType != expectBindEip.InstanceType {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}

		_, ok = query["unbind"]
		if ok {
			if eip != expectBindEip.Ip {
				w.WriteHeader(400)
			}
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
	if r.Method == http.MethodDelete {
		if eip != expectUnbindEip.Ip {
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(400)
}
