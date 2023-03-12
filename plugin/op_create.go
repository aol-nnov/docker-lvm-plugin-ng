package plugin

import (
	"docker-lvm-plugin-ng/lvm"
	"fmt"
	"github.com/docker/go-plugins-helpers/volume"
	"log"
	"os/exec"
)

func (l *localLvmStoragePlugin) Create(req *volume.CreateRequest) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if _, exists := l.volumes[req.Name]; exists {
		return nil
	}

	vgName, err := getVgName(req.Options["vg"])
	if err != nil {
		return
	}

	cmdArgs := []string{"-y", "-n", req.Name, "--setactivationskip", "n"}
	snap, ok := req.Options["snapshot"]
	isSnapshot := ok && snap != ""
	isThinSnap := false

	moutpoint, err := provisionMoutpoint(req.Name)

	if err != nil {
		return
	}


	v := &vol{
		Name:       req.Name,
		MountPoint: moutpoint,
		Vg:         vgName,
		Count:      0,
	}

	if isSnapshot {
		if isThinSnap, _, err = lvm.IsThinlyProvisioned(vgName, snap); err != nil {
			return fmt.Errorf("Error creating volume")
		}
	}
	size, ok := req.Options["size"]
	hasSize := ok && size != ""

	if !hasSize && !isThinSnap {
		return fmt.Errorf("Please specify a size with --opt size=")
	}

	if hasSize && isThinSnap {
		return fmt.Errorf("Please don't specify --opt size= for thin snapshots")
	}

	if isSnapshot {
		cmdArgs = append(cmdArgs, "--snapshot")
		if hasSize {
			cmdArgs = append(cmdArgs, "--size", size)
		}
		cmdArgs = append(cmdArgs, vgName+"/"+snap)

		v.Origin = snap

	} else if thin, ok := req.Options["thinpool"]; ok && thin != "" {
		v.Thinpool = thin
		cmdArgs = append(cmdArgs,
			"--virtualsize", size,
			"--thin", vgName+"/"+thin,
		)
	} else {
		cmdArgs = append(cmdArgs, "--size", size, vgName)
	}

	log.Printf("lvcreate %+q", cmdArgs)

	cmd := exec.Command("lvcreate", cmdArgs...)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("Create: lvcreate error: %s output %s", err, string(out))
	}

	defer func() {
		if err != nil {
			lvm.RemoveLogicalVolume(req.Name, vgName)
		}
	}()

	if !isSnapshot {
		device := lvm.LogicalDevice(vgName, req.Name)

		cmd = exec.Command("mkfs.xfs", device)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("Create: mkfs.xfs error: %s output %s", err, string(out))
		}
	}

	l.volumes[v.Name] = v

	err = saveToDisk(l.volumes)

	return
}
