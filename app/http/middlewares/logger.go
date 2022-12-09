package middlewares

import (
	"bytes"
	"io"
	"time"

	"github.com/spf13/cast"

	"sample/pkg/logger"

	"sample/pkg/helpers"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// responseBodyWriter 装饰器模式
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseBodyWriter) Writer(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		w := &responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		ctx.Writer = w

		var requestBody []byte
		if ctx.Request.Body != nil {
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			// body 只能读取一次
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		start := time.Now()
		ctx.Next()
		cost := time.Since(start)

		// 状态码
		status := ctx.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", status),
			zap.String("request", ctx.Request.Method+" "+ctx.Request.URL.String()),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}

		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			// 请求的内容
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))

			// 响应的内容
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if status > 400 && status <= 499 {
			// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
			logger.Warn("HTTP Warning "+cast.ToString(status), logFields...)
		} else if status >= 500 && status <= 599 {
			// 除了内部错误，记录 error
			logger.Error("HTTP Error "+cast.ToString(status), logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}
}
