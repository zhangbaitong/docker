package daemon

import (
	"github.com/docker/docker/engine"
)

func (daemon *Daemon) ContainerStop(job *engine.Job) engine.Status {
	if len(job.Args) != 1 {
		return job.Errorf("Usage: %s CONTAINER\n", job.Name)
	}
	var (
		name = job.Args[0]
		t    = 10
	)
	if job.EnvExists("t") {
		t = job.GetenvInt("t")
	}
	container, err := daemon.Get(name)
	if err != nil {
		return job.Error(err)
	}
	if !container.IsRunning() {
		return job.Errorf("Container already stopped")
	}
	if err := container.Stop(int(t)); err != nil {
		return job.Errorf("Cannot stop container %s: %s\n", name, err)
	}
	container.LogEvent("stop")
	return engine.StatusOK
}
