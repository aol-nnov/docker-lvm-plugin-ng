package lvm

import (
	"fmt"
	"log"
	"os/exec"
)

//func IsThinlyProvisioned(vg, lv string) (bool, string, error) {
//	return lvdisplayGrep(vg, lv, "LV Pool")
//}

func CreateLogicalVolume() {

}

func RemoveLogicalVolume(vg string, lv string) error {
	cmd := exec.Command("lvremove", "--force", fmt.Sprintf("%s/%s", vg, lv))
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("RemoveLogicalVolume: %s", out)
		return err
	}
	return nil
}

func LogicalDevice(vgName, lvName string) string {
	return fmt.Sprintf("/dev/%s/%s", vgName, lvName)
}

//func lvdisplayGrep(vgName, lvName, keyword string) (bool, string, error) {
//	var b2 bytes.Buffer
//
//	cmd1 := exec.Command("lvdisplay", fmt.Sprintf("/dev/%s/%s", vgName, lvName))
//	cmd2 := exec.Command("grep", keyword)
//
//	r, w := io.Pipe()
//	cmd1.Stdout = w
//	cmd2.Stdin = r
//	cmd2.Stdout = &b2
//
//	if err := cmd1.Start(); err != nil {
//		return false, "", err
//	}
//	if err := cmd2.Start(); err != nil {
//		return false, "", err
//	}
//	if err := cmd1.Wait(); err != nil {
//		return false, "", err
//	}
//	w.Close()
//	if err := cmd2.Wait(); err != nil {
//		//exitCode, inErr := system.GetExitCode(err)
//		//if inErr != nil {
//		//	return false, "", inErr
//		//}
//		//if exitCode != 1 {
//		//	return false, "", err
//		//}
//		return false, "", err
//	}
//
//	if b2.Len() != 0 {
//		return true, b2.String(), nil
//	}
//	return false, "", nil
//}
