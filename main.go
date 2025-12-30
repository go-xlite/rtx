package rtx

import (
	args "github.com/go-xlite/rtx/handler/args"
	pid "github.com/go-xlite/rtx/handler/pid"
	rtm "github.com/go-xlite/rtx/handler/rtm"
)

type PidHandler = pid.PidHandler

var NewPidHandler = pid.NewPidHandler
var Rtm = rtm.Rtm
var Args = args.Args
