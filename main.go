package main

import (
	"log"
	"os"

	"go-automation-api/__test__/stepdef"

	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
)

func init() {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	stepdef.LoginSteps(ctx)
}

func main() {
	opts := godog.Options{
		Format:    "progress",
		Strict:    true,
		Randomize: 0,
	}

	suite := godog.TestSuite{
		Name:                "Login API",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}

	if status := suite.Run(); status > 0 {
		os.Exit(status)
	}
}
