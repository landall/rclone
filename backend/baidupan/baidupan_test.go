// Test BaiduPan filesystem interface
package baidupan_test

import (
	"testing"

	"github.com/rclone/rclone/backend/baidupan"
	"github.com/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestBaiduPan:",
		NilObject:  (*baidupan.Object)(nil),
	})
}
