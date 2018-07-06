//dep_demo/dep_1.go
package main

import(
	"net/http"
	"go.uber.org/zap"
	"github.com/beego/mux"
)

func init(){
	test()
}

func main(){

	logger, _ = zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	mx := mux.New()
	mx.Handler("GET", "/", http.FileServer(http.Dir(".")))
	sugar.Fatal(http.ListenAndServe(("127.0.0.1:8000", mx)))
}