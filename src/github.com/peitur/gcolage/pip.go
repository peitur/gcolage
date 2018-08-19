package gcolage

import (
	"fmt"
)

func PipProjectInfoJson(project string) string {
	return fmt.Sprintf("https://pypi.org/pypi/%s/json", project)
}
