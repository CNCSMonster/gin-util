package responsewriter

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type BufferedResponseWriter struct {
	gin.ResponseWriter
	buffer *bytes.Buffer
}

func NewBufferedResponseWriter(w gin.ResponseWriter) *BufferedResponseWriter {
	return &BufferedResponseWriter{w, bytes.NewBufferString("")}
}
func (w *BufferedResponseWriter) Write(data []byte) (int, error) {
	w.buffer.Write(data)
	return w.ResponseWriter.Write(data)
}

func (w *BufferedResponseWriter) Base() gin.ResponseWriter {
	return w.ResponseWriter
}
func (w *BufferedResponseWriter) Data() []byte {
	return w.buffer.Bytes()
}
