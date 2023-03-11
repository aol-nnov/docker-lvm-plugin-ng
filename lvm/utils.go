package lvm

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

func IsThinlyProvisioned(vg, lv string) (bool, string, error) {
	return lvdisplayGrep(vg, lv, "LV Pool")
}

func CreateLogicalVolume() {

}

func RemoveLogicalVolume(vg string, lv string) ([]byte, error) {
	cmd := exec.Command("lvremove", "--force", fmt.Sprintf("%s/%s", vg, lv))
	if out, err := cmd.CombinedOutput(); err != nil {
		return out, err
	}
	return nil, nil
}

func LogicalDevice(vgName, lvName string) string {
	return fmt.Sprintf("/dev/%s/%s", vgName, lvName)
}

func GetVolumeCreationDateTime(vg, lv string) (time.Time, error) {
	_, creationDateTime, err := lvdisplayGrep(vg, lv, "LV Creation host")
	if err != nil {
		return time.Time{}, err
	}

	// creationDateTime is in the form "LV Creation host, time localhost, 2018-11-18 13:46:08 -0100"
	tokens := strings.Split(creationDateTime, ",")
	date := strings.TrimSpace(tokens[len(tokens)-1])
	return time.Parse("2006-01-02 15:04:05 -0700", date)
}

func lvdisplayGrep(vgName, lvName, keyword string) (bool, string, error) {
	var b2 bytes.Buffer

	cmd1 := exec.Command("lvdisplay", fmt.Sprintf("/dev/%s/%s", vgName, lvName))
	cmd2 := exec.Command("grep", keyword)

	r, w := io.Pipe()
	cmd1.Stdout = w
	cmd2.Stdin = r
	cmd2.Stdout = &b2

	if err := cmd1.Start(); err != nil {
		return false, "", err
	}
	if err := cmd2.Start(); err != nil {
		return false, "", err
	}
	if err := cmd1.Wait(); err != nil {
		return false, "", err
	}
	w.Close()
	if err := cmd2.Wait(); err != nil {
		//exitCode, inErr := system.GetExitCode(err)
		//if inErr != nil {
		//	return false, "", inErr
		//}
		//if exitCode != 1 {
		//	return false, "", err
		//}
		return false, "", err
	}

	if b2.Len() != 0 {
		return true, b2.String(), nil
	}
	return false, "", nil
}
