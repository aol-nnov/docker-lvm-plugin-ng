package lvm

import (
	"fmt"
	"time"
)

type LvmResponse struct {
	Reports []Report `json:"report"`
}
type Report struct {
	Lvs []Volume `json:"lv"`
}

type Volume struct {
	Name        string       `json:"lv_name"`
	Size        string       `json:"lv_size"`
	Used        string       `json:"data_percent"`
	Vg          string       `json:"vg_name"`
	DevPath     string       `json:"lv_path"`
	Attributes  LvAttributes `json:"lv_attr"`
	ThinPool    string       `json:"pool_lv"`
	Origin      string       `json:"origin"`
	UsedMeta    string       `json:"metadata_percent"`
	Created     CreatedTime  `json:"lv_time"`
	CreatedHost string       `json:"lv_host"`
	FullName    string       `json:"lv_full_name"`
}

func (v *Volume) String() string {
	return fmt.Sprintf(
		"Name: %s (%s)\n"+
			"Size: %s\n"+
			"Used: %s%%\n"+
			"UsedMeta: %s%%\n"+
			"Vg: %s\n"+
			"DevPath: %s\n"+
			"Attributes:\n%s\n"+
			"ThinPool: %s\n"+
			"Origin: %s\n"+
			"Created: %s\n"+
			"CreatedHost: %s",
		v.Name, v.FullName,
		v.Size,
		v.Used,
		v.UsedMeta,
		v.Vg,
		v.DevPath,
		v.Attributes.String(),
		v.ThinPool,
		v.Origin,
		v.Created.Format(time.RFC822),
		v.CreatedHost,
	)
}
