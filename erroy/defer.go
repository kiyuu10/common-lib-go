package erroy

import (
	"fmt"
	"log"
)

var vRootDeferFunctions []func()

func RegisterRootDefer(deferFunc func()) {
	vRootDeferFunctions = append(vRootDeferFunctions, deferFunc)
}

func ExecuteRootDefers() {
	for _, deferFunc := range vRootDeferFunctions {
		func() {
			defer recoverRootDefer()
			deferFunc()
		}()
	}
}

func recoverRootDefer() {
	errObj := recover()
	if errObj == nil {
		return
	}
	err, ok := errObj.(error)
	if !ok {
		err = fmt.Errorf("error: %v", errObj)
	}
	log.Printf("execute root defer failed | err=%s\n", err.Error())
}
