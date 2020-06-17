package limiter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLimit(t *testing.T){
	burst := 10
	cl := NewCustomLimiter(burst)
	handler := cl.Limit(http.HandlerFunc(okHandler))

	for i :=0 ; i < burst; i++ {
	    req, err := http.NewRequest("GET", "/health-check", nil)
	    if err != nil {
	    	t.Fatal(err)
	    }
	    w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusTooManyRequests {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusTooManyRequests)
	}
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}