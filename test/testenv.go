package test

import (
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"ynabtui/internal/app"
)

type TestEnv struct {
	t         *testing.T
	Ynab      *FakeYnab
	testModel *teatest.TestModel
}

func NewTestEnv(t *testing.T) *TestEnv {
	return &TestEnv{
		t:    t,
		Ynab: NewFakeYnab(),
	}
}

func (env *TestEnv) Start() *teatest.TestModel {
	env.testModel = teatest.NewTestModel(env.t, app.InitialModel(env.Ynab.Api()))
	return env.testModel
}

func (env *TestEnv) Type(s string) {
	env.testModel.Type(s)
}

func (env *TestEnv) AssertOutput(f func([]byte) bool) {
	teatest.WaitFor(
		env.t, env.testModel.Output(),
		f,
		teatest.WithCheckInterval(5*time.Millisecond),
		teatest.WithDuration(500*time.Millisecond),
	)
}

func (env *TestEnv) Quit() {
	assert.NoError(env.t, env.testModel.Quit())
}
