package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PingModelMock struct {
	mock.Mock
}

func (p *PingModelMock) getPing() string {
	args := p.Called()
	return args.Get(0).(string)
}

func (p *PingModelMock) getPong() string {
	args := p.Called()
	return args.Get(0).(string)
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

//Method 이름이 Testxxxx로 시작해야 test가 작동
func TestPingRouter(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		pingpong := PingPong{"ping", "pong"}
		body := gin.H{
			"message": pingpong.getPing(),
		}

		mockPingModel := new(PingModelMock)
		mockPingModel.On("getPong").Return("ping")

		router := gin.Default()
		c := &Config{
			R:         router,
			PingModel: mockPingModel,
		}
		setupRouter(router, c)
		request, err := http.NewRequest(http.MethodGet, "/game/ping", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)

		fmt.Println(rr.Body.String())

		var response map[string]string
		jsonErr := json.Unmarshal([]byte(rr.Body.String()), &response)

		assert.Nil(t, jsonErr)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, body["message"], response["message"])
		mockPingModel.AssertExpectations(t)
	})

	// w := performRequest(router, "GET", "/game/ping")

	// assert.Equal(t, http.StatusOK, w.Code)

	// var response map[string]string
	// err := json.Unmarshal([]byte(w.Body.String()), &response)

	// value, exists := response["message"]

	// assert.Nil(t, err)
	// assert.True(t, exists)
	// assert.Equal(t, body["message"], value)
}
