package server

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/m-mizutani/httpdump/pkg/utils"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (x *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Logger().Error("Failed to read body", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Logger().Info("HTTP Access",
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("query", r.URL.RawQuery),
		slog.String("remoteAddr", r.RemoteAddr),
		slog.String("userAgent", r.UserAgent()),
		slog.String("referer", r.Referer()),
		slog.Any("headers", r.Header),
		slog.String("body", string(body)),
	)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World!"))
}
