package plugin

import (
	"docker-lvm-plugin-ng/lvm"
	"fmt"
	"log"
	"os/exec"

	"github.com/docker/go-plugins-helpers/volume"
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

	snapshotOrigin, ok := req.Options["snapshot"]
	isSnapshot := ok && snapshotOrigin != ""

	//snapshotOriginVol, ok := l.volumes[snapshotOrigin]
	//isThinSnap := ok && snapshotOriginVol.Thinpool != ""

	if err = provisionMoutpoint(req.Name); err != nil {
		return
	}

	v := &vol{
		Vg:       vgName,
		RefCount: 0,
	}

	//if isSnapshot {
	//	if isThinSnap, _, err = lvm.IsThinlyProvisioned(vgName, snapshotOrigin); err != nil {
	//		return fmt.Errorf("Error creating volume")
	//	}
	//}

	size, ok := req.Options["size"]
	hasSize := ok && size != ""

	//if !hasSize && !isThinSnap {
	//	return fmt.Errorf("Please specify a size with --opt size=")
	//}
	//
	//if hasSize && isThinSnap {
	//	return fmt.Errorf("Please don't specify --opt size= for thin snapshots")
	//}

	if isSnapshot {
		cmdArgs = append(cmdArgs, "--snapshot")
		if hasSize {
			cmdArgs = append(cmdArgs, "--size", size)
		}
		cmdArgs = append(cmdArgs, vgName+"/"+snapshotOrigin)

		v.Origin = snapshotOrigin

	} else if thin, ok := req.Options["thinpool"]; ok && thin != "" {
		v.Thinpool = thin
		cmdArgs = append(cmdArgs,
			"--virtualsize", size,
			"--thin", vgName+"/"+thin,
		)
	} else {
		cmdArgs = append(cmdArgs, "-Zn", "--size", size, vgName)
	}

	log.Printf("lvcreate %+q", cmdArgs)

	cmd := exec.Command("lvcreate", cmdArgs...)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("Create: lvcreate error: %s output %s", err, string(out))
	}

	defer func() {
		if err != nil {
			lvm.RemoveLogicalVolume(vgName, req.Name)
		}
	}()

	if !isSnapshot {
		device := lvm.LogicalDevice(vgName, req.Name)

		mkfsCmdArgs := []string{
			"-f",
			"-m", "bigtime=1",
			device,
		}

		cmd = exec.Command("mkfs.xfs", mkfsCmdArgs...)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("Create: mkfs.xfs error: %s output %s", err, string(out))
		}
	}

	l.volumes[req.Name] = v

	err = saveToDisk(l.volumes)

	return
}
