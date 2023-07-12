package test

import (
	"testing"

	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/simple"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
