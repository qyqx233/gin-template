package controller

import (
	"bytes"
	"encoding/json"
	"gin-template/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	model.InitDB()
	m.Run()
}

func TestA(t *testing.T) {
	testNewTranslate(t)
	// t.Log(2323)
}

func TestQueryTranslate(t *testing.T) {
	testQueryTranslate(t)
}

func TestNewTranslate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/translate", NewTranslate)

	t.Run("successful request", func(t *testing.T) {
		reqBody := TranslateRq{
			OriginalText:   "Hello",
			TranslatedText: "Hola",
			Direction:      "en-es",
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/translate", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message":"success"}`, w.Body.String())
	})

	t.Run("invalid request", func(t *testing.T) {
		reqBody := `{"originalText": "Hello"}` // Missing translatedText and direction
		req, _ := http.NewRequest(http.MethodPost, "/translate", bytes.NewBuffer([]byte(reqBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error"`)
	})
}
