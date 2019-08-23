package db

import (
	"github.com/winjeg/go-commons/log"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sony/sonyflake"
)

// package that provides unique ids
// 1. snowflake id
// 2. uuid

var (
	timeStart, _  = time.Parse("2006-01-02 15:04:05", "2019-07-01 12:00:00")
	snowFlakeInst = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime:      timeStart, // start time
		MachineID:      nil,       // can be replaced, default is private ip address
		CheckMachineID: nil,       // the method to make sure the machine id is unique
	})

	logger = log.GetLogger(nil)
)

func NextID() uint64 {
	id, err := snowFlakeInst.NextID()
	if err != nil {
		logger.Error(err)
		return 0
	}
	return id
}

// export UUID
// is to generate unique ids
func UUID() string {
	uid := uuid.NewV4()
	return uid.String()
}

// export UUIDShort
// this method will generate a unique id using uuid, but the result is too long
// so we just use the digits from 0 to 8, thus, increasing the possibility to get a
// duplicated id, but It's okay
// not true uuid, not for tons of ids
func UUIDShort() string {
	u2 := uuid.NewV4()
	d := u2.String()
	return d[24:] + d[9:13]
}
