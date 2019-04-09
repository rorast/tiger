package main

import (
	"io"
	"net/http"
	"os/exec"
	"log"
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



func main() {
	http.HandleFunc("/deploy",deployPage)
	http.ListenAndServe(":3450", nil)
}