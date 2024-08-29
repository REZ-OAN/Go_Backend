package api

import (
	"os"
	"testing"
	"time"

	db "github.com/REZ-OAN/simplebank/database/sqlc"
	"github.com/REZ-OAN/simplebank/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TOKEN_SYM_KEY:         utils.RandomString(32),
		ACCESS_TOKEN_DURATION: time.Minute,
	}
	server, err := NewServer(store, config)
	require.NoError(t, err)
	return server
}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
