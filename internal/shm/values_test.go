package shm

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestValues_read(t *testing.T) {
	// given
	var values Values

	file, err := os.Open("testdata/data.bin")
	require.NoError(t, err)

	// when
	err = values.Read(file)

	// then
	assert.NoError(t, err)

	assert.True(t, values.SdkActive)
	assert.True(t, values.Paused)
	assert.Equal(t, uint64(5166460), values.Time)
	assert.Equal(t, uint64(252773222), values.SimulatedTime)
	assert.Equal(t, uint64(252745580), values.RenderTime)
	assert.Equal(t, int64(0), values.MultiplayerTimeOffset)

	assert.Equal(t, uint32(12), values.ScsValues.TelemetryPluginRevision)
	assert.Equal(t, uint32(1), values.ScsValues.GameVersionMajor)
	assert.Equal(t, uint32(18), values.ScsValues.GameVersionMinor)
	assert.Equal(t, uint32(1), values.ScsValues.Game)
	assert.Equal(t, uint32(1), values.ScsValues.GameTelemetryVersionMajor)
	assert.Equal(t, uint32(18), values.ScsValues.GameTelemetryVersionMinor)
}
