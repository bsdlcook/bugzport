package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

const (
	defaultSpinner      int           = 11
	defaultSpinnerTime  time.Duration = 120
	defaultSpinnerColor string        = "blue"
)

type SpinMessage *spinner.Spinner

func Spinner(suffix string) SpinMessage {
	spin := spinner.New(spinner.CharSets[defaultSpinner], defaultSpinnerTime*time.Millisecond)
	spin.Color(defaultSpinnerColor)
	spin.Suffix = suffix

	return spin
}

func SpinStart(s *spinner.Spinner) {
	s.Start()
}

func SpinStop(s *spinner.Spinner) {
	s.Stop()
}
