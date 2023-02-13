package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"testing"
)

/*
online
*/

/*
test
*/
var (
	apiRZ  = "https://sandbox-gw.insuremo.com.cn"
	secret = "maMbrpeVWtRbjXFlxSa7ka4qGcI6XqJP"
	token = "eBaoCCwzDk-_0T4BPyR9Og0IR4AjIesA"
)

func TestGetLink(t *testing.T) {

	linkReq := LinkReq{}
	linkReq.QuotationNumber = "SMECLE23002001"
	linkReq.ProductCode = "EL_JianFei"
	linkReq.ChannelCode = "0000159"
	linkReq.AgreementCode = "SH2023187B"
	linkReq.AgentCode = "bbyfu.app"
	linkReq.InsuranceCompanyCode = "SRSH"
	linkReq.BusinessType = "Quo"
	linkReq.BusinessFrom = "Agent"
	linkReq.ExternalInfo = "{}"
	linkReq.Factors = "{}"

	//获取sign
	bodyByte, _ := json.Marshal(linkReq)
	md5Param := string(bodyByte) + secret
	sign := Md5(&md5Param)

	addr := apiRZ + "/swissre/1.0.0/pa/api/aquila/general/quotation/url?sign=" + *sign
	//请求
	request, err := http.NewRequest("POST", addr, bytes.NewReader(bodyByte))
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("res:", string(body))

}
func TestOrderDetail(t *testing.T) {

	detailReq := OrderDetailReq{}
	detailReq.QuotationNumber = "SMECLE23003802"

	//获取sign
	bodyByte, _ := json.Marshal(detailReq)
	md5Param := string(bodyByte) + secret
	sign := Md5(&md5Param)

	addr := apiRZ + "/swissre/1.0.0/pa/api/aquila/general/quotation/info?sign=" + *sign
	//请求
	request, err := http.NewRequest("POST", addr, bytes.NewReader(bodyByte))
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("res:", string(body))

}
func TestChannelAuth(t *testing.T) {

	authReq := AuthReq{}
	authReq.ProductID = "KoNPRtXdgL:1"
	authReq.ExternalInfo = "A12345678"

	//获取sign
	bodyByte, _ := json.Marshal(authReq)
	md5Param := string(bodyByte) + secret
	sign := Md5(&md5Param)
	fmt.Println("sign:", *sign)
	addr :=  "http://localhost:11818/ruizai/channel/authenticate?sign=" + *sign
	//请求
	request, err := http.NewRequest("POST", addr, bytes.NewReader(bodyByte))
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("res:", string(body))

}
type LinkReq struct {
	QuotationNumber      string `json:"quotationNumber"`
	ProductCode          string `json:"productCode"`
	ChannelCode          string `json:"channelCode"`
	AgreementCode        string `json:"agreementCode"`
	AgentCode            string `json:"agentCode"`
	InsuranceCompanyCode string `json:"insuranceCompanyCode"`
	BusinessType         string `json:"businessType"`
	BusinessFrom         string `json:"businessFrom"`
	ExternalInfo         string `json:"externalInfo"`
	Factors              string `json:"factors"`
}
type OrderDetailReq struct{
	QuotationNumber      string `json:"quotationNumber"`	
}
	
type AuthReq struct {
	ProductID    string `json:"productId"`
	ExternalInfo string `json:"externalInfo"`
}