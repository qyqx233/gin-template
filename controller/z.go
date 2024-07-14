package controller

import (
	"bytes"
	"encoding/json"
	"gin-template/model"
	"gin-template/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type TranslateRq struct {
	OriginalText   string `json:"originalText"`
	TranslatedText string `json:"translatedText"`
	Direction      string `json:"direction"`
}

// write code to test NewTranslate
func testNewTranslate(_ *testing.T) {
	// create a new router
	r := gin.Default()
	// register the NewTranslate handler
	r.POST("/translate", NewTranslate)
	// create a new request
	reqBody := TranslateRq{
		OriginalText:   "Hello",
		TranslatedText: "Hola",
		Direction:      "en-es",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/translate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	// create a new response recorder
	w := httptest.NewRecorder()
	// serve the request
	r.ServeHTTP(w, req)
	// assert the response
}

func NewTranslate(c *gin.Context) {
	var rq = new(TranslateRq)
	var err error
	if err = util.ShouldBindJSON(c, rq, false); err != nil {
		return
	}
	t := &model.Translate{
		OriginalText:   rq.OriginalText,
		TranslatedText: rq.TranslatedText,
		Direction:      rq.Direction,
	}
	if err = t.Create(); err != nil {
		return
	}
	util.RenderSuccess(c)
}

func QueryTranslate(c *gin.Context) {
	ts, err := model.SearchTranslates()
	if err != nil {
		util.RenderError(c, err)
		return
	}
	c.JSON(http.StatusOK, ts)
}

func testQueryTranslate(t *testing.T) {
	// write code
	// create a new router
	r := gin.Default()
	// register the QueryTranslate handler
	r.GET("/translate", QueryTranslate)
	// create a new request
	req, _ := http.NewRequest(http.MethodGet, "/translate", nil)
	// create a new response recorder
	w := httptest.NewRecorder()
	// serve the request
	r.ServeHTTP(w, req)
	// 输出结果
	t.Log(w.Body.String())
}
