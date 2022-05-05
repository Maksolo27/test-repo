package main
import (
  "net/http"
  "time"
  "encoding/json"
) 
type Time struct {
  NowTime string `json:"time"`
} 
func main() {
  const RFC3339 = "2006-01-02T15:04:05Z07:00" 
  http.HandleFunc("/time", func(rw http.ResponseWriter, r *http.Request) {
    now := time.Now()
    formattedTime := now.Format(RFC3339) 
    structTime := Time{formattedTime} 
    data, _ := json.MarshalIndent(structTime, "", "\t")
    rw.Write(data)
  })
  http.ListenAndServe(":8795", nil)
}

