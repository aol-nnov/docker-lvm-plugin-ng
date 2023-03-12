package plugin

import (
	"docker-lvm-plugin-ng/lvm"
	"fmt"
	"github.com/docker/go-plugins-helpers/volume"
	"github.com/moby/sys/mountinfo"
	"os/exec"
)

func (l *localLvmStoragePlugin) Mount(req *volume.MountRequest) (*volume.MountResponse, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	vol, found := l.volumes[req.Name]
	if !found {
		return &volume.MountResponse{}, fmt.Errorf("mount: volume %s not found", req.Name)
	}

	mountpoint := getMountpoint(req.Name)
	isVolMounted, err := mountinfo.Mounted(mountpoint)
	if err != nil {
		return nil, fmt.Errorf("error mounting volume '%s': %s", req.Name, err)
	}

	if !isVolMounted && l.volumes[req.Name].RefCount > 0 {
		return nil, fmt.Errorf("Volume '%s' not mounted and RefCount is %d! Something nasty", req.Name, l)
	}

	if vol.RefCount == 0 {
		device := lvm.LogicalDevice(vol.Vg, req.Name)

		mountArgs := []string{device, mountpoint}
		if vol.Origin != "" { // snapshot
			mountArgs = append([]string{"-o", "nouuid"}, mountArgs...)
		}
		cmd := exec.Command("mount", mountArgs...)
		if out, err := cmd.CombinedOutput(); err != nil {
			return &volume.MountResponse{}, fmt.Errorf("Mount: mount error: %s output %s", err, string(out))
		}
	}
	l.volumes[req.Name].RefCount++
	//if err := saveToDisk(l.volumes); err != nil {
	//	return &volume.MountResponse{}, err
	//}

	return &volume.MountResponse{Mountpoint: mountpoint}, nil
}
