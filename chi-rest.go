package main

import (
  "net/http"
  "fmt"
  "encoding/json"

  "github.com/go-chi/chi/v5"
  // "github.com/go-chi/chi/v5/middleware"
)

// Articleデータ構造体定義
type ArticleData struct {
  Id int          `json:"id"`
  Title string    `json:"title"`
  Content string  `json:"content"`
  Author string   `json:"author"`
}

func main() {
  r := chi.NewRouter()

  // "articles"以下のURLをルーティング
  r.Route("/articles", func(r chi.Router) {

    r.Post("/", createArticle)                            // POST /articles
    // r.Get("/search", searchArticles)                   // GET /articles/search

    // 正規表現を使ったURLパラメータも可能:
    // r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)  // GET /articles/home-is-toronto

    // サブルータ:
    r.Route("/{articleID}", func(r chi.Router) {
      r.Get("/", getArticle)                               // GET /articles/123
      r.Put("/", updateArticle)                            // PUT /articles/123
      r.Delete("/", deleteArticle)                         // DELETE /articles/123
    })
  })

  // サーバーリッスン
  http.ListenAndServe(":3000", r)
}

// GET /articles/xxxx ハンドラー
func getArticle(w http.ResponseWriter, r *http.Request) {

  // URLパラメータ取得
  articleID := chi.URLParam(r, "articleID")
  fmt.Printf("article id = %v\n", articleID)

  // 記事データ作成
  aData := &ArticleData{}
  aData.Id = 1
  aData.Title = "Test Artcle"
  aData.Content = "This is test !"
  aData.Author = "Anonymous"

  // 記事データをJSONに変換
  encoder := json.NewEncoder(w)
  err := encoder.Encode(aData)
  if err != nil {
    fmt.Print("json encoder error\n")
  }

  // article, err := dbGetArticle(articleID)
  // if err != nil {
  //   http.Error(w, http.StatusText(404), 404)
  //   return
  // }

  // w.Write([]byte(fmt.Sprintf("title:%s", article.Title)))
}

// POST /articles ハンドラー
func createArticle(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("article created!\n")
}

// PUT /articles/xxxxxx ハンドラー
func updateArticle(w http.ResponseWriter, r *http.Request){
  fmt.Printf("article updated!\n")
}

// DELETE /articles/xxxxxx ハンドラー
func deleteArticle(w http.ResponseWriter, r *http.Request){
  fmt.Print("article deleted!\n")
}
