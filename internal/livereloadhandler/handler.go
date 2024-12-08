package livereloadhandler

import "net/http"

const defaultPath = "_livereload"

const snippetTemplate = `
<script>
  ;(function() {
    const es = new EventSource('%s')
    es.addEventListener("message", evt => {
      es.close()
      window.location.reload()
    })
  })()
</script>
`

type Config struct {
	Path string // Path is the url path used for the livereload endpoint ("/_livereload")
}

type Handler struct {
	root http.FileSystem
	cnf  *Config
}

func New(root http.FileSystem, cnf *Config) *Handler {
	if cnf.Path == "" {
		cnf.Path = defaultPath
	}
	return &Handler{root, cnf}
}
