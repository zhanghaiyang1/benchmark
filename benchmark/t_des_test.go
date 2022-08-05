package main

import (
	"bufio"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)
var desKey string = "c593b4a4fb444aca822f932be5aa4d8e" 
func TestDecrypt(t *testing.T) {
	var data = "STU1bW9SYVZZK28vQ21DV2lZR1lod3IycDJiYVBwczRjTnBMbldLYXlBamhvQ2RwQ0IyZmF0M2wwMVVaUWtibDFMTVlybXdYR3k5R0luNW9rR0s4NXdRNkJOdUtpM1FnUzRjY1NPN3d4UG0yQTVKTjh2OWQ5aGdiTHY2S1RqQnNITkx0cGtVcFZoWGRWMUJyZWw5WmJZTlpDWC9IU29wY0hDNFB3cDlXbFp1NHYzNDJFZld6VXNZNXBFSS9yOFBtNy9wQmsrc2srMVhBOWpVWGxRMjkzVlBRT0hVSnA1WmcwR3dmNWh4bjhVd3I3U0FMVTMvRW1sQ1RGQ3JGZnlINlVPSlk0NG9LZWk0NENhcitPaUdKNTVST1VXUVFGR1pVelZueVh0TGV1bFA5b3UyQU1kVlJnSDcveXBnYXZFd0x3U3RienFEVnc4TmM1WjZzYXNkOStkTmtvM09iMmxLRTMyUWZKMUtEazBQQksxdk9vTlhEdzF6bG5xeHF4MzM1Y3ViL1RBQjFmRThPWXhaVXQrajBmZVg4OHBSMDIvazUzdUxJSnNEZTM3STlCNllCa0NtMXBMU0k1anRkK0hyR3I3NnY2Y2pLNmU4Nzhab0JXcEpnR2tMV2RoYmxCTHFtbEF0T0lvanAxUVcwaU9ZN1hmaDZ4bDZOcmRibVFWaENPL0dhQVZxU1lCcUZIMHVzZVpjOCt0VGtmR0JhS1BYWlVSN0VoVXRYMW54ZHdGbCtNVGdWQVlwNmd6NllGMEtoVjFaMG5TZWp3bWxPNkdGYjBkQjhHSE1EMEI0bHdHUkNoS3NsT0RCSWJzVU51WnptVzZVaThSSEY4OVJZdU9icnhMdDFBT1IrSjR5SXIwc2xDWG52OHNtYU5JNEh2eGNFNGNZcStjOWc3RU9aNmYvSDhlN25Nb01YT2EvVkdVOFhXK1Fkck1DbEZheDJIRjVkMVNmUkNmV2dJZi9tZlUya0o1b3ZiVjNDR09OYVN0TlgvMnRuSGxJd0ZpOEt2Z2QwMWNRdktLOE1ZWDBaeitmN1l5ZHpmNE5vRTcyajNVNTZPeXc1K1lTZndMczFYU0hRNzEzR2dXTFJzazZhU2dLbzc0UHBMOXJ3ejcwOHJPRWlpZFk0WE15OWZjTTFzbkhZcE0vc01IQXBXVDdCQ2ZxTERSUTl4aTlCcmxack1tWTZkVnlhakJPY3o2V0JtQ0dlQWdua0wxMTk2OUJPcUY0d3NFNHpVSDdDVHVETWZDOVEveTdsSVZDanBXYWJNa3JTSFl0cHRseDVHb3RrcDI1Vk8zNVBwK3BOQ0M5RlQ4SGl2V3BHSGdiTjkrZklrVzFrd1FrK1l4RzZjYW8yK2V6OWhmRGlEUFc0c1JKQnV6c1JBRlRtd2x2dmh0bHYyQ1ZGaHZzSk56NTlWemdYWFhOTSs3T011bkEzeGJ4WlpTaEd4OGFOSUo3QVZHM3FVY2RZYldlb1NBPT0="
	data1, _ := base64.StdEncoding.DecodeString(data)
	origin := Decrypt(data1, []byte(desKey[:des.BlockSize]))

	//加密
	Encrypt([]byte(origin), []byte(desKey[:des.BlockSize]))
}
func TestEncrypt(t *testing.T) {
	//解密 secret->decode->decode->des.decrypt->origin
	//加密 origin->des.encrypt->encode->encode->secret
	origin := `{"head":{"firm_code":"","sub_signature_no":"AQY20220719163013499i68X","app_type":"APP_TYPE:H5","emp_num":"","biz_code":"payNotice","subjectStyle":"","main_part_code":"B202103190022","application_id":"YY20220719163549998OBoR","userCode":"HZZ1658213564629","partner_id":"HZZ1658213564629","amountStyle":"","remark9":"","front_env":"","remark8":"","remark5":"","remark10":"","remark4":"","remark7":"","remark6":"","remark1":"4","main_signature_no":"MQY20220719162852682jVG5","remark3":"","remark2":""},"body":{"msg":"TRADE_SUCCESS","code":"0000","data":{"pay_status":"ODR_PAY_STATUS:2","pay_chanel":"alipay","pay_serial_number":"2022080322001400171418376855","order_group_id":"a9dccf96719844ba9aac73bd230","pay_amount":"0.01","pay_card_no":"2088612964400174","order_id":"a9dccf96719844ba9aac73bd230","pay_time":"20220803114836"}}}`
	Encrypt([]byte(origin), []byte(desKey[:des.BlockSize]))	
}
func Encrypt(data, key []byte) {
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	bs := block.BlockSize()

	out := make([]byte, len(data))
    dst := out
    for len(data) > 0 {
        block.Encrypt(dst, data[:bs])
        data = data[bs:]
        dst = dst[bs:]
    }
	
	baseOne := base64.StdEncoding.EncodeToString(out)
	baseTwo := base64.StdEncoding.EncodeToString([]byte(baseOne))

	fmt.Println(baseTwo)
	Write(&baseTwo)
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

	// 返回结果
	// decryptData := make([]byte, len(decode_data))

	// //存储每次解密的数据
	// tmpData := make([]byte, bs)
	// // //分组分块解密
	// for index := 0; index < len(decode_data); index += bs {
	// 	block.Decrypt(tmpData, decode_data[index:index+bs])
	// 	copy(decryptData[index:index+bs], tmpData)
	// }
	// fmt.Println(string(decryptData))
	// res := ""
	// res = string(decryptData)
	// Write(&res)
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
	res := ""
	res = string(out)
	fmt.Println("res:", res)
	Write(&res)
	return out
}

func Write(text *string){

		filePath := "./a.txt"
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("文件打开失败", err)
		}
		//及时关闭file句柄
		defer file.Close()
		//写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)
		write.WriteString(*text)
		//Flush将缓存的文件真正写入到文件中
		write.Flush()
}