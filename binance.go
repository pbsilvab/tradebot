package tradebot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const baseurl string = "https://api.binance.com"
const secret string = "dlC5QpjX3XsGLnMYXtGEai3qwUWVbTlevwb0p2mF57YpWZMIW6YuUISPgk9DBe05"
const apikey string = "BjasgLPc4dV2fLevjr0XbRP6LPZ7aYk31wTY00PBezR8mheDffCvHQVyFARFNqMn"

type requestStruct struct {
	endpoint string
	method   string
	params   map[string][]string
}

func prepareRequest(endpoint string, method string, params map[string][]string) requestStruct {

	newReq := requestStruct{
		endpoint: endpoint,
		method:   method,
		params:   params,
	}

	return newReq
}

func currentTimestamps() string {
	now := time.Now()
	ts := fmt.Sprintf("%v", now.UnixMilli())
	return ts
}

func createSignature(qstring string) string {
	secret := secret
	data := qstring
	fmt.Println(qstring)
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	h.Write([]byte(data))
	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

func CallApi(endpoint string, method string, p map[string][]string) string {

	r := prepareRequest(endpoint, method, p)

	timestamp := currentTimestamps()
	rawqs := url.Values(r.params)
	rawqs.Add("timestamp", timestamp)
	rawqs.Add("recvWindow", "5000")

	signature := createSignature(rawqs.Encode())
	rawqs.Add("signature", signature)

	params := strings.NewReader(rawqs.Encode())
	finalUrl := fmt.Sprintf("%v%v?%v", baseurl, r.endpoint, rawqs.Encode())

	client := &http.Client{}
	req, errr := http.NewRequest(r.method, finalUrl, nil)

	if r.method != "GET" && len(r.params) > 0 {
		req, errr = http.NewRequest(r.method, finalUrl, params)
	}

	if errr != nil {
		log.Fatal(errr)
	}

	req.Header.Add("X-MBX-APIKEY", apikey)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	b := string(body)

	return b
}
