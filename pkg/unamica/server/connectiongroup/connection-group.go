package connectiongroup

import (
	"context"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

type ConnectionGroup struct {
	connections map[uuid.UUID]chan string
	m           sync.Mutex
}

func New() *ConnectionGroup {
	return &ConnectionGroup{
		connections: make(map[uuid.UUID]chan string),
	}
}

func (c *ConnectionGroup) Add(ctx context.Context) (uuid.UUID, <-chan string) {
	key := uuid.New()
	ch := make(chan string)

	c.m.Lock()
	c.connections[key] = ch
	c.m.Unlock()

	// automatically remove when context is closed
	go func() {
		<-ctx.Done()
		c.Remove(key)
	}()

	return key, ch
}

func (c *ConnectionGroup) Remove(key uuid.UUID) {
	c.m.Lock()
	close(c.connections[key])
	delete(c.connections, key)
	c.m.Unlock()
}

func (c *ConnectionGroup) Send(msg string) {
	log.Debug("send message to all channels", "message", msg)
	c.m.Lock()
	for _, ch := range c.connections {
		ch <- msg
	}
	c.m.Unlock()
}

func (c *ConnectionGroup) Count() int {
	return len(c.connections)
}
