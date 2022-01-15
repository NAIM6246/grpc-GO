package param

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Int(r *http.Request, p string) int {
	idStr := chi.URLParam(r, p)
	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return 0
		}
		return id
	}
	return 0
}

func UInt(r *http.Request, p string) uint {
	return uint(Int(r, p))
}