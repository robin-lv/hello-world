package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloRoute(t *testing.T) {
	// 设置 Gin 路由
	router := setupRouter()

	// 创建一个 HTTP 请求
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 创建一个 HTTP 响应记录器
	w := httptest.NewRecorder()

	// 发送请求并记录响应
	router.ServeHTTP(w, req)

	// 检查状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 检查响应体
	expectedResponse := `{"message":"Hello, World!"}`
	assert.Equal(t, expectedResponse, w.Body.String())
}
