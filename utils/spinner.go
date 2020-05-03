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

var SpinStart = func(s *spinner.Spinner) {
	s.Start()
}

var SpinStop = func(s *spinner.Spinner) {
	s.Stop()
}

func Spinner(suffix string) SpinMessage {
	spin := spinner.New(spinner.CharSets[defaultSpinner], defaultSpinnerTime*time.Millisecond)
	spin.Color(defaultSpinnerColor)
	spin.Suffix = suffix
	return spin
}
