package main

import (
	"crypto/des"
	"encoding/json"
	"fmt"
	"testing"
)
func TestGetOrder(t *testing.T){
	//1,加密
	req := &OrderReq{}
	req.Body.Data.OrderID = "a9dccf96719844ba9aac73bd230"
	origin, err := json.Marshal(req)
    if err != nil {
        fmt.Printf("Error: %s", err)
    }
	str := string(origin)
	Write(&str)
    secret := Encrypt(origin, []byte(desKey[:des.BlockSize]))
	Write(secret)
	//2，请求

	//
}

type OrderReq struct {
	Head struct {
		FirmCode        string `json:"firm_code"`
		SubSignatureNo  string `json:"sub_signature_no"`
		AppType         string `json:"app_type"`
		EmpNum          string `json:"emp_num"`
		BizCode         string `json:"biz_code"`
		SubjectStyle    string `json:"subjectStyle"`
		MainPartCode    string `json:"main_part_code"`
		ApplicationID   string `json:"application_id"`
		UserCode        string `json:"userCode"`
		PartnerID       string `json:"partner_id"`
		AmountStyle     string `json:"amountStyle"`
		Remark9         string `json:"remark9"`
		FrontEnv        string `json:"front_env"`
		Remark8         string `json:"remark8"`
		Remark5         string `json:"remark5"`
		Remark10        string `json:"remark10"`
		Remark4         string `json:"remark4"`
		Remark7         string `json:"remark7"`
		Remark6         string `json:"remark6"`
		Remark1         string `json:"remark1"`
		MainSignatureNo string `json:"main_signature_no"`
		Remark3         string `json:"remark3"`
		Remark2         string `json:"remark2"`
	} `json:"head"`
	Body struct {
		Msg  string `json:"msg"`
		Code string `json:"code"`
		Data struct {
			PayStatus       string `json:"pay_status"`
			PayChanel       string `json:"pay_chanel"`
			PaySerialNumber string `json:"pay_serial_number"`
			OrderGroupID    string `json:"order_group_id"`
			PayAmount       string `json:"pay_amount"`
			PayCardNo       string `json:"pay_card_no"`
			OrderID         string `json:"order_id"`
			PayTime         string `json:"pay_time"`
		} `json:"data"`
	} `json:"body"`
}