package common

import (
	"fmt"
	"github.com/astrix-network/astrixd/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunAstrixdForTesting runs astrixd for testing purposes
func RunAstrixdForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	astrixdRunCommand, err := StartCmd("ASTRIXD",
		"astrixd",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Astrixd started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := astrixdRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Astrixd closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := astrixdRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Astrixd stopped")
	}
}
