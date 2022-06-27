package middlewares

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type logRespWriter struct {
	gin.ResponseWriter
	respBody *bytes.Buffer
}

func (w logRespWriter) Write(b []byte) (int, error) {
	w.respBody.Write(b)
	return w.ResponseWriter.Write(b)
}

func logMiddleware(c *gin.Context) {
	if c.Request.Method == "GET" || c.Request.Method == "OPTIONS" || c.Request.Method == "HEAD" {
		c.Next()
		return
	}

	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
	logWriter := &logRespWriter{respBody: bytes.NewBuffer(make([]byte, 0, 512)), ResponseWriter: c.Writer}
	c.Writer = logWriter

	c.Next()

	statusCode := c.Writer.Status()
	log.Info().Int("statusCode", statusCode).
		Str("reqBody", string(reqBody)).
		Str("respBody", logWriter.respBody.String()).
		Msg("HTTP request completed")
}
