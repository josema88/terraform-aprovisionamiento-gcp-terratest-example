
package test

import (
	"crypto/tls"
	"testing"
	"time"
	"strings"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
    
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	
	// Define terraform variables - Set this value with your project ID
	project := ""

	
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../gcp-instance-module",
		Vars:  map[string]interface{} {
			"project": project,
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	output := terraform.Output(t, terraformOptions, "instance_name")
	assert.Equal(t, "my-instance", output)


	// Setup a TLS configuration to submit with the helper, a blank struct is acceptable
	tlsConfig := tls.Config{}

	// It can take a minute or so for the Instance to boot up, so retry a few times
	maxRetries := 30
	timeBetweenRetries := 5 * time.Second

	// Get app url
	app_url := terraform.Output(t, terraformOptions, "app_url")
	
	
	// Verify that we get back a 200 OK and the body contains specific text
	
	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		app_url,
		&tlsConfig,
		maxRetries,
		timeBetweenRetries,
		func(statusCode int, body string) bool {
			return statusCode == 200 && strings.Contains(body, "Hello world! From Woman Who Code Guatemala")
		},
	)
}