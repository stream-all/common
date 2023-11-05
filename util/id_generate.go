package util

import "github.com/sony/sonyflake"

var sf *sonyflake.Sonyflake = nil

func InitIDGenerate(machineID uint16) {
	st := sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return machineID, nil
		},
	}
	sf = sonyflake.NewSonyflake(st)
}

func GenID() (uint64, error) {
	return sf.NextID()
}
