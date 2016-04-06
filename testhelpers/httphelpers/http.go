package httphelpers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

// Ways to get data from *gin.Context
// Param
// DefaultQuery
// Query
// PostForm
// DefaultPostForm

// MockRequest is an interface for mocking requests
type MockRequest interface {
	MakeRequest(*RequestOptions) *httptest.ResponseRecorder
}

// TestRequest holds info for mocking a request
type TestRequest struct {
	Method string
	Path   string
	Router *gin.Engine
}

// RequestOptions hold test options to be set when making a request
type RequestOptions struct {
	Params   map[string]string
	PostForm map[string]string
	Headers  map[string]string
}

// AddParams adds a key, value to Params
func (o *RequestOptions) AddParams(key string, value string) {
	o.Params[key] = value
}

// AddPostForm adds a key, value to PostForm
func (o *RequestOptions) AddPostForm(key string, value string) {
	o.PostForm[key] = value
}

// AddHeaders adds a key, value to Headers
func (o *RequestOptions) AddHeaders(key string, value string) {
	o.Headers[key] = value
}

// NewRequestOptions returns a new RequestOptions
func NewRequestOptions() (o *RequestOptions) {
	o = new(RequestOptions)
	o.Params = make(map[string]string)
	o.PostForm = make(map[string]string)
	o.Headers = make(map[string]string)
	return
}

func constructQuery(params map[string]string) (query string) {
	for k, v := range params {
		query += k + "=" + v + "&"
	}
	return
}

func addHeadersToRequest(req *http.Request, o *RequestOptions) {
	for k, v := range o.Headers {
		req.Header.Add(k, v)
	}
}

// MakeRequest makes a http request and returns the ResponseRecorder
func (r *TestRequest) MakeRequest(o *RequestOptions) *httptest.ResponseRecorder {
	query := "?"
	path := r.Path
	query += constructQuery(o.Params)
	if query != "?" {
		path += query
	}
	formQuery := ""
	formQuery += constructQuery(o.PostForm)
	body := bytes.NewBufferString(formQuery)
	req, _ := http.NewRequest(r.Method, path, body)
	if r.Method == "POST" || r.Method == "PUT" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	addHeadersToRequest(req, o)
	resp := httptest.NewRecorder()
	r.Router.ServeHTTP(resp, req)
	return resp
}

func mockRoute(router *gin.Engine, method string, path string, handler func(c *gin.Context)) {
	switch {
	case method == "GET":
		router.GET(path, handler)
	case method == "POST":
		router.POST(path, handler)
	case method == "PUT":
		router.PUT(path, handler)
	case method == "DELETE":
		router.DELETE(path, handler)
	}
}

// MockHTTPRequest sets up a gin server with only the handler action to test it
func MockHTTPRequest(handler func(c *gin.Context), method string, callback func(MockRequest)) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	testPath := "/test_path"
	method = strings.ToUpper(method)

	mockRoute(router, method, testPath, handler)

	testRequest := new(TestRequest)
	testRequest.Method = method
	testRequest.Path = testPath
	testRequest.Router = router
	callback(testRequest)
}
