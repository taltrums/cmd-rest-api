package controllers

import (
	"api/models"
	"context"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ExecuteCommand executes the provided shell command within the given context.
// It returns the command output as a string and any error that occurred during execution.
func ExecuteCommand(ctx context.Context, command string) (string, error) {
	parts := strings.Fields(command)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, parts[0], parts[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// ExecuteShellCommandHandler is an HTTP request handler that accepts a JSON payload
// containing a shell command and executes it. It returns the command output as a response.
func ExecuteShellCommandHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var request models.CommandRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Command == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty command"})
		return
	}
	parts := strings.Fields(request.Command)
	if parts[0] == "sudo" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot run sudo command"})
		return
	}

	output, err := ExecuteCommand(ctx, request.Command)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": exitError.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	response := models.CommandResponse{
		Output: output,
	}

	c.JSON(http.StatusOK, response)
}
