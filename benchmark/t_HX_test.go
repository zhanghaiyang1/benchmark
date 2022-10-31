package main

import (
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

/*
online
*/
var (
	desKey        = "bac03ac72e2346118e18f0362c991870"
	md5Key        = "1d6083e02a0044abac3f4eb3d9267623"
	partnerID     = "HZZ1659349268861"
	bizCode       = "orderDetailStatus"
	api           = "https://cloud.hexiangibc.com/ipolicy-api/API/v1"
	oid           = "7dbb97fc2e2547eda0c05ce8e84"
)

/*
test
*/
// var (
// 	desKey        = "c593b4a4fb444aca822f932be5aa4d8e"
// 	md5Key        = "0468a8ae77314b99b3e737cd0c84da50"
// 	partnerID     = "HZZ1658213564629"
// 	bizCode       = "orderDetailStatus"
// 	mainSignNO    = "MQY20220719162852682jVG5"
// 	subSignNO     = "AQY20220719163013499i68X"
// 	commodityCode = "RCP00003945"
// 	api           = "http://app.xiaozhibaoxian.com/bxpf-api/API/v1"
// 	oid           = "a9dccf96719844ba9aac73bd230"
// )

func TestGetOrder(t *testing.T) {
	orderID := oid
	md5Key := md5Key

	bizCode := bizCode
	//1,加密
	req := &OrderReq{}
	req.Body.Data.OrderID = orderID
	req.Head.PartnerID = partnerID
	req.Head.BizCode = bizCode
	origin, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	str := string(origin)
	Write(&str)
	secret := Encrypt(origin, []byte(desKey[:des.BlockSize]))
	Write(secret)
	//2，请求
	param := "data=" + *secret
	param += "&partner_id=" + partnerID

	mill := time.Now().UnixNano() / 1e6
	param += "&time=" + strconv.Itoa(int(mill))
	signParam := "data=" + *secret + "&time=" + strconv.Itoa(int(mill)) + "&" + md5Key
	sign := Md5(&signParam)
	//data=xxx&time=xxx&md5key
	param += "&sign=" + *sign
	resp := POSTUrlencoded(&api, &param)
	Write(resp)
	//
}
func Md5(param *string) *string {
	h := md5.New()
	h.Write([]byte(*param))
	sign := hex.EncodeToString(h.Sum(nil))
	return &sign
}
func POSTUrlencoded(url, param *string) *string {
	resp, err := http.Post(*url,
		"application/x-www-form-urlencoded",
		strings.NewReader(*param))
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
	res := string(body)
	fmt.Println(res)
	return &res
}

type OrderReq struct {
	Head struct {
		BizCode         string `json:"biz_code"`
		PartnerID       string `json:"partner_id"`
		// CommodityCode   string `json:"commodity_code"`
	} `json:"head"`
	Body struct {
		Msg  string `json:"msg"`
		Code string `json:"code"`
		Data struct {
			OrderID         string `json:"order_id"`
		} `json:"data"`
	} `json:"body"`
}
