package main

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"testing"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

func Test_Pic(t *testing.T) {
	//生成图片
	picCreate()
	//dc.SetColor(color.White)
	//dc.SetHexColor("#19344180")
	// picCreate1()

}

func picCreate() {
	code := GetRandStr(4)
	res := ImgText(200, 60, code)
	str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(res)
	Write(&str)
}

func ImgText(width, height int, text string) (b []byte) {
	textLen := len(text)
	dc := gg.NewContext(width, height)
	// bgR, bgG, bgB, bgA := getRandColorRange(240, 255)
	// dc.SetRGBA255(bgR, bgG, bgB, bgA)
	dc.SetRGBA255(0, 0, 0, 0)
	dc.Clear()

	// 干扰线
	// for i := 0; i < 10; i++ {
	// 	x1, y1 := getRandPos(width, height)
	// 	x2, y2 := getRandPos(width, height)
	// 	r, g, b, a := getRandColor(255)
	// 	w := float64(rand.Intn(3) + 1)
	// 	dc.SetRGBA255(r, g, b, a)
	// 	dc.SetLineWidth(w)
	// 	dc.DrawLine(x1, y1, x2, y2)
	// 	dc.Stroke()
	// }

	fontSize := float64(height/2) + 5
	face := loadFontFace(fontSize)
	dc.SetFontFace(face)
	colors := [4]string{"#f4275d","#00dadc", "#ebd36c", "#bb6b4e"}
	for i := 0; i < len(text); i++ {
		// r, g, b, _ := getRandColor(100)
		// dc.SetRGBA255(r, g, b, 255)
		dc.SetHexColor(colors[RandNumber(4)])
		fontPosX := float64(width/textLen*i) + fontSize*0.6

		writeText(dc, text[i:i+1], float64(fontPosX), float64(height/2))
	}
	buffer := bytes.NewBuffer(nil)
	dc.EncodePNG(buffer)
	b = buffer.Bytes()
	return
}

func GetRandStr(n int) (randStr string) {
    chars := "abcdefghijkmnpqrstuvwxyz23456789"
    charsLen := len(chars)
    if n > 10 {
        n = 10
    }
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < n; i++ {
        randIndex := rand.Intn(charsLen)
        randStr += chars[randIndex : randIndex+1]
    }
    return randStr
}

// 渲染文字
func writeText(dc *gg.Context, text string, x, y float64) {
	xfload := 5 - rand.Float64()*10 + x
	yfload := 5 - rand.Float64()*10 + y

	radians := 40 - rand.Float64()*80
	dc.RotateAbout(gg.Radians(radians), x, y)
	dc.DrawStringAnchored(text, xfload, yfload, 0.2, 0.5)
	dc.RotateAbout(-1*gg.Radians(radians), x, y)
	dc.Stroke()
}


// 加载字体
func loadFontFace(points float64) font.Face {
	//  // 这里是将字体TTF文件转换成了 byte 数据保存成了一个 go 文件 文件较大可以到附录下
	// // 通过truetype.Parse可以将 byte 类型的数据转换成TTF字体类型
	// f, err := ioutil.ReadFile("dk.ttf")
	// if err != nil {
	//     fmt.Println("read fail:", err)
	// }
	font, err := truetype.Parse(goregular.TTF)

	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: points,
	})
	return face
}

/* 
github.com/mojocn/base64Captcha
*/
