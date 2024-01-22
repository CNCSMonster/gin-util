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
	// 将响应体的内容同时写入缓存区
	w.buffer.Write(data)
	// 调用原始的Write方法，将响应体写入HTTP响应中
	return w.ResponseWriter.Write(data)
}

func (w *BufferedResponseWriter) Base() gin.ResponseWriter {
	return w.ResponseWriter
}
func (w *BufferedResponseWriter) Data() []byte {
	return w.buffer.Bytes()
}
