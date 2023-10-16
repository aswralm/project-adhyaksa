package integrationtest_test

import (
	"context"
	"net/http/httptest"
	createid "project-adhyaksa/pkg/create-id"
	"project-adhyaksa/services/event/internal/repository/model"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
)

func TestRegisterEvent(t *testing.T) {
	database()
	gin.SetMode(gin.TestMode)

	//refreshEventTable()

	defer DB.Close()

	//create branch
	branch := model.Branch{
		Name:      "branch123",
		ID:        createid.CreateID(),
		Address:   "jakarta",
		CreatedAt: time.Now(),
	}
	insertBranchTable(branch, context.TODO())

	samples := []struct {
		testname   string
		inputJSON  string
		statusCode int
		success    bool
		message    string
		data       interface{}
	}{
		{
			testname: "positive case",
			inputJSON: `{
				"branch_id": "branch123",
				"name": "pertemuan seluruh indonesia",
				"start_time": "2023-10-01T08:00:00Z",
				"end_time": "2023-10-01T17:00:00Z",
				"location": "jakarta",
				"description": "meningkatkan silaturahmi"
			}`,
			statusCode: 201,
			success:    true,
			message:    "Register event successfully",
			data:       nil,
		},
	}

	for _, v := range samples {
		handler := gin.New()

		server := httptest.NewServer(handler)
		server.Listener.Close()
		server.Listener = httptest.NewServer(handler).Listener

		e := httpexpect.WithConfig(httpexpect.Config{
			TestName: v.testname,
			BaseURL:  server.URL,
			Reporter: httpexpect.NewAssertReporter(t),
		})

		resp := e.POST("/api/v1/event").
			WithJSON(v.inputJSON).
			Expect().
			Status(v.statusCode)

		if v.success {
			resp.JSON().Object().HasValue("message", v.message)
		} else {
			resp.JSON().Object().HasValue("error", v.message)
		}

		// Jika perlu verifikasi data yang dikembalikan
		if v.data != nil {
			resp.JSON().Object().ContainsSubset(v.data.(map[string]interface{}))
		}
	}
}
