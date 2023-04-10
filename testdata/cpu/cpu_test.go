package cpu

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestBpfBindingsSize(t *testing.T) {
	require.Equal(t, int(unsafe.Sizeof(UnwinderConfigT{})), 2)
	require.Equal(t, int(unsafe.Sizeof(UnwinderStatsT{})), 80)
	require.Equal(t, int(unsafe.Sizeof(ChunkInfoT{})), 40)
	require.Equal(t, int(unsafe.Sizeof(UnwindInfoChunksT{})), 1200)
	require.Equal(t, int(unsafe.Sizeof(StackTraceT{})), 1024)
	require.Equal(t, int(unsafe.Sizeof(StackCountKeyT{})), 20)
	require.Equal(t, int(unsafe.Sizeof(MappingT{})), 40)
	require.Equal(t, int(unsafe.Sizeof(ProcessInfoT{})), 10016)
	require.Equal(t, int(unsafe.Sizeof(UnwindStateT{})), 1056)
	require.Equal(t, int(unsafe.Sizeof(StackUnwindRowT{})), 16, "failed unpacked struct")
	require.Equal(t, int(unsafe.Sizeof(StackUnwindRowT{})), 14, "failed packed struct")
	require.Equal(t, int(unsafe.Sizeof(StackUnwindTableT{})), 3500000)
}
