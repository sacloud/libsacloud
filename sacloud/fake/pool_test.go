package fake

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextSubnet(t *testing.T) {

	DataStore = NewInMemoryStore()

	first := pool().nextSubnet(24)
	require.Equal(t, "24.0.1.0", first.networkAddress)
	require.Equal(t, 24, first.networkMaskLen)
	require.Len(t, first.addresses, 251)
	require.Equal(t, "24.0.1.4", first.addresses[0])
	require.Equal(t, "24.0.1.254", first.addresses[len(first.addresses)-1])

	next := pool().nextSubnet(24)
	require.Equal(t, "24.0.2.0", next.networkAddress)
	require.Equal(t, 24, next.networkMaskLen)
	require.Len(t, next.addresses, 251)
	require.Equal(t, "24.0.2.4", next.addresses[0])
	require.Equal(t, "24.0.2.254", next.addresses[len(next.addresses)-1])

}
