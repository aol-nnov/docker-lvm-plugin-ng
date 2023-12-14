package lvm

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func GetVolumeInfo(args ...string) ([]Volume, error) {
	if len(args) > 2 {
		return nil, fmt.Errorf("lvm.GetInfo: Two args max. Vg name, lv name")
	}

	cmdArgs := []string{
		"--reportformat", "json",
		"-o", "+lv_time,lv_host,lv_full_name,lv_path",
	}

	if devPath := strings.Join(args, "/"); devPath != "" {
		cmdArgs = append(cmdArgs, devPath)
	}

	var parsedResponse LvmResponse
	cmd := exec.Command("lvs", cmdArgs...)
	pipe, err := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	if err == nil {
		err = json.NewDecoder(pipe).Decode(&parsedResponse)
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	if len(parsedResponse.Reports[0].Lvs) < 1 {
		return nil, fmt.Errorf("GetVolumeInfo: Results parsing failed or volume not found. Error: %v", err)
	}

	return parsedResponse.Reports[0].Lvs, err
}
