// Package httpx is a thin HTTP toolkit over the Go 1.22+ net/http ServeMux:
// a Router with method+pattern routing, per-route and global middleware in
// standard func(http.Handler) http.Handler form, a route registry (for spec
// drift tests and metrics labels), and a request Context that carries the
// handful of helpers the v1 API handlers rely on (path params, JSON
// rendering, body binding).
//
// The package exists so the frozen v1 API could move off the gin framework
// without rewriting ~160 handlers: Context reproduces the exact wire
// behavior the parity goldens pin (content types, marshaling, error
// semantics), while middleware and routing are plain net/http so the v2 API
// can reuse them directly. It has no dependencies outside the standard
// library.
package httpx
