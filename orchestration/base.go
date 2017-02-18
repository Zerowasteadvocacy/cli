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

func CreateBase() (*DockerBase, error) {
	base := &DockerBase{}
	if err := base.ConfigureContainer(); err != nil {
		return nil, err
	}
	if err = base.ConfigureImage(); err != nil {
		return nil, err
	}
	return base, nil
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

func (base *DockerBase) Pull() error {
	if err := DockerClient.PullImage(base.PullImageOptions, base.AuthConfiguration); err != nil {
		return nil, fmt.Errorf("Error in pulling image: %v", DockerError(err))
	}
	return after, nil
}

func (base *DockerBase) ListImages() (interface{}, error) { //actual signature ([]APIImages, error)
	if images, err := DockerClient.ListImages(base.ListImagesOptions); err != nil {
		return nil, fmt.Errorf("Error in listing images: %v", DockerError(err))
	} else {
		return images, nil
	}
}

func (base *DockerBase) ListContainers() (interface{}, error) { //actual signature ([]APIContainers, error)
	if containers, err := DockerClient.ListContainers(base.ListContainersOptions); err != nil {
		return nil, fmt.Errorf("Error in listing containers: %v", DockerError(err))
	} else {
		return containers, nil
	}
}

func (base *DockerBase) Create() (interface{}, error) { //(*Container, error)
	if container, err := DockerClient.CreateContainer(base.CreateContainerOptions); err != nil {
		return nil, fmt.Errorf("Error in creating container: %v", DockerError(err))
	} else {
		return container, nil
	}
}

func (base *DockerBase) Logs() (interface{}, error) { // error
	if err := DockerClient.Logs(base.LogsOptions); err != nil {
		return nil, fmt.Errorf("Error in getting logs: %v", DockerError(err))
	} else {
		return nil, nil
	}
}

func (base *DockerBase) Download(id string) (interface{}, error) {
	if err := DockerClient.DownloadFromContainer(id, base.DownloadFromContainerOptions); err != nil {
		return nil, fmt.Errorf("Error in downloading from container: %v", DockerError(err))
	} else {
		return nil, nil
	}
}

func (base *DockerBase) Upload(id string) (interface{}, error) {
	if err := DockerClient.UploadToContainer(id, base.UploadToContainerOptions); err != nil {
		return nil, fmt.Errorf("Error in uploading to container: %v", DockerError(err))
	} else {
		return nil, nil
	}
}

func (base *DockerBase) Inspect(id string) (interface{}, error) {
	if container, err := DockerClient.InspectContainer(id); err != nil {
		return nil, fmt.Errorf("Error in inspecting container: %v", DockerError(err))
	} else {
		return container, nil
	}
}

func (base *DockerBase) Stop(id string, timeout uint) (interface{}, error) {
	if err := DockerClient.StopContainer(id, timeout); err != nil {
		return nil, fmt.Errorf("Error in stopping container: %v", DockerError(err))
	} else {
		return nil, nil
	}
}

func (base *DockerBase) Remove() (interface{}, error) {
	if err := DockerClient.RemoveContainer(base.RemoveContainerOptions); err != nil {
		return nil, fmt.Errorf("Error in removing container: %v", DockerError(err))
	} else {
		return nil, nil
	}
}

func (base *DockerBase) Wait(id string) (interface{}, error) {
	if exitCode, err := DockerClient.WaitContainer(id); err != nil {
		return nil, fmt.Errorf("Error in waiting on container: %v", DockerError(err))
	} else {
		return exitCode, nil
	}
}

func (base *DockerBase) Attach() (interface{}, error) {
	if err := DockerClient.AttachToContainer(base.AttachToContainerOptions); err != nil {
		return nil, fmt.Errorf("Error in attaching to container: %v", DockerError(err))
	} else {
		return nil, nil
	}
}

func (base *DockerBase) Rename() (interface{}, error) {
	if err := DockerClient.RenameContainer(base.RenameContainerOptions); err != nil {
		return nil, fmt.Errorf("Error in renaming container: %v", DockerError(err))
	} else {
		return nil, nil
	}
}

func (base *DockerBase) Start(id string) (interface{}, error) {
	if err := DockerClient.StartContainer(id, base.HostConfig); err != nil {
		return nil, fmt.Errorf("Error in starting container: %v", DockerError(err))
	} else {
		return nil, nil
	}
}
