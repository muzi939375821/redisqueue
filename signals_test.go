package redisqueue

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestNewSignalHandler(t *testing.T) {
	t.Run("closes the returned channel on SIGINT", func(tt *testing.T) {
		ch := newSignalHandler()
		// 获取当前进程的 PID
		pid := os.Getpid()

		// 发送 SIGINT 信号给当前进程
		cmd := exec.Command("kill", "-SIGINT", fmt.Sprintf("%d", pid))
		err := cmd.Run()
		//err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		require.NoError(tt, err)

		select {
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for signal")
		case <-ch:
		}
	})
}
