package responsewriters

import (
	"compress/gzip"
	"net/http"
)

type GzipResponseWriter struct {
	rw http.ResponseWriter
	w  gzip.Writer
}

func NewGzipResponseWriter(w http.ResponseWriter) (grw GzipResponseWriter) {
	return GzipResponseWriter{
		rw: w,
		w:  *gzip.NewWriter(w),
	}
}

func (grw *GzipResponseWriter) Close() error {
	return grw.w.Close()
}

func (grw *GzipResponseWriter) Write(input []byte) (int, error) {
	return grw.w.Write(input)
}

func (grw *GzipResponseWriter) Header() http.Header {
	return grw.rw.Header()
}

func (grw *GzipResponseWriter) WriteHeader(statusCode int) {
	grw.rw.WriteHeader(statusCode)
}
