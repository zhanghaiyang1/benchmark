package main

import (
	"crypto/des"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"testing"
)

var desKey string = "c593b4a4fb444aca822f932be5aa4d8e" 
func TestDecrypt(t *testing.T) {
	var data = "bTVoRVhwUEZ0NEtqM01QYjN6UTc2UXZBa2FNbysxT3YxNHdKNlEwRFpSYWZPNVdUdzNoSEtIU3NvN1pkajJNb01QWmRtTzBZUkljVnVXcTR5azlmUXc9PQ=="
	data1, _ := base64.StdEncoding.DecodeString(data)
	origin := Decrypt(data1, []byte(desKey[:des.BlockSize]))
	originStr := string(origin)
	fmt.Println("originstr:", originStr)
	Write(&originStr)
	//加密
	secret := Encrypt([]byte(origin), []byte(desKey[:des.BlockSize]))
	Write(secret)

}
func TestEncrypt(t *testing.T) {
	//解密 secret->decode->decode->des.decrypt->origin
	//加密 origin->des.encrypt->encode->encode->secret
	origin := `{"head":{"firm_code":"","sub_signature_no":"AQY20220719163013499i68X","app_type":"APP_TYPE:H5","emp_num":"","biz_code":"payNotice","subjectStyle":"","main_part_code":"B202103190022","application_id":"YY20220719163549998OBoR","userCode":"HZZ1658213564629","partner_id":"HZZ1658213564629","amountStyle":"","remark9":"","front_env":"","remark8":"","remark5":"","remark10":"","remark4":"","remark7":"","remark6":"","remark1":"4","main_signature_no":"MQY20220719162852682jVG5","remark3":"","remark2":""},"body":{"msg":"TRADE_SUCCESS","code":"0000","data":{"pay_status":"ODR_PAY_STATUS:2","pay_chanel":"alipay","pay_serial_number":"2022080322001400171418376855","order_group_id":"a9dccf96719844ba9aac73bd230","pay_amount":"0.01","pay_card_no":"2088612964400174","order_id":"a9dccf96719844ba9aac73bd230","pay_time":"20220803114836"}}}`
	secret := Encrypt([]byte(origin), []byte(desKey[:des.BlockSize]))	
	Write(secret)
}
func Encrypt(data, key []byte) *string{
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	bs := block.BlockSize()
	length := len(data)
	out := make([]byte, length)
	fmt.Println("len+:", length + (8 - length % 8 ))
	if length % 8 > 0 {

		out = make([]byte, length + (8 - length % 8 ))	
		add := 8 - length % 8 
		for add > 0 {
			data = append(data, 0)
			add--
		}
	}
    dst := out
	// mark := false
    for len(data) > 0 {

		block.Encrypt(dst, data[:bs])
        data = data[bs:]
        dst = dst[bs:]

    }
	
	baseOne := base64.StdEncoding.EncodeToString(out)
	baseTwo := base64.StdEncoding.EncodeToString([]byte(baseOne))

	return &baseTwo
}

func Decrypt(str, key []byte) []byte{
	src, err := base64.StdEncoding.DecodeString(string(str))
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}

	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	bs := block.BlockSize()

	out := make([]byte, len(src))
    dst := out
    if len(src) % bs != 0 {
        fmt.Println("err------->")
    }
    for len(src) > 0 {
        block.Decrypt(dst, src[:bs])
        src = src[bs:]
        dst = dst[bs:]
    }
    // out = PKCS5UnPadding(out)

	fmt.Println("out:", out)
	return out
}

func Write(text *string){


		var filename = "./log.txt"
		var f *os.File

		f, _ = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm) //打开文件
		fmt.Println("文件存在")

		defer f.Close()
		n, err1 := io.WriteString(f, *text) 
		if err1 != nil {
		   panic(err1)
		}
		fmt.Printf("写入 %d 个字节n", n)
}