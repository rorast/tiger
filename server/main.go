package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
)

type (
	// 建立一個樂透控制器
	lotterController struct {
		Ctx iris.Context
	}

	// 傳給前端的結果及下次轉盤的數字
	ResultResponse struct {
		Code       []int  `json:"code"`
		Result     []int  `json:"result"`
		NextSymbol []int  `json:"symbol"`
	}
)


// 重啟服務
func reLaunch() {
	cmd := exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}


// 啟動一個http server
func deployPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1> Hello deplay success</h2>")
	reLaunch()  // 重啟新服務
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotterController{})
	return app
}

// 既開既得型 http://localhost:8088
func (c *lotterController) Get() string {
	var prize string
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Intn(10)
	switch {
	case code == 1 :
		prize = "一等獎"
	case code >= 2 && code <= 3:
		prize = "二等獎"
	case code >=4 && code <= 6:
		prize = "三等獎"
	default:
		return fmt.Sprintf("尾號為1獲得一等獎<br/>"+
			"尾號為2或者3獲得二等獎<br/>"+
			"尾號為4/5/6獲得三等獎<br/>"+
			"code=%d<br/>"+
			"很遺憾，沒有獲獎", code)
	}
	return fmt.Sprintf("尾號為1獲得一等獎<br/>"+
		"尾號為2或者3獲得二等獎<br/>"+
		"尾號為4/5/6獲得三等獎<br/>"+
		"code=%d<br/>"+
		"恭禧您獲得:%s", code,prize)
}

// 開獎 vatility L1 ~ L5
func (c *lotterController) GetPrize() string {
	//seed := time.Now().UnixNano()
	//r := rand.New(rand.NewSource(seed))
	//// 先隨機計算結果給前端 0 :=沒中，1 := 中彩金，2 :=freeSpin, 3 := bigWin, 4 := jackPort
	//var resultCode [5]int
	//for i :=0 ; i < len(resultCode) ; i ++ {
	//	resultCode[i] = r.Intn(5)+1
	//}
	//var prize [5][3]int
	//// 15個滾輪，13個symbol 1-13
	//for i:=0 ; i < 5 ; i++ {
	//	for j:=0 ; j < 3 ; j++ {
	//		prize[i][j] = r.Intn(13)+1
	//	}
	//	//prize[i] = r.Intn(33)+1
	//}
	//// reel為25個symbol， 有 5 個 reel 帶
	//var nextPrize [5][25]int
	//for i:=0 ; i < 5 ; i++ {
	//	for j:=0 ; j < 25 ; j++ {
	//		nextPrize[i][j] = r.Intn(13)+1
	//	}
	//	//prize[i] = r.Intn(33)+1
	//}
	//
	//return fmt.Sprintf("code = : %v",resultCode,"result = : %v", prize,"symbols = : %v", nextPrize)


	rrr := ResultResponse{}
	rrr.Code = make([]int, 5)
	rrr.Result = make([]int,15)

	rrr.NextSymbol = make([]int,125)
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	//先隨機計算結果給前端 0 :=沒中，1 := 中彩金，2 :=freeSpin, 3 := bigWin, 4 := jackPort
	var resultCode [5]int
	for i :=0 ; i < len(resultCode) ; i ++ {
		rrr.Code[i] = r.Intn(5)+1
	}

	// 15個滾輪，13個symbol 1-13
	for i:=0 ; i < len(rrr.Result) ; i++ {
		rrr.Result[i] = r.Intn(13)+1
	}

	// reel為25個symbol， 有 5 個 reel 帶
	for i:=0 ; i < len(rrr.NextSymbol) ; i++ {
		rrr.NextSymbol[i] = r.Intn(13)+1
	}
	fmt.Println("Spin:",rrr.String())

	return "111"
}

func (conf *ResultResponse) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}

func main() {

	app := newApp()
	app.Run(iris.Addr(":8088"))

	// 執行 bash 程式
	//http.HandleFunc("/deploy",deployPage)
	//http.ListenAndServe(":3450", nil)
}
