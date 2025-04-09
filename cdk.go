package main

import (
	"cdk/register"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewMainAppStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	register.MainApp(stack)
	return stack
}
func main() {
	defer jsii.Close()
	app := awscdk.NewApp(nil)

	envStr := os.Getenv("ENVIRONMENT")
	if envStr == "" {
		envStr = "stg"
	}
	name := fmt.Sprintf("%sCdkStack", envStr)
	NewMainAppStack(app, name, &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
