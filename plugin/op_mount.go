package plugin

import (
	"docker-lvm-plugin-ng/lvm"
	"fmt"
	"github.com/docker/go-plugins-helpers/volume"
	"os/exec"
)

func (l *localLvmStoragePlugin) Mount(req *volume.MountRequest) (*volume.MountResponse, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	v, found := l.volumes[req.Name]

	if !found {
		return &volume.MountResponse{}, fmt.Errorf("mount: volume %s not found", req.Name)
	}

	isSnap := func() bool {
		if v, ok := l.volumes[req.Name]; ok {
			return v.Origin != ""
		}
		return false
	}()

	if l.volumes[req.Name].Count == 0 {
		device := lvm.LogicalDevice(v.Vg, req.Name)

		mountArgs := []string{device, getMountpoint(req.Name)}
		if isSnap {
			mountArgs = append([]string{"-o", "nouuid"}, mountArgs...)
		}
		cmd := exec.Command("mount", mountArgs...)
		if out, err := cmd.CombinedOutput(); err != nil {
			return &volume.MountResponse{}, fmt.Errorf("Mount: mount error: %s output %s", err, string(out))
		}
	}
	l.volumes[req.Name].Count++
	if err := saveToDisk(l.volumes); err != nil {
		return &volume.MountResponse{}, err
	}

	return &volume.MountResponse{Mountpoint: getMountpoint(req.Name)}, nil
}
