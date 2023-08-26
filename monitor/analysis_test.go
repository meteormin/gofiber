package monitor

import (
	"github.com/miniyus/gofiber"
	"testing"
)

func TestNewAnalysis(t *testing.T) {
	a := gofiber.New()
	a.Bootstrap()
	ai := NewAnalysis(a)

	marshal, err := ai.Marshal(true)

	if err != nil {
		t.Log(ai)
		t.Error(err)
	}

	t.Log(marshal)
}
