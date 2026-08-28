package version

import (
	"fmt"
	"runtime"
)

var (
	// Version é a versão semântica da aplicação (injetada via ldflags no build).
	Version = "dev"

	// Commit é o hash do commit git atual (injetado via ldflags no build).
	Commit = "none"

	// Date é a data ISO8601 da compilação (injetada via ldflags no build).
	Date = "unknown"
)

// Info contém a estrutura de metadados da compilação da aplicação.
type Info struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	Date      string `json:"date"`
	GoVersion string `json:"go_version"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}

// GetInfo retorna a estrutura preenchida com os metadados de versão da aplicação.
func GetInfo() Info {
	return Info{
		Version:   Version,
		Commit:    Commit,
		Date:      Date,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// Formatted retorna a representação textual formatada da versão.
func Formatted(appName string) string {
	if appName == "" {
		appName = "md_server"
	}
	return fmt.Sprintf("%s version: %s (commit: %s, built at: %s, %s/%s)",
		appName, Version, Commit, Date, runtime.GOOS, runtime.GOARCH)
}
