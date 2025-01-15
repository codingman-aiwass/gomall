package snowflake

import (
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	settings := sonyflake.Settings{
		StartTime: time.Now(),
	}
	sf = sonyflake.NewSonyflake(settings)
	if sf == nil {
		panic("Sonyflake not initialized")
	}
}

func GenerateTransactionID() uint64 {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return id
}
