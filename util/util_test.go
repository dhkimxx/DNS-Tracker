package util_test

import (
	"testing"
	"tracker/util"

	"github.com/stretchr/testify/assert"
)

func TestCompareIpAddress(t *testing.T) {
	t.Run("Compare empty slices", func(t *testing.T) {
		existingIps := []string{}
		newIps := []string{}

		isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, newIps)

		assert.True(t, isEqual)
		assert.Zero(t, len(addedIps))
		assert.Zero(t, len(deletedIps))
	})

	t.Run("Added IPs only", func(t *testing.T) {
		existingIps := []string{"192.168.0.1", "192.168.0.2"}
		newIps := []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"}

		isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, newIps)

		assert.False(t, isEqual)
		assert.Equal(t, []string{"192.168.0.3"}, addedIps)
		assert.Zero(t, len(deletedIps))
	})

	t.Run("Deleted IPs only", func(t *testing.T) {
		existingIps := []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"}
		newIps := []string{"192.168.0.1", "192.168.0.2"}

		isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, newIps)

		assert.False(t, isEqual)
		assert.Zero(t, len(addedIps))
		assert.Equal(t, []string{"192.168.0.3"}, deletedIps)
	})

	t.Run("Added and Deleted IPs", func(t *testing.T) {
		existingIps := []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"}
		newIps := []string{"192.168.0.1", "192.168.0.4"}

		isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, newIps)

		assert.False(t, isEqual)
		assert.Equal(t, []string{"192.168.0.4"}, addedIps)
		assert.Equal(t, []string{"192.168.0.2", "192.168.0.3"}, deletedIps)
	})

	t.Run("No change", func(t *testing.T) {
		existingIps := []string{"192.168.0.1", "192.168.0.2"}
		newIps := []string{"192.168.0.1", "192.168.0.2"}

		isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, newIps)

		assert.True(t, isEqual)
		assert.Zero(t, len(addedIps))
		assert.Zero(t, len(deletedIps))
	})

	t.Run("Added duplicate IP", func(t *testing.T) {
		existingIps := []string{"192.168.0.1", "192.168.0.2"}
		newIps := []string{"192.168.0.1", "192.168.0.2", "192.168.0.2"}

		isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, newIps)

		assert.True(t, isEqual)
		assert.Zero(t, len(addedIps))
		assert.Zero(t, len(deletedIps))
	})

	t.Run("Deleted duplicate IP", func(t *testing.T) {
		existingIps := []string{"192.168.0.1", "192.168.0.2", "192.168.0.2"}
		newIps := []string{"192.168.0.1", "192.168.0.2"}

		isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, newIps)

		assert.True(t, isEqual)
		assert.Zero(t, len(addedIps))
		assert.Zero(t, len(deletedIps))
	})

	t.Run("All IPs changed", func(t *testing.T) {
		existingIps := []string{"192.168.0.1", "192.168.0.2"}
		newIps := []string{"192.168.0.3", "192.168.0.4"}

		isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, newIps)

		assert.False(t, isEqual)
		assert.Equal(t, []string{"192.168.0.3", "192.168.0.4"}, addedIps)
		assert.Equal(t, []string{"192.168.0.1", "192.168.0.2"}, deletedIps)
	})
}
