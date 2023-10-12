package myanalizer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test_ExitInMainAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), OsExitCheckAnalyzer, "./...")
}
