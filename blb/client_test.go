package blb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/drinktee/bce-sdk-go/bce"
	"github.com/gorilla/mux"
)

var (
	testHTTPServer *httptest.Server
	blbClient      *Client
)

func init() {
	var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

	//var bceConfig = bce.NewConfig(credentials)
	var bceConfig = &bce.Config{
		Credentials: credentials,
		Checksum:    true,
		Region:      os.Getenv("BOS_REGION"),
	}
	var bccConfig = NewConfig(bceConfig)
	blbClient = NewBLBClient(bccConfig)
	r := mux.NewRouter()
	r.HandleFunc("/v1/blb", handleGetBLB).Methods("GET")
	r.HandleFunc("/v1/blb/{blbid}", handleDeleteBLB).Methods("DELETE")
	r.HandleFunc("/v1/blb/{blbid}", handleUpdateBLB).Methods("PUT")
	r.HandleFunc("/v1/blb", handleCreateBLB).Methods("POST")
	testHTTPServer = httptest.NewServer(r)
	blbClient.Endpoint = testHTTPServer.URL
}

func handleGetBLB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := `{
    "blbList":[
        {
            "blbId":"lb-a7e5zPPk",
            "status":"available",
            "name":"test-blb",
            "desc":"用于生产环境",
            "vpcId":"vpc-fy6vdykpwkqb",
            "address":"10.32.249.113"
        },
        {
            "blbId": "lb-gj5gVpeq",
            "status":"available",
            "name": "nametest",
            "desc": "用于测试环境",
            "vpcId":"vpc-a8n5p6kybbx4",
            "address": "10.32.251.4"
        }
    ],
    "marker": "blb-0A20F971",
    "nextMarker": "blb-0A20FB09",
    "isTruncated": true,
    "maxKeys": 2
}`
	fmt.Fprint(w, response)
}

func handleCreateBLB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	args := &CreateLoadBalancerArgs{}
	json.Unmarshal(body, args)
	if args.Name != expectCreateBLB.Name {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response := `{
    "blbId": "lb-BLuOPSLZ",
    "name": "blb-for-test",
    "desc": "",
    "address": "10.32.251.93"
}`
	fmt.Fprint(w, response)
}

func handleUpdateBLB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	blbid := vars["blbid"]
	if blbid == "lb-e5b33752" {
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(400)
}

func handleDeleteBLB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	blbid := vars["blbid"]
	if blbid == "lb-426fad2b" {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}

}
