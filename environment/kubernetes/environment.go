package kubernetes

import (
	"context"
	"os"
	"time"

	"github.com/pterodactyl/wings/environment"
	"github.com/pterodactyl/wings/events"
	"github.com/pterodactyl/wings/remote"
)

type Metadata struct {
	Image string
	Stop  remote.ProcessStopConfiguration
}

type KubernetesEnvironment struct {
}

func New(id string, m *Metadata, c *environment.Configuration) (*KubernetesEnvironment, error) {

}

func (k KubernetesEnvironment) Type() string {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Config() *environment.Configuration {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Events() *events.Bus {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Exists() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) IsRunning(ctx context.Context) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) InSituUpdate() error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) OnBeforeStart(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Stop(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) WaitForStop(ctx context.Context, duration time.Duration, terminate bool) error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Terminate(ctx context.Context, signal os.Signal) error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Destroy() error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) ExitState() (uint32, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Create() error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Attach(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) SendCommand(s string) error {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Readlog(i int) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) State() string {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) SetState(s string) {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) Uptime(ctx context.Context) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (k KubernetesEnvironment) SetLogCallback(f func([]byte)) {
	//TODO implement me
	panic("implement me")
}
