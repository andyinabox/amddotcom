package watchservice

import (
	"context"
)

func (s *Service) Watch(ctx context.Context, paths []string) (<-chan string, <-chan error) {
	changes := make(chan string)
	errs := make(chan error)

	return changes, errs
}
