package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api/controllers"

	"github.com/gin-gonic/gin"
)

func TestExecuteShellCommandHandler(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Set up a route for the ExecuteShellCommandHandler
	router.POST("/execute", controllers.ExecuteShellCommandHandler)

	// Create a test request body
	requestBody := `{"command": "echo hello"}`

	// Create a new request with the test body
	req, err := http.NewRequest("POST", "/execute", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status 200 but got %d", recorder.Code)
	}

	// Check the response body
	expectedResponse := `{"output":"hello\n"}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("expected response body %s but got %s", expectedResponse, recorder.Body.String())
	}
}

func TestExecuteCommand(t *testing.T) {
	// Create a test context
	ctx := context.Background()

	// Execute a command with a known output
	output, err := controllers.ExecuteCommand(ctx, "echo test")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Check the command output
	expectedOutput := "test\n"
	if output != expectedOutput {
		t.Errorf("expected output %s, got %s", expectedOutput, output)
	}

	// Execute a command that will produce an error
	_, err = controllers.ExecuteCommand(ctx, "unknown-command")
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}
