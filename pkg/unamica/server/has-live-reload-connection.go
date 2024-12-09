package server

func (s *Server) HasLiveReloadConnection() bool {
	return s.liveReloadConnections > 0
}
