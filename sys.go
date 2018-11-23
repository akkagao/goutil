package goutil

import (
	"log"
	"runtime"
)

/**
* 获取当前方法的名称
 */
func GetMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

/**
* 获取执行者的名称
 */
func GetExecutorName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

/**
* 判断error是否为nil如果不为nil 输出执行者名称，跑出panic
 */
func ChkErr(err error) {
	if err != nil {
		log.Println("executor:", GetExecutorName())
		log.Panic(err)
	}
}
