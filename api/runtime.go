package api

import (
	"fmt"
	"strings"

	"github.com/muka/redzilla/docker"
	"github.com/muka/redzilla/model"
	"github.com/sirupsen/logrus"
)

//Reset reset container runtime information
func (i *Instance) Reset() error {

	i.instance.IP = ""
	i.instance.Status = model.InstanceStopped

	return nil
}

//GetIP return the container IP
func (i *Instance) GetIP() (string, error) {

	ip := ""

	if len(i.instance.IP) > 0 {
		return i.instance.IP, nil
	}

	net, err := docker.GetNetwork(i.cfg.Network)
	if err != nil {
		return "", err
	}

	if len(net.Containers) == 0 {
		return "", fmt.Errorf("Network '%s' has no container attached", i.cfg.Network)
	}

	for _, container := range net.Containers {
		if container.Name == i.instance.Name {
			ip = container.IPv4Address[:strings.Index(container.IPv4Address, "/")]
			logrus.Debugf("Container IP %s", ip)
			break
		}
	}

	if ip == "" {
		return ip, fmt.Errorf("IP not found for container `%s`", i.instance.Name)
	}

	i.instance.IP = ip

	return ip, nil
}
