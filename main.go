package main

import (
	"fmt"
	"go_todo_app/todo_app/config"
	"log"
	"net/http" // config パッケージをインポート
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Docker!")
}

func main() {
	// 設定値の出力
	fmt.Println("Configurations:")
	fmt.Println("Port:", config.Config.Port)
	fmt.Println("SQLDriver:", config.Config.SQLDriver)
	fmt.Println("DbName:", config.Config.DbName)
	fmt.Println("LogFile:", config.Config.LogFile)
  log.Println("test")

	// HTTP サーバーの起動
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
