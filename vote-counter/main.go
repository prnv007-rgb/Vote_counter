package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strings" // ✅ Added
  "time"
)

var counts = map[string]int{"A": 0, "B": 0}

func voteHandler(w http.ResponseWriter, r *http.Request) {
  opt := r.URL.Query().Get("option")
  if opt != "A" && opt != "B" {
    countsHandler(w, r)
    return
  }
  counts[opt]++
  go sendLog(opt)
  w.WriteHeader(http.StatusOK)
}

func countsHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(counts)
}

func sendLog(opt string) {
  payload := fmt.Sprintf(`{"option":"%s","time":"%s"}`, opt, time.Now().Format(time.RFC3339))
  http.Post("http://vote-logger:4567/log", "application/json", strings.NewReader(payload))
}

func main() {
  http.HandleFunc("/vote", voteHandler)
  http.HandleFunc("/counts", countsHandler)
  log.Println("Vote counter running on :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
