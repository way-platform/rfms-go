//go:build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// CI runs the full CI pipeline.
func CI() {
	mg.Deps(Generate, Test)
}

// Clean removes all build artifacts.
func Clean() error {
	return sh.Run("rm", "-rf", "build")
}

// Test runs the Go tests.
func Test() error {
	mg.Deps(Generate)
	return sh.RunV("go", "test", "-v", "./...")
}

// Generate runs all code generators.
func Generate() error {
	mg.Deps(Tools.OpenAPIGenerator)
	return sh.RunV("go", "generate", "./...")
}

type Tools mg.Namespace

// OpenAPIGenerator downloads the OpenAPI generator JAR file.
func (Tools) OpenAPIGenerator() error {
	const (
		openAPIGeneratorVersion = "7.13.0"
		openAPIGeneratorURL     = "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/" + openAPIGeneratorVersion + "/openapi-generator-cli-" + openAPIGeneratorVersion + ".jar"
		targetFile              = "build/openapi-generator-cli-" + openAPIGeneratorVersion + ".jar"
	)
	if _, err := os.Stat(targetFile); err == nil {
		return nil
	}
	return sh.Run("curl", "--create-dirs", "-L", "-o", targetFile, openAPIGeneratorURL)
}
