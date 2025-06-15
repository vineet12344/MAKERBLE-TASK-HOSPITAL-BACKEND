package middleware

import (
	"hospital-backend/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func init() {
	os.Setenv("JWT_SECRET", "testsecret123")
}

func performRequestWithToken(t *testing.T, token string, handler gin.HandlerFunc) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	c.Request = req

	handler(c)
	return w
}

func TestRequireAuth_MissingToken(t *testing.T) {
	rec := performRequestWithToken(t, "", RequireAuth())
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestRequireAuth_ValidToken(t *testing.T) {
	token, _ := utils.GenerateToken(1, "test@doc.com", "doctor")
	rec := performRequestWithToken(t, token, RequireAuth())
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCheckRole_CorrectRole(t *testing.T) {
	token, _ := utils.GenerateToken(1, "a@b.com", "doctor")

	rec := performRequestWithToken(t, token, func(c *gin.Context) {
		// Simulate middleware chain
		RequireAuth()(c)
		if c.IsAborted() {
			return
		}

		CheckRole("doctor")(c)
		if c.IsAborted() {
			return
		}

		// If all middleware passes, return OK
		c.String(http.StatusOK, "Authorized")
	})

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Authorized", rec.Body.String())
}

func TestCheckRole_WrongRole(t *testing.T) {
	token, _ := utils.GenerateToken(1, "x@y.com", "receptionist")
	rec := performRequestWithToken(t, token, func(c *gin.Context) {
		RequireAuth()(c)
		if c.IsAborted() {
			return
		}
		CheckRole("doctor")(c)
	})
	assert.Equal(t, http.StatusForbidden, rec.Code)
}
