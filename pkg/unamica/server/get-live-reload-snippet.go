package server

import "fmt"

const liveReloadSnippet = `
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

func (s *Server) GetLiveReloadSnippet() string {
	url := fmt.Sprintf("http://%s%s", s.srv.Addr, s.cnf.LiveReloadPath)
	return fmt.Sprintf(liveReloadSnippet, url)
}
