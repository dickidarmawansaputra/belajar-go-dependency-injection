package test

import (
	"testing"

	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/simple"
	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceSuccess(t *testing.T) {
	// bisa simulasi error, jika di injector ada paramater error
	simpleService, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestSimpleServiceError(t *testing.T) {
	// bisa simulasi error, jika di injector ada paramater error
	simpleService, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}
