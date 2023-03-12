package lvm

import (
	"fmt"
)

// https://man7.org/linux/man-pages/man8/lvs.8.html

/**
Позоции атрибутов в сырой строке 'Vwi-a-tz--'
в нулевой позиции - кавычки. см. комментарий в UnmarshalJSON
*/
const (
	attrPosVolumeType int = iota + 1
	attrPosPermissions
	attrPosAllocPolicy
	attrPosFixedMinor
	attrPosState
	attrPosDevice
	attrPosTargetType
	attrPosZeroNewBlocksOnAlloc
	attrPosHealth
	attrPosSkipActivation
)

var lvVolumeTypeFlags = map[byte]LvVolumeType{
	'C': Cache,
	'm': Mirrored,
	'M': MirroredNoInitialSync,
	'o': Origin,
	'O': OriginMerging,
	'r': Raid,
	'R': RaidNoInitialSync,
	's': Snapshot,
	'S': SnapshotMerging,
	'p': PvMove,
	'v': Virtual,
	'i': Image,
	'I': ImageOutOfSync,
	'l': Log,
	'c': UnderConversion,
	'V': ThinVolume,
	't': ThinPool,
	'T': ThinPoolData,
	'd': VdoPool,
	'D': VdoPoolData,
	'e': PoolMetaSpare,
}

var lvPermissionFlags = map[byte]LvPermissions{
	'w': Writeable,
	'r': ReadOnly,
	'R': RoActivationOfRw,
}

var lvAllocPolicyFlags = map[byte]LvAllocPolicy{
	'a': Anywhere,
	'c': Contiguous,
	'i': Inherited,
	'l': Cling,
	'n': Normal,
	'A': AnywhereLocked,
	'C': ContiguousLocked,
	'I': InheritedLocked,
	'L': ClingLocked,
	'N': NormalLocked,
}

var lvFixedMinorFlags = map[byte]bool{
	'm': true,
	'-': false,
}

var lvStateFlags = map[byte]LvState{
	'-': Inactive,
	'a': Active,
	'h': Historical,
	's': Suspended,
	'I': InvalidSnap,
	'S': InvalidSuspendedSnap,
	'm': SnapMergeFailed,
	'M': SuspendedSnapMergeFailed,
	'd': MappedDevWithoutTables,
	'i': MappedDevWithInactiveTable,
	'c': CheckThinPool,
	'C': CheckSuspendedThinPool,
	'X': StateUnknown,
}

var lvDeviceStatusFlags = map[byte]LvDeviceStatus{
	'o': Mounted,
	'-': NotMounted,
	'X': DevStateUnknown,
}

var lvTargetTypeFlags = map[byte]LvTargetType{
	'C': TargetTypeCache,
	'm': TargetTypeMirror,
	'r': TargetTypeRaid,
	's': TargetTypeSnapshot,
	't': TargetTypeThin,
	'u': TargetTypeUnknown,
	'v': TargetTypeVirtual,
}

var lvZeroNewBlocksOnAllocFlags = map[byte]bool{
	'z': true,
	'-': false,
}

var lvHealthFlags = map[byte]LvHealth{
	'p': HealthPartial,
	'X': HealthUnknown,
	'-': HealthUnknown,
	// RAID
	'r': RaidRefreshNeeded,
	'm': RaidMismatches,
	'w': RaidWritemostly,
	's': RaidReshaping,
	'R': RaidRemoveAfterReshape,
	//thin
	'F': ThinFailed,
	'D': ThinOutOfData,
	'M': ThinMetaFailed,
	//writecache
	'E': WritecacheError,
}

var lvSkipActivationFlags = map[byte]bool{
	'k': true,
	'-': false,
}

func (attrs *LvAttributes) UnmarshalJSON(raw []byte) error {
	if len(raw) != 12 {
		// 12 символов - сами флаги и кавычки ("Vwi-a-tz--"). так поступает в UnmarshalJSON
		return fmt.Errorf("10 symbols needed to parse LvAttributes. Got %d: '%s'", len(raw), raw)
	}

	*attrs = LvAttributes{
		VolumeType:           lvVolumeTypeFlags[raw[attrPosVolumeType]],
		Permissions:          lvPermissionFlags[raw[attrPosPermissions]],
		AllocPolicy:          lvAllocPolicyFlags[raw[attrPosAllocPolicy]],
		FixedMinor:           lvFixedMinorFlags[raw[attrPosFixedMinor]],
		State:                lvStateFlags[raw[attrPosState]],
		Device:               lvDeviceStatusFlags[raw[attrPosDevice]],
		TargetType:           lvTargetTypeFlags[raw[attrPosTargetType]],
		ZeroNewBlocksOnAlloc: lvZeroNewBlocksOnAllocFlags[raw[attrPosZeroNewBlocksOnAlloc]],
		Health:               lvHealthFlags[raw[attrPosHealth]],
		SkipActivation:       lvSkipActivationFlags[raw[attrPosSkipActivation]],
		Raw:                  string(raw),
	}

	return nil
}

func (attrs *LvAttributes) Map() map[string]interface{} {
	return map[string]interface{}{
		"Volume type":       attrs.VolumeType.String(),
		"Permissions":       attrs.Permissions.String(),
		"Allocation policy": attrs.AllocPolicy.String(),
		"Fixed minor":       attrs.FixedMinor,
		"State":             attrs.State.String(),
		"Device":            attrs.Device.String(),
		"Target type":       attrs.TargetType.String(),
		"Newly-allocated data blocks are zeroed before use": attrs.ZeroNewBlocksOnAlloc,
		"Health":          attrs.Health.String(),
		"Skip activation": attrs.SkipActivation,
		//"raw":             attrs.Raw,
	}

}
