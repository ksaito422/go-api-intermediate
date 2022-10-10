package middlewares

import (
	"context"
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := req.Context()
		ctx = context.WithValue(ctx, traceIDKey{}, traceID)
		req = req.WithContext(ctx)
		// 自作のResponseWriterを作って
		rlw := NewResLoggingWriter(w)

		// それをハンドラに渡す
		next.ServeHTTP(rlw, req)

		// 自作ResponseWriterからロギングしたいデータを出す
		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}
