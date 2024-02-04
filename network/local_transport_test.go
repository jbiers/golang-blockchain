package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddr(t *testing.T) {
	trA := NewLocalTransport("A")

	addr := trA.Addr()
	expectedAddr := NetAddr("A")

	assert.Equal(t, addr, expectedAddr)
}

func TestConnect(t *testing.T) {
	trA := NewLocalTransport("A")
	trB := NewLocalTransport("B")

	trA.Connect(trB)
	trB.Connect(trA)

	assert.Equal(t, trA.peers[trB.Addr()], trB)
	assert.Equal(t, trB.peers[trA.Addr()], trA)
}

func TestSendMessage(t *testing.T) {
	trA := NewLocalTransport("A")
	trB := NewLocalTransport("B")

	trA.Connect(trB)
	trB.Connect(trA)

	msg := []byte("Hello World")

	assert.Nil(t, trA.SendMessage(trB.Addr(), msg))

	rpc := <-trB.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, trA.Addr())
}
