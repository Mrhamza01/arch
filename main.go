package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("❌ Usage: arch <FeatureName> [TargetPath]")
		return
	}

	featureName := os.Args[1]
	targetPath := "."

	if len(os.Args) >= 3 {
		targetPath = os.Args[2]
	}

	basePath := filepath.Join(targetPath, featureName)

	fmt.Printf("📂 Creating feature structure for: %s in %s\n", featureName, targetPath)

	// Domain
	makeDir(filepath.Join(basePath, "Domain", "Entities"))
	makeDir(filepath.Join(basePath, "Domain", "Repositories"))
	makeDir(filepath.Join(basePath, "Domain", "UseCases"))
	makeDir(filepath.Join(basePath, "Domain", "Datasources"))

	// Data
	makeDir(filepath.Join(basePath, "Data", "Datasources"))
	makeDir(filepath.Join(basePath, "Data", "Repositories"))
	makeDir(filepath.Join(basePath, "Data", "DTOs"))
	makeDir(filepath.Join(basePath, "Data", "Mappings"))

	// Presentation
	makeDir(filepath.Join(basePath, "Presentation", "Pages"))

	makeDir(filepath.Join(basePath, "Tests", "Unit"))
	makeDir(filepath.Join(basePath, "Tests", "Integration"))
	makeDir(filepath.Join(basePath, "Tests", "Ui"))

	fmt.Printf("✅ Feature '%s' created at %s\n", featureName, basePath)
}

func makeDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("❌ Failed to create %s: %v\n", path, err)
	}
}
