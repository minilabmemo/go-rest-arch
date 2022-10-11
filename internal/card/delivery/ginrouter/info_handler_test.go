package ginrouter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/minilabmemo/go-rest-arch/internal/card/usecase"
	"github.com/minilabmemo/go-rest-arch/internal/config"
	"github.com/minilabmemo/go-rest-arch/internal/models"
)

func init() {
	config.ConfigData = &config.CofigDefinition{Service: config.ServiceInfo{Name: "rest-demo"}} // 測試資料
}

func setUpEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	loadRoutes(engine) //設定你寫路由
	return engine
}

func loadRoutes(engine *gin.Engine) {
	engineGrp := engine.Group("service/api/v1")
	iu := usecase.NewInfoUsecase(*config.ConfigData)
	NewInfoHandler(engineGrp, iu)

}

//go test -v .  //	0.280s	coverage: 100.0% of statements
func TestInfoHandler(t *testing.T) {
	// main
	engine := setUpEngine()
	//StartHttpServer(errs, engine) // 略 不會真的執行

	//模擬發出ＧＥＴ請求
	w := httpTestGetRequest("/service/api/v1/info", engine)

	//驗證發起請求後的回覆是否正確
	assert.Equal(t, http.StatusOK, w.Code)
	t.Log("body", w.Body.String()) //使用go test -v 才可以印出
	//驗證回覆
	expactedResp := `{"name":"rest-demo","version":"1.0.0"}`
	if !reflect.DeepEqual(w.Body.String(), expactedResp) {
		t.Errorf("Resp(%s) != expactedResp(%s)", w.Body.String(), expactedResp)
	}
}

func TestUpdateInfoHandler(t *testing.T) {
	engine := setUpEngine()

	//模擬發出MethodPut請求
	body := models.Info{Name: "test"}
	w := httpTestRequestWithBody(http.MethodPut, "/service/api/v1/info", body, engine)

	//驗證發起請求後的回覆是否正確
	assert.Equal(t, http.StatusOK, w.Code)

	//模擬發出ＧＥＴ請求，再次確認內容有被修改
	w2 := httpTestGetRequest("/service/api/v1/info", engine)
	//驗證發起請求後的回覆是否正確
	assert.Equal(t, http.StatusOK, w2.Code)
	t.Log("body", w2.Body.String()) //使用go test -v 才可以印出
	expactedResp := `{"name":"test","version":"1.0.0"}`
	if !reflect.DeepEqual(w2.Body.String(), expactedResp) {
		t.Errorf("Resp(%s) != expactedResp(%s)", w2.Body.String(), expactedResp)
	}
}

func Test_UpdateInfoBadHandler(t *testing.T) {
	engine := setUpEngine()

	//模擬發出MethodPut請求 //故意放入nil body
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/service/api/v1/info", nil)
	engine.ServeHTTP(w, req)
	t.Log("Code", w.Code)
	//驗證發起請求後的回覆是否正確
	assert.Equal(t, http.StatusBadRequest, w.Code)

	//模擬發出MethodPut請求 //故意放入nil body
	body := models.Info{Name: ""}
	w2 := httpTestRequestWithBody(http.MethodPut, "/service/api/v1/info", body, engine)

	t.Log("Code", w2.Code)
	//驗證發起請求後的回覆是否正確
	assert.Equal(t, http.StatusExpectationFailed, w2.Code)
}

// 模擬MethodGet發生
func httpTestGetRequest(url string, engine *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()                         // 利用 httptest 取得 ResponseRecorder 物件
	req, _ := http.NewRequest(http.MethodGet, url, nil) //使用 http 建立一個 Request
	engine.ServeHTTP(w, req)                            //使用 gin 提供的 ServeHTTP 模擬 API 呼叫．response 會放入ResponseRecorder
	defer w.Result().Body.Close()
	return w
}

//模擬帶json body的請求發生
func httpTestRequestWithBody(method, url string, body interface{}, engine *gin.Engine) *httptest.ResponseRecorder {
	jsonByte, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, bytes.NewReader(jsonByte))
	engine.ServeHTTP(w, req)
	defer w.Result().Body.Close()
	return w
}
