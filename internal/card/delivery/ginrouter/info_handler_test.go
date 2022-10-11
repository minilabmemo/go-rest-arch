package ginrouter

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/minilabmemo/go-rest-arch/internal/card/usecase"
	"github.com/minilabmemo/go-rest-arch/internal/config"
)

func init() {
	config.ConfigData = &config.CofigDefinition{Service: config.ServiceInfo{Name: "rest-demo"}} // 測試資料
}

//go test -v .
func TestStartHttpServer(t *testing.T) {
	// main
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	loadRoutes(engine) //取得路由
	//StartHttpServer(errs, engine) // 略 不會真的執行

	w := httptest.NewRecorder()                                   // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/service/api/v1/info", nil) //使用 http 建立一個 Request
	engine.ServeHTTP(w, req)                                      //使用 gin 提供的 ServeHTTP 模擬 API 呼叫．response 會放入ResponseRecorder

	//驗證發起請求後的回覆是否正確
	assert.Equal(t, http.StatusOK, w.Code)
	t.Log("body", w.Body.String()) //使用go test -v 才可以印出
	expactedResp := `{"name":"rest-demo","version":"1.0.0"}`
	if !reflect.DeepEqual(w.Body.String(), expactedResp) {
		t.Errorf("Resp(%s) != expactedResp(%s)", w.Body.String(), expactedResp)
	}
}

func loadRoutes(engine *gin.Engine) {
	engineGrp := engine.Group("service/api/v1")

	iu := usecase.NewInfoUsecase(*config.ConfigData)
	NewInfoHandler(engineGrp, iu)

}
