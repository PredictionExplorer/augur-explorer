package httpx

import (
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
)

func serve(r *Router, method, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(method, path, nil))
	return w
}

func TestRouterRoutesByMethodAndPattern(t *testing.T) {
	r := NewRouter()
	r.GET("/api/things/{id}", func(c *Context) {
		c.JSON(http.StatusOK, H{"id": c.Param("id")})
	})
	r.POST("/api/things", func(c *Context) {
		c.Status(http.StatusCreated)
	})

	if w := serve(r, http.MethodGet, "/api/things/42"); w.Code != http.StatusOK || w.Body.String() != `{"id":"42"}` {
		t.Errorf("GET: %d %q", w.Code, w.Body.String())
	}
	if w := serve(r, http.MethodPost, "/api/things"); w.Code != http.StatusCreated {
		t.Errorf("POST: %d", w.Code)
	}
}

func TestRouterNotFound(t *testing.T) {
	r := NewRouter()
	r.GET("/known", func(c *Context) { c.Status(http.StatusOK) })

	w := serve(r, http.MethodGet, "/unknown")
	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
	if body := w.Body.String(); body != "404 page not found\n" {
		t.Errorf("body = %q (stdlib NotFound text expected)", body)
	}
}

func TestRouterMethodNotAllowed(t *testing.T) {
	r := NewRouter()
	r.GET("/read-only", func(c *Context) { c.Status(http.StatusOK) })

	w := serve(r, http.MethodPost, "/read-only")
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want 405", w.Code)
	}
	if allow := w.Header().Get("Allow"); !strings.Contains(allow, http.MethodGet) {
		t.Errorf("Allow header = %q, want it to list GET", allow)
	}
}

// TestRouterHeadServedByGetRoute pins the ServeMux behavior that a GET
// pattern answers HEAD requests; the real server (not httptest recorders)
// discards the body. Legacy gin answered 404 — the new behavior is standard
// and strictly more correct for CDNs and health checks.
func TestRouterHeadServedByGetRoute(t *testing.T) {
	r := NewRouter()
	r.GET("/doc", func(c *Context) { c.String(http.StatusOK, "body") })

	srv := httptest.NewServer(r)
	defer srv.Close()

	resp, err := http.Head(srv.URL + "/doc")
	if err != nil {
		t.Fatalf("HEAD request: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("HEAD status = %d, want 200", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if len(body) != 0 {
		t.Errorf("HEAD body = %q, want empty", body)
	}
}

func TestRouterTrailingSlashRedirect(t *testing.T) {
	r := NewRouter()
	r.GET("/api/list", func(c *Context) { c.Status(http.StatusOK) })
	r.POST("/api/submit", func(c *Context) { c.Status(http.StatusOK) })

	t.Run("get_301_and_query_preserved", func(t *testing.T) {
		w := serve(r, http.MethodGet, "/api/list/?limit=5")
		if w.Code != http.StatusMovedPermanently {
			t.Fatalf("status = %d, want 301", w.Code)
		}
		if loc := w.Header().Get("Location"); loc != "/api/list?limit=5" {
			t.Errorf("Location = %q", loc)
		}
	})

	t.Run("post_307", func(t *testing.T) {
		w := serve(r, http.MethodPost, "/api/submit/")
		if w.Code != http.StatusTemporaryRedirect {
			t.Fatalf("status = %d, want 307", w.Code)
		}
		if loc := w.Header().Get("Location"); loc != "/api/submit" {
			t.Errorf("Location = %q", loc)
		}
	})

	t.Run("no_alternative_stays_404", func(t *testing.T) {
		if w := serve(r, http.MethodGet, "/api/other/"); w.Code != http.StatusNotFound {
			t.Errorf("status = %d, want 404", w.Code)
		}
	})

	t.Run("wrong_method_not_redirected", func(t *testing.T) {
		// /api/list/ with POST: the GET-only alternative must not trigger a
		// redirect loophole; the mux answers 405 for the closest match.
		w := serve(r, http.MethodPost, "/api/list/")
		if w.Code == http.StatusTemporaryRedirect || w.Code == http.StatusMovedPermanently {
			t.Errorf("status = %d, redirect must not apply across methods", w.Code)
		}
	})

	t.Run("scheme_relative_location_never_emitted", func(t *testing.T) {
		// A "//host/path" Location is scheme-relative: browsers would leave
		// the site. The trailing-slash logic refuses such paths; the mux's
		// own clean-path redirect (which may still answer) always targets
		// the cleaned, single-slash path — assert no "//" ever escapes.
		req := httptest.NewRequest(http.MethodGet, "/api/list/", nil)
		req.URL.Path = "//api/list/"
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if loc := w.Header().Get("Location"); strings.HasPrefix(loc, "//") {
			t.Fatalf("scheme-relative Location %q escaped (status %d)", loc, w.Code)
		}
	})
}

func TestRouterMiddlewareOrderAndScope(t *testing.T) {
	var order []string
	mark := func(name string) Middleware {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				order = append(order, name)
				next.ServeHTTP(w, r)
			})
		}
	}

	r := NewRouter()
	r.Use(mark("global1"), mark("global2"))
	r.GET("/route", func(c *Context) { order = append(order, "handler") }, mark("route1"), mark("route2"))

	serve(r, http.MethodGet, "/route")
	want := []string{"global1", "global2", "route1", "route2", "handler"}
	if !slices.Equal(order, want) {
		t.Errorf("execution order = %v, want %v", order, want)
	}
}

