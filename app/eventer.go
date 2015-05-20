package app

import (
	"fmt"
	"reflect"

	"github.com/hprose/hprose-go/hprose"
)

type myServiceEvent struct{}

func (myServiceEvent) OnBeforeInvoke(name string, args []reflect.Value, byref bool, context hprose.Context) {
	fmt.Println("OnBeforeInvoke", name, args, byref)
}

func (myServiceEvent) OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context hprose.Context) {
	fmt.Println("OnAfterInvoke", name, args, byref, result)
}

func (myServiceEvent) OnSendError(err error, context hprose.Context) {
	fmt.Println("OnSendError", err)
}
