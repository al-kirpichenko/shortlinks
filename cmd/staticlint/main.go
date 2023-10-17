package main

import (
	"log"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/defers"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/unmarshal"
	"honnef.co/go/tools/staticcheck"

	"staticlint/myanalizer"
)

func main() {

	var mychecks []*analysis.Analyzer

	for _, v := range staticcheck.Analyzers {
		mychecks = append(mychecks, v.Analyzer)
	}

	mychecks = append(mychecks, myanalizer.OsExitCheckAnalyzer)
	mychecks = append(mychecks, printf.Analyzer)
	mychecks = append(mychecks, shadow.Analyzer)
	mychecks = append(mychecks, structtag.Analyzer)
	mychecks = append(mychecks, defers.Analyzer)
	mychecks = append(mychecks, unmarshal.Analyzer)

	for _, s := range mychecks {
		log.Println(s)
	}

	multichecker.Main(
		mychecks...,
	)
}
