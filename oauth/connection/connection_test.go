package connection_test

import (
	"github.com/ory-am/hydra/Godeps/_workspace/src/github.com/pborman/uuid"
	"github.com/ory-am/hydra/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	. "github.com/ory-am/hydra/oauth/connection"
	"testing"
)

func TestConnection(t *testing.T) {
	c := &DefaultConnection{
		ID:            uuid.New(),
		LocalSubject:  "peter",
		RemoteSubject: "peter@gmail.com",
		Provider:      "google",
	}

	assert.Equal(t, c.ID, c.GetID())
	assert.Equal(t, c.Provider, c.GetProvider())
	assert.Equal(t, c.LocalSubject, c.GetLocalSubject())
	assert.Equal(t, c.RemoteSubject, c.GetRemoteSubject())
}
