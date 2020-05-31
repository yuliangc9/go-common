package sessioncontext

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/google/uuid"
)

type SessionCtx struct {
	Id string
}

func NewSessionCtx(id string) *SessionCtx {
	if id == "" {
		id = uuid.New().String()
	}
	return &SessionCtx{
		Id: id,
	}
}

func (this *SessionCtx) preLog() string {
	return fmt.Sprintf("[session:%s]", this.Id)
}

func (this *SessionCtx) Debug(f interface{}, v ...interface{}) {
	args := append([]interface{}{f}, v...)
	logs.Debug(this.preLog(), args...)
}

func (this *SessionCtx) Info(f interface{}, v ...interface{}) {
	args := append([]interface{}{f}, v...)
	logs.Info(this.preLog(), args...)
}

func (this *SessionCtx) Notice(f interface{}, v ...interface{}) {
	args := append([]interface{}{f}, v...)
	logs.Notice(this.preLog(), args...)
}

func (this *SessionCtx) Warn(f interface{}, v ...interface{}) {
	args := append([]interface{}{f}, v...)
	logs.Warn(this.preLog(), args...)
}

func (this *SessionCtx) Error(f interface{}, v ...interface{}) {
	args := append([]interface{}{f}, v...)
	logs.Error(this.preLog(), args...)
}

func (this *SessionCtx) Alert(f interface{}, v ...interface{}) {
	args := append([]interface{}{f}, v...)
	logs.Alert(this.preLog(), args...)
}
