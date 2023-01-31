package main

import (

	"encoding/json"
	"fmt"

	"testing"
)

/*
online
*/

/*
test
*/
var (
	apiRZ = "https://sandbox-gw.insuremo.com.cn"
	secret = "maMbrpeVWtRbjXFlxSa7ka4qGcI6XqJP"
)


func TestGetLink(t *testing.T) {
	/*

		1.将HTTP Body的内容和secret key进行字符串拼接；
		2.对于拼接后的内容使用MD5算法进行签名运算。
	*/
	linkReq := LinkReq{}
	linkReq.QuotationNumber = "Q123210321"
	linkReq.ProductCode = "P123"
	linkReq.ChannelCode = "C001"
	linkReq.AgreementCode = "A001"
	linkReq.AgentCode = "A001"
	linkReq.InsuranceCompanyCode = "SRSH"
	linkReq.BusinessType = "Pro"
	linkReq.BusinessFrom = "Agent"
	linkReq.ExternalInfo = "{}"
	linkReq.Factors = "{}"

	//获取sign
	bodyByte, _ := json.Marshal(linkReq)
	fmt.Println("bodyByte:", string(bodyByte))

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