package presenters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	SetupMocks()
}

func TestAddress_GetAddress(t *testing.T) {
	address := AddressPresenter.GetAddress(nil, 0)
	assert.NotNil(t, address)
	assert.Equal(t, 0, address.ID)
	t.Logf("Address: %v", address)
}
