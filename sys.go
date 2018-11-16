package goutil

import (
	"log"
	"runtime"
)

func GetMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func GetExecutorName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func ChkErr(err error) {
	if err != nil {
		log.Println("executor:", GetExecutorName())
		log.Panic(err)
	}
}
