package test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestGoogleServiceAccountInProject(t *testing.T) {

	run := strings.ToLower(random.UniqueId())

	// Can be any project and region, need to source from env var
	gcpProject := "test-terraform-project-01"
	gcpIndonesiaRegion := "asia-southeast2"

	// GCP credentials will be sourced from this var. Do not use `GOOGLE_CREDENTIALS`
	// since we will be using that for validating Terraform-created credentials.
	googleCredentialsEnvVarName := "TERRATEST_GOOGLE_CREDENTIALS"
	googleCredentialsFromEnvVar := os.Getenv(googleCredentialsEnvVarName)
	assert.NotZero(t, len(googleCredentialsFromEnvVar), googleCredentialsEnvVarName+" environment variable not set!")
	terraformOptions := &terraform.Options{}

	// TODO - not needed currently
	//t.Run("CreateGoogleServiceAccountTestCase", func(t *testing.T) {
	//	t.Parallel()
	//})

	defer test_structure.RunTestStage(t, "terraform_destroy", func() {
		terraform.Destroy(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "create_test_copy", func() {
		tempTestDir := test_structure.CopyTerraformFolderToTemp(t, "..", "modules/google_service_account")
		logger.Log(t, "Temporary test directory created at %s", tempTestDir)

		copyErr := copySupportingFiles(
			[]string{
				"providers.tf",
			}, tempTestDir)

		if copyErr != nil {
			t.FailNow()
		}

		terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
			TerraformDir: tempTestDir,
			Vars: map[string]interface{}{
				"account_id":         "terratest-" + run,
				"description":        "created by terratest run: " + run,
				"display_name":       "terratest-" + run,
				"key_aliases":        []string{"primary"},
				"project":            gcpProject,
				"google_credentials": os.Getenv(googleCredentialsEnvVarName),
				"google_region":      gcpIndonesiaRegion,
				"in_project_roles":   []string{"roles/viewer"},
			},
		})
	})

	test_structure.RunTestStage(t, "terraform_init", func() {
		terraform.Init(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "terraform_plan", func() {
		terraform.Plan(t, terraformOptions)
	})

	// Create the service account and credentials
	test_structure.RunTestStage(t, "terraform_apply", func() {
		terraform.Apply(t, terraformOptions)
	})

	// Set up the credentials locally and "log in"
	test_structure.RunTestStage(t, "validate_service_account_key", func() {

		// TODO: Try to use subshell
		//privateKeyMap := terraform.OutputMap(t, terraformOptions, "service_account_keys")
		jsonOutput := terraform.OutputJson(t, terraformOptions, "service_account_keys")
		//
		//primaryKey := privateKeyMap["primary"]
		//privateKey := primaryKey["private_key"]
		//
		//os.WriteFile("key.json", []byte(privateKey), 0400)
		//os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "key.json")
		//
		//serviceAccountEmail := terraform.Output(t, terraformOptions, "account_email")
		//
		//_ = gcp.GetLoginProfile(t, serviceAccountEmail)
		//os.Remove("key.json")

		//log.Println("private key map: ", privateKeyMap)
		log.Println("JSON output: ", jsonOutput)

		// Unmarshaling the JSON above requires a struct etc etc
		// For now just checking that we got something back.
		assert.NotZero(t, len(jsonOutput), "Output should not be empty.")
	})
}
