// +build go1.8

package doggy

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hnlq715/doggy/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// ListenAndServeGracefully always returns a non-nil error.
func ListenAndServeGracefully(handler http.Handler) error {
	l := utils.LogFromContext(context.Background())

	h := &http.Server{Addr: viper.GetString("listen"), Handler: handler}
	go func() {
		if err := h.ListenAndServe(); err != nil {
			l.Error("h.ListenAndServe failed", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	l.Info("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	h.Shutdown(ctx)

	l.Info("Server gracefully stopped")
	return nil
}
