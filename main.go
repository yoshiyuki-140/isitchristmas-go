package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageData struct {
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 現在の日付を取得
		now := time.Now()

		// メッセージを設定
		message := "No"
		if now.Month() == time.December && now.Day() == 25 {
			message = "Yes"
		}

		// テンプレートファイルをパース
		tmpl, err := template.ParseFiles("./html/index.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		// ページデータを作成
		data := PageData{
			Message: message,
		}

		// テンプレートを実行し、レスポンスとして出力
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	})

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
