package pipeline

import "time"

type BaseTask struct {
	outputs []Task
	inputs  []Task

	id      int64
	dotID   string
	Index   int32         `mapstructure:"index" json:"-" `
	Timeout time.Duration `mapstructure:"timeout"`
}

func NewBaseTask(id int64, dotID string, outputs []Task, index int32) BaseTask {
	return BaseTask{dotID: dotID, outputs: outputs, Index: index}
}

func (t *BaseTask) Base() *BaseTask {
	return t
}

func (t BaseTask) ID() int64 {
	return t.id
}

func (t BaseTask) DotID() string {
	return t.dotID
}

func (t BaseTask) OutputIndex() int32 {
	return t.Index
}

func (t BaseTask) Outputs() []Task {
	return t.outputs
}

func (t BaseTask) Inputs() []Task {
	return t.inputs
}

func (t BaseTask) TaskTimeout() (time.Duration, bool) {
	if t.Timeout == time.Duration(0) {
		return time.Duration(0), false
	}
	return t.Timeout, true
}
