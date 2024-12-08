package webservice

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

func (s *Service) GetLiveReloadSnippet() string {
	return fmt.Sprintf(liveReloadSnippet, fmt.Sprintf("%s", liveReloadPath))
}
