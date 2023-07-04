package main

import (
	"go-automation-api/__test__/stepdef"

	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	stepdef.LoginSteps(ctx)
}
