package test

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"wobbs-server/config"
	"wobbs-server/router"
)

func Init() {
	var configPath string
	flag.StringVar(&configPath, "f", "../config/config.yaml", "配置文件路径")
	flag.Parse()
	// 初始化配置
	config.InitConfig(configPath)
	// 初始化日志
	config.InitLogger(config.Conf.LogConfig)
	// 初始化数据库
	config.InitDB(config.Conf.MySQLConfig)
}

func TestLoginApi(t *testing.T) {
	Init()
	gin.SetMode(gin.TestMode)
	r := router.InitRouter()
	//w := httptest.NewRecorder()
	data := make(map[string]interface{})
	data["username"] = "admin"
	data["password"] = "admin"
	//jsonData, _ := json.Marshal(data)

	//req, _ := http.NewRequest(http.MethodPost, "/api/user/login", bytes.NewReader(jsonData))
	//r.ServeHTTP(w, req)

	//assert.Equal(t, 200, w.Code)
	//fmt.Println(w)
	//assert.Equal(t, "pong", w.Body.String())
	body := PostJson("/api/user/login", data, r, t)
	fmt.Println(body)
}

// PostJson 根据特定请求uri和参数param，以Json形式传递参数，发起post请求返回响应
func PostJson(uri string, param map[string]interface{}, router *gin.Engine, t *testing.T) []byte {
	// 将参数转化为json比特流
	jsonByte, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}
	// 构造post请求，json数据以请求body的形式传递
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	fmt.Println(w)
	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
