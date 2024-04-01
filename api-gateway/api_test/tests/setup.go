package tests

import (
  "api-gateway/api_test/storage/kv"
  "bytes"
  "github.com/gin-gonic/gin"
  "github.com/spf13/viper"
  "net/http"
  "net/http/httptest"
  "os"
)

const testHost = "http://localhost"

func SetupMinimumInstance(path string) error {
  _ = path
  cnf := viper.New()
  cnf.Set("mode", "test")
//   db := storage.ConnDB()
  kv.Init(kv.NewInMemoryInst())
  return nil
}

func Serve(handler gin.HandlerFunc, uri string, req *http.Request, router *gin.Engine) (*httptest.ResponseRecorder, error) {

  resp := httptest.NewRecorder()

  // Define route based on the request method
  switch req.Method {
  case http.MethodPost:
    router.POST(uri, handler)
  case http.MethodGet:
    router.GET(uri, handler)
  case http.MethodDelete:
    router.DELETE(uri, handler)
  case http.MethodPatch:
    router.PATCH(uri, handler)
  }

  router.ServeHTTP(resp, req)
  return resp, nil
}

func NewResponse() *http.Response {
  return httptest.NewRecorder().Result()
}

func NewRequest(method string, uri string, body []byte) *http.Request {
  req, _ := http.NewRequest(method, testHost+uri, bytes.NewReader(body))
  req.Header.Set("Host", "localhost")
  req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
  req.Header.Set("X-Forwarded-For", "79.104.42.249")
  return req
}

func OpenFile(fileName string) ([]byte, error) {
  return os.ReadFile(fileName)
}
