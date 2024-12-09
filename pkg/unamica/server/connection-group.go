package server

import (
	"context"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

// connectionGroup maintains a pool of string channels
// so that a string message can be sent to all open connections.
// it will also automatically handle cleanup when the context expires
type connectionGroup struct {
	connections map[uuid.UUID]chan string
	m           sync.Mutex
}

func newConnectionGroup() *connectionGroup {
	return &connectionGroup{
		connections: make(map[uuid.UUID]chan string),
	}
}

func (c *connectionGroup) Add(ctx context.Context) (uuid.UUID, <-chan string) {
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

func (c *connectionGroup) Remove(key uuid.UUID) {
	c.m.Lock()
	close(c.connections[key])
	delete(c.connections, key)
	c.m.Unlock()
}

func (c *connectionGroup) Send(msg string) {
	log.Debug("send message to all channels", "message", msg)
	c.m.Lock()
	for _, ch := range c.connections {
		ch <- msg
	}
	c.m.Unlock()
}

func (c *connectionGroup) Count() int {
	return len(c.connections)
}
