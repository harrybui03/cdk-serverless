package utils

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
	"os"
)

func NameWithEnv(name string, camelCase ...bool) string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "stg"
	}
	if len(camelCase) > 0 && camelCase[0] {
		return fmt.Sprintf("%s%s", env, name)
	}
	return fmt.Sprintf("%s_%s", env, name)
}

func JsiiWithEnv(s string, camelCase ...bool) *string {
	return jsii.String(NameWithEnv(s, camelCase...))
}
