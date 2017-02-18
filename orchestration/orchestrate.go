package orchestration

import (
	"fmt"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

var DockerClient *docker.Client

type dockerFunc func(...interface{}) (interface{}, error)

type ContainerMaintainer interface {
	ConfigureContainer()
	Create()
	Start()
	Stop()
	Remove()
	Rename()
	Exec()
	ListContainers()
	Inspect()
	Upload()
	Download()
	Wait()
	Attach()
}

type ImageMaintainer interface {
	ConfigureImage()
	Pull()
	ListImages()
}

type DockerBase struct {
	docker.PullImageOptions
	docker.ListImagesOptions
	docker.ListContainersOptions
	docker.CreateContainerOptions
	docker.LogsOptions
	docker.DownloadFromContainerOptions
	docker.UploadToContainerOptions
	docker.AttachToContainerOptions
	*docker.HostConfig
	docker.RemoveContainerOptions
	docker.RenameContainerOptions
	docker.AuthConfiguration
}

func (base *DockerBase) ConfigureContainer() error {
	//Default container configuration
	base.ListContainersOptions = docker.ListContainersOptions{All: true}
	base.HostConfig = &docker.HostConfig{
		ReadonlyRootfs: false,
		RestartPolicy:  docker.NeverRestart(),
	}
	base.CreateContainerOptions = docker.CreateContainerOptions{
		Config: &docker.Config{
			AttachStderr:    false,
			AttachStdin:     false,
			AttachStdout:    false,
			Tty:             false,
			OpenStdin:       false,
			NetworkDisabled: false,
		},
		HostConfig: base.HostConfig,
	}
	base.LogsOptions = docker.LogsOptions{
		OutputStream: os.Stdout,
		ErrorStream:  os.Stderr,
		Follow:       true,
		Since:        0,
		Timestamps:   false,
		Tail:         true,
	}
	base.AttachToContainerOptions = docker.AttachToContainerOptions{
		OutputStream: os.Stdout,
		ErrorStream:  os.Stderr,
		Logs:         false,
		Stream:       true,
		Stdout:       true,
		Stderr:       true,
		RawTerminal:  true,
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	base.DownloadFromContainerOptions = docker.DownloadFromContainerOptions{
		OutputStream: os.Stdout,
		Path:         pwd,
	}
	base.UploadToContainerOptions = docker.UploadToContainerOptions{
		Path:                 pwd,
		NoOverwriteDirNonDir: true,
	}
	return nil
}

func (base *DockerBase) ConfigureImage() {
	base.PullImageOptions = docker.PullImageOptions{
		RawJSONStream: true,
		OutputStream:  os.Stdout,
	}
}

func DockerError(err error) error {
	if _, ok := err.(*docker.Error); ok {
		return fmt.Errorf("Docker: %v", err.(*docker.Error).Message)
	}
	return err
}

func (base *DockerBase) PullImage() error {
	if err := DockerClient.PullImage(opts, auth); err != nil {
		return nil, fmt.Errorf("Error in pulling image: %v", DockerError(err))
	}
	return after, nil
}

func (base *DockerBase) ListImages() (interface{}, error) { //actual signature ([]APIImages, error)

	return nil, nil
}

func (base *DockerBase) ListContainers() (interface{}, error) { //actual signature ([]APIContainers, error)

}

func (base *DockerBase) CreateContainer() (interface{}, error) { //(*Container, error)

}

func (base *DockerBase) Logs() (interface{}, error) { // error

}

func (base *DockerBase) DownloadFromContainer(id string) error {

}

func (base *DockerBase) UploadToContainer(id string) error {

}

func (base *DockerBase) InspectContainer(id string) (*Container, error) {

}

func (base *DockerBase) StopContainer(id string, timeout uint) error {

}

func (base *DockerBase) RemoveContainer() error {

}

func (base *DockerBase) WaitContainer(id string) (int, error) {

}

func (base *DockerBase) AttachToContainer() error {

}

func (base *DockerBase) RenameContainer() error {

}

func (base *DockerBase) StartContainer(id string, hostConfig *HostConfig) error {

}
