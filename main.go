package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/healthz", HealthHandler)
	http.HandleFunc("/errexample", ErrorResponseHandler)
	http.ListenAndServe(":80", nil)
}

// RootHandler 默认处理方法
func RootHandler(w http.ResponseWriter, r *http.Request) {

	r.Header.Add("runtime version", runtime.Version())
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, ","))
	}
	//w.WriteHeader(200)要写在Header().set方法后，不然无法解析
	w.WriteHeader(200)
	fmt.Fprintln(w, w.Header())
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	//获取客户端IP，并打印
	ClientIp := strings.Split(r.RemoteAddr, ":")[0]
	if len(ClientIp) == 1 {
		fmt.Println("Client ip is : ", "127.0.0.1", " statusCode is : ", "200")
	} else {
		fmt.Println("Client ip is : ", ClientIp, " statusCode is : ", "200")
	}

	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(body))
}

// ErrorResponseHandler /errexample的访问处理方法
func ErrorResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	fmt.Fprintln(w, "Server error.")
	//获取客户端IP，并打印
	ClientIp := strings.Split(r.RemoteAddr, ":")[0]
	if len(ClientIp) == 1 {
		fmt.Println("Client ip is : ", "127.0.0.1", " statusCode is : ", "200")
	} else {
		fmt.Println("Client ip is : ", ClientIp, " statusCode is : ", "200")
	}
}

// HealthHandler /healthz的访问处理方法
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprint(w, `{"statusCode":200}`)
	//获取客户端IP，并打印
	ClientIp := strings.Split(r.RemoteAddr, ":")[0]
	if len(ClientIp) == 1 {
		fmt.Println("Client ip is : ", "127.0.0.1", " statusCode is : ", "200")
	} else {
		fmt.Println("Client ip is : ", ClientIp, " statusCode is : ", "200")
	}
}
