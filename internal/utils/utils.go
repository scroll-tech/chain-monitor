package utils

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/modern-go/reflect2"
	"github.com/scroll-tech/go-ethereum/common"
)

var (
	Precision = big.NewInt(1e14)
)

// TryTimes try run several times until the function return true.
func TryTimes(ctx context.Context, times int, run func() bool) bool {
	for i := 0; i < times; i++ {
		select {
		case <-ctx.Done():
			return false
		default:
			if run() {
				return true
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
	return false
}

// LoopWithContext Run the f func with context periodically.
func LoopWithContext(ctx context.Context, period time.Duration, f func(ctx context.Context)) {
	tick := time.NewTicker(period)
	defer tick.Stop()
	for true {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			f(ctx)
		}
	}
}

// Loop Run the f func periodically.
func Loop(ctx context.Context, period time.Duration, f func()) {
	tick := time.NewTicker(period)
	defer tick.Stop()
	for true {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			f()
		}
	}
}

// IsNil Check if the interface is empty.
func IsNil(i interface{}) bool {
	return i == nil || reflect2.IsNil(i)
}

func Copy(dest, source interface{}) error {
	dv := reflect.ValueOf(source)
	//sv := reflect.ValueOf(source)
	for i := 0; i < dv.NumField(); i++ {
		val := dv.Index(i)
		fmt.Println(val.String())
	}
	return nil
}

func GetID(data []byte) string {
	if len(data) >= 4 {
		return common.Bytes2Hex(data[:4])
	}
	return common.Bytes2Hex(data)
}
