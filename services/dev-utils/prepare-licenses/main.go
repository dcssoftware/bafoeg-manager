package main

import (
	"encoding/json"
	"flag"
	"os"

	"golang.org/x/mod/modfile"
)

type LicenseFile struct {
	Backend  []string `json:"backend"`
	Frontend []string `json:"frontend"`
}

func main() {
	var backendGoModFile string
	var frontendPackageJsonFile string
	var exportFilePath string

	flag.StringVar(&backendGoModFile, "backend", "../../backend/go.mod", "Specify location of backend go.mod file")
	flag.StringVar(&frontendPackageJsonFile, "frontend", "../../frontend/package.json", "Specify location of frontend package.json file")
	flag.StringVar(&exportFilePath, "export", "../../backend/src/static/assets/licenses/licenses.json", "Specify location for the exported json file")

	backendLicenses := getBackendLicenses(backendGoModFile)
	frontendLicenses := getFrontendLicenses(frontendPackageJsonFile)

	licenses := LicenseFile{
		Backend:  backendLicenses,
		Frontend: frontendLicenses,
	}

	json, err := json.Marshal(licenses)
	if err != nil {
		panic(err)
	}

	os.WriteFile(exportFilePath, json, 0755)
}

func getBackendLicenses(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	modFile, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		panic(err)
	}

	var licenses []string

	for _, require := range modFile.Require {
		if !require.Indirect {
			licenses = append(licenses, require.Mod.Path)
		}
	}

	return licenses
}

func getFrontendLicenses(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var packageJSON struct {
		Dependencies    map[string]string `json:"dependencies"`
		DevDependencies map[string]string `json:"devDependencies"`
	}

	err = json.Unmarshal(content, &packageJSON)
	if err != nil {
		panic(err)
	}

	var licenses []string

	for depName := range packageJSON.Dependencies {
		licenses = append(licenses, depName)
	}

	return licenses
}
