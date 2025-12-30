package pid

import (
	"fmt"
	"os"
	"syscall"
)

func NewPid() *Pid {
	p := &Pid{}
	p.Pid = os.Getpid()
	return p
}

type Pid struct {
	Pid         int
	PidFilePath string
}

func (p *Pid) IsProcessRunning(processId int) bool {
	process, err := os.FindProcess(processId)
	if err != nil {
		return false
	}
	// Sending signal 0 to a process is a way to check for its existence
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func (p *Pid) DoesPidFileExist() bool {
	if err := p.ensure_filePath(); err != nil {
		panic(err)
	}
	_, err := os.Stat(p.PidFilePath)
	return err == nil
}

func (p *Pid) KillProcess(processId int) error {
	process, err := os.FindProcess(processId)
	if err != nil {
		return err
	}
	return process.Kill()
}

func (p *Pid) ReadPidFromFile() (int, error) {

	data, err := os.ReadFile(p.PidFilePath)
	if err != nil {
		return 0, err
	}

	var pid int
	_, err = fmt.Sscanf(string(data), "%d", &pid)
	if err != nil {
		return 0, err
	}

	return pid, nil
}

func (p *Pid) RemovePidFile() error {
	if err := p.ensure_filePath(); err != nil {
		return err
	}
	return os.Remove(p.PidFilePath)
}

func (p *Pid) ensure_filePath() error {
	if p.PidFilePath == "" {
		return fmt.Errorf("PID file path is not set")
	}
	return nil
}

func (p *Pid) WritePidFile() error {
	if err := p.ensure_filePath(); err != nil {
		return err
	}

	err := os.WriteFile(p.PidFilePath, []byte(fmt.Sprintf("%d", p.Pid)), 0644)
	if err != nil {
		return err
	}

	return nil
}
