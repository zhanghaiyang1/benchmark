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

/*
test
*/
var (
	apiRZ = "https://sandbox-gw.insuremo.com.cn"
)

func TestGetLink(t *testing.T) {
	api := apiRZ
	/*

		1.将HTTP Body的内容和secret key进行字符串拼接；
		2.对于拼接后的内容使用MD5算法进行签名运算。
	*/
	//
}
