package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"time"

	"github.com/maliur/cardmarket/pkg/listing"
)

func loggingMiddleware(l hclog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Debug("Incoming", "method", r.Method, "path", r.URL.Path)
			next.ServeHTTP(w, r)
			l.Debug("Took", "ms", time.Since(start), "method", r.Method, "path", r.URL.Path)
		}

		return http.HandlerFunc(fn)
	}
}

func NewRouter(l hclog.Logger, ls listing.Service) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/order/paid", getPaidOrders(ls))
	r.HandleFunc("/order/sent", getSentOrders(ls))

	if l.IsDebug() {
		r.Use(loggingMiddleware(l))
	}

	return r
}

// Orders
// TODO: Extract this into a separate file
func getPaidOrders(ls listing.Service) http.HandlerFunc {
	return func (w http.ResponseWriter, _ *http.Request) {
		orders, err := ls.GetPaidOrders()
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(orders); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}
}

func getSentOrders(ls listing.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		orders, err := ls.GetSentOrders()
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(orders); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}
}
