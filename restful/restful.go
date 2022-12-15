package restful

import "net/http"

var (
	ModeDebug   = "debug"
	ModeRelease = "release"
	Mode        = ModeDebug
)

// GoRestful struct
type GORestful struct {
	*http.ServeMux
}

// New GoRestful
func New() *GORestful {
	//debugPrintWARNINGNew()
	return &GORestful{
		ServeMux: http.NewServeMux(),
	}
}

// HandlerServeHTTP
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
func (r *GORestful) HandlerServeHTTP(method, pattern string, handler http.HandlerFunc) {
	//debugPrintMessage(method, pattern, handler)
	r.Handle(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(405)
			_, _ = w.Write([]byte("Method Not Allowed"))
			return
		}
		handler.ServeHTTP(w, r)
	}))
}

// GET is a shortcut for GORestful.HandlerServeHTTP(http.MethodGet, path, handle)
func (r *GORestful) GET(pattern string, handler http.HandlerFunc) {
	r.HandlerServeHTTP(http.MethodGet, pattern, handler)
}

// GET is a shortcut for GORestful.HandlerServeHTTP(http.MethodHead, path, handle)
func (r *GORestful) HEAD(pattern string, handler http.HandlerFunc) {
	r.HandlerServeHTTP(http.MethodHead, pattern, handler)
}

// GET is a shortcut for GORestful.HandlerServeHTTP(http.MethodPost, path, handle)
func (r *GORestful) POST(pattern string, handler http.HandlerFunc) {
	r.HandlerServeHTTP(http.MethodPost, pattern, handler)
}

// GET is a shortcut for GORestful.HandlerServeHTTP(http.MethodPut, path, handle)
func (r *GORestful) PUT(pattern string, handler http.HandlerFunc) {
	r.HandlerServeHTTP(http.MethodPut, pattern, handler)
}

// GET is a shortcut for GORestful.HandlerServeHTTP(http.MethodPatch, path, handle)
func (r *GORestful) PATCH(pattern string, handler http.HandlerFunc) {
	r.HandlerServeHTTP(http.MethodPatch, pattern, handler)
}

// GET is a shortcut for GORestful.HandlerServeHTTP(http.MethodDelete, path, handle)
func (r *GORestful) DELETE(pattern string, handler http.HandlerFunc) {
	r.HandlerServeHTTP(http.MethodDelete, pattern, handler)
}

// GET is a shortcut for GORestful.HandlerServeHTTP(http.MethodConnect, path, handle)
func (r *GORestful) CONNECT(pattern string, handler http.HandlerFunc) {
	r.HandlerServeHTTP(http.MethodConnect, pattern, handler)
}

// GET is a shortcut for GORestful.HandlerServeHTTP(http.MethodTrace, path, handle)
func (r *GORestful) TRACE(pattern string, handler http.HandlerFunc) {
	r.HandlerServeHTTP(http.MethodTrace, pattern, handler)
}
