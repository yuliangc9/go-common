package main

import (
	"common"
	"common/sessioncontext"
)

func main() {
	common.Debug("test debug")
	common.Info("test info")

	sc := sessioncontext.NewSessionCtx("")
	sc.Debug("hehe", "haha")
	sc.Warn("nothing")
}