func TestRouterGlobalMiddlewareCovers404(t *testing.T) {
	r := NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("X-Traced", "1")
			next.ServeHTTP(w, req)
		})
	})
	r.GET("/exists", func(c *Context) { c.Status(http.StatusOK) })

	w := serve(r, http.MethodGet, "/does-not-exist")
	if w.Code != http.StatusNotFound || w.Header().Get("X-Traced") != "1" {
		t.Errorf("global middleware must wrap unmatched requests: %d, X-Traced=%q",
			w.Code, w.Header().Get("X-Traced"))
	}
}

func TestRouterPatternSeenByOuterMiddleware(t *testing.T) {
	var pattern string
	r := NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(w, req)
			pattern = PatternPath(req) // metrics-style: read after dispatch
		})
	})
	r.GET("/api/x/{id}", func(c *Context) { c.Status(http.StatusOK) })

	serve(r, http.MethodGet, "/api/x/9")
	if pattern != "/api/x/{id}" {
		t.Errorf("outer middleware saw pattern %q, want /api/x/{id}", pattern)
	}

	pattern = "sentinel"
	serve(r, http.MethodGet, "/nope")
	if pattern != "" {
		t.Errorf("unmatched request must leave pattern empty, got %q", pattern)
	}
}

func TestRouterLiteralSegmentBeatsParam(t *testing.T) {
	// Regression guard for the real route table:
	// /donations/nft/claims/{offset}/{limit} coexists with
	// /donations/nft/claims/by_user/{user_addr}; the literal wins.
	r := NewRouter()
	r.GET("/claims/{offset}/{limit}", func(c *Context) { c.String(http.StatusOK, "paged") })
	r.GET("/claims/by_user/{user_addr}", func(c *Context) { c.String(http.StatusOK, "by_user") })

	if w := serve(r, http.MethodGet, "/claims/by_user/0xabc"); w.Body.String() != "by_user" {
		t.Errorf("literal segment must win: got %q", w.Body.String())
	}
	if w := serve(r, http.MethodGet, "/claims/0/10"); w.Body.String() != "paged" {
		t.Errorf("param route must still match: got %q", w.Body.String())
	}
}

func TestRouterConflictingPatternsPanicAtRegistration(t *testing.T) {
	r := NewRouter()
	r.GET("/a/{x}", func(c *Context) {})
	defer func() {
		if recover() == nil {
			t.Fatal("registering an ambiguous pattern must panic at startup")
		}
	}()
	r.GET("/a/{y}", func(c *Context) {})
}

func TestRouterFreezesAfterFirstRequest(t *testing.T) {
	r := NewRouter()
	r.GET("/x", func(c *Context) { c.Status(http.StatusOK) })
	serve(r, http.MethodGet, "/x")

	assertPanics := func(name string, fn func()) {
		t.Helper()
		defer func() {
			if recover() == nil {
				t.Errorf("%s after serving must panic", name)
			}
		}()
		fn()
	}
	assertPanics("Use", func() { r.Use(func(h http.Handler) http.Handler { return h }) })
	assertPanics("Handle", func() { r.GET("/y", func(c *Context) {}) })
}

func TestRouterRoutesRegistry(t *testing.T) {
	r := NewRouter()
	r.GET("/a", func(c *Context) {})
	r.POST("/b/{id}", func(c *Context) {})
	r.HEAD("/c", func(c *Context) {})

	got := r.Routes()
	want := []Route{
		{Method: "GET", Pattern: "/a"},
		{Method: "POST", Pattern: "/b/{id}"},
		{Method: "HEAD", Pattern: "/c"},
	}
	if !slices.Equal(got, want) {
		t.Errorf("Routes() = %v, want %v", got, want)
	}

	// The registry is a copy: mutating it must not corrupt the router.
	got[0].Pattern = "/mutated"
	if r.Routes()[0].Pattern != "/a" {
		t.Error("Routes() must return a copy")
	}
}
