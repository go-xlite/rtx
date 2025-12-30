package pid

import (
	"fmt"

	pid_s "github.com/go-xlite/rtx/svc/pid"
)

type PidHandler struct {
	Pid *pid_s.Pid
}

func NewPidHandler() *PidHandler {
	h := &PidHandler{}
	h.Pid = pid_s.NewPid()
	return h
}

// strategy: exit if another instance is running
func (h *PidHandler) Handle_ExitOnDuplicate() error {

	if h.Pid.DoesPidFileExist() {
		existingPid := pid_s.NewPid()
		existingPid.PidFilePath = h.Pid.PidFilePath
		pidFromFile, err := existingPid.ReadPidFromFile()
		if err != nil {
			return err
		}
		if existingPid.IsProcessRunning(pidFromFile) {
			return fmt.Errorf("service already running")
		} else {
			// Stale PID file, remove it
			existingPid.RemovePidFile()
		}
	}
	h.Pid.WritePidFile()
	return nil
}

func (h *PidHandler) Handle_CleanupOnExit() error {
	return h.Pid.RemovePidFile()
}

// strategy: kill existing process and start new one
func (h *PidHandler) Handle_RestartOnDuplicate() error {

	if h.Pid.DoesPidFileExist() {
		existingPid := pid_s.NewPid()
		existingPid.PidFilePath = h.Pid.PidFilePath
		pidFromFile, err := existingPid.ReadPidFromFile()
		if err != nil {
			return err
		}
		if existingPid.IsProcessRunning(pidFromFile) {
			err := existingPid.KillProcess(pidFromFile)
			if err != nil {
				return err
			}
			existingPid.RemovePidFile()
		} else {
			// Stale PID file, remove it
			existingPid.RemovePidFile()
		}
	}
	h.Pid.WritePidFile()
	return nil
}
