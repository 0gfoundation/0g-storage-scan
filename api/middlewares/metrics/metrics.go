package metrics

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	nhMetrics "github.com/0glabs/0g-storage-scan/metrics"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/http/middlewares"
	"github.com/sirupsen/logrus"
)

type ResponseWriter struct {
	writer http.ResponseWriter
	code   int
	// cache response data
	buf *bytes.Buffer
}

func newResponseWriter(writer http.ResponseWriter, code int) *ResponseWriter {
	var buf bytes.Buffer
	return &ResponseWriter{
		writer: writer,
		code:   code,
		buf:    &buf,
	}
}

func (w *ResponseWriter) Write(bs []byte) (int, error) {
	w.buf.Write(bs)
	return w.writer.Write(bs)
}

func (w *ResponseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.code = code
	w.writer.WriteHeader(code)
}

func Metrics() middlewares.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrappedW := newResponseWriter(w, http.StatusOK)
			next.ServeHTTP(wrappedW, r)

			if len(wrappedW.buf.Bytes()) > 0 {
				urlType, ok := r.Context().Value(CtxKeyURLType).(string)
				if !ok {
					return
				}

				var resp api.BusinessError
				err := json.Unmarshal(wrappedW.buf.Bytes(), &resp)
				if err != nil {
					return
				}

				UpdateDuration(urlType, wrappedW.code, resp.Code, start)
				logrus.WithFields(logrus.Fields{
					"Path":         r.URL.Path,
					"urlType":      urlType,
					"http status":  wrappedW.code,
					"resp code":    resp.Code,
					"resp message": resp.Message,
				}).Debug("Report API metrics")
			}
		})
	}
}

func UpdateDuration(url string, status, code int, start time.Time) {
	path := url[5:]
	nhMetrics.Registry.API.UpdateDuration(path, status, code, start)
}

var CtxKeyURLType = middlewares.CtxKey("X-URL-TYPE")

func URLType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := validPath(r.URL.Path)

		var urlType string
		if url != nil {
			urlType = *url
		} else {
			urlType = "invalid url"
		}

		ctx := context.WithValue(r.Context(), CtxKeyURLType, urlType)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

var URLPattern = map[*regexp.Regexp]string{
	regexp.MustCompile(`^/api/accounts/0[xX][0-9a-fA-F]{40}$`):       "/api/accounts/address",
	regexp.MustCompile(`^/api/accounts/0[xX][0-9a-fA-F]{40}/txs$`):   "/api/accounts/address/txs",
	regexp.MustCompile(`^/api/miners/0[xX][0-9a-fA-F]{40}$`):         "/api/miners/address",
	regexp.MustCompile(`^/api/miners/0[xX][0-9a-fA-F]{40}/rewards$`): "/api/miners/address/rewards",
	regexp.MustCompile(`^/api/txs/0[xX][0-9a-fA-F]{64}$`):            "/api/txs/txSeq",
}

func validPath(url string) *string {
	validPaths := map[string]bool{
		"/api/miners":           true,
		"/api/rewards":          true,
		"/api/stats/address":    true,
		"/api/stats/fee":        true,
		"/api/stats/layer1-tx":  true,
		"/api/stats/miner":      true,
		"/api/stats/reward":     true,
		"/api/stats/storage":    true,
		"/api/stats/summary":    true,
		"/api/stats/top/data":   true,
		"/api/stats/top/fee":    true,
		"/api/stats/top/files":  true,
		"/api/stats/top/reward": true,
		"/api/stats/top/txs":    true,
		"/api/txs":              true,
	}

	if _, exists := validPaths[url]; exists {
		return &url
	}

	for r, s := range URLPattern {
		if r.MatchString(url) {
			return &s
		}
	}

	return nil
}
