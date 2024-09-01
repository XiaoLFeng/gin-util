package bstruct

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
)

type CustomWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
	size int
}

func (w *CustomWriter) Write(b []byte) (int, error) {
	n, _ := w.Body.Write(b)
	w.size += n
	return w.size, nil
}

func (w *CustomWriter) WriteConfirm() (int, error) {
	_, err := io.WriteString(w.ResponseWriter, w.Body.String())
	return w.size, err
}

func (w *CustomWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}
