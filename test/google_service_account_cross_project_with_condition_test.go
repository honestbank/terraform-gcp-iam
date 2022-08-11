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

func TestGoogleServiceAccountCrossProjectWithCondition(t *testing.T) {
	run := strings.ToLower(random.UniqueId())

	// [roles/iam.serviceAccountAdmin, roles/iam.serviceAccountKeyAdmin] required in this project
	gcpServiceAccountHostProject := os.Getenv("GOOGLE_PROJECT")

	// [roles/iam.securityAdmin] required in this project
	gcpCrossProjectIamRoleMembershipProjectId := "storage-0994"
	gcpIndonesiaRegion := "asia-southeast2"

	// GCP credentials will be sourced from this var. Do not use `GOOGLE_CREDENTIALS`
	// since we will be using that for validating Terraform-created credentials.
	googleCredentialsEnvVarName := "TERRATEST_GOOGLE_CREDENTIALS"
	googleCredentialsFromEnvVar := os.Getenv(googleCredentialsEnvVarName)
	assert.NotZero(t, len(googleCredentialsFromEnvVar), googleCredentialsEnvVarName+" environment variable not set!")
	terraformOptions := &terraform.Options{}

	testCaseName := "google_service_account_cross_project"

	defer test_structure.RunTestStage(t, testCaseName+"_deferred_destroy", func() {
		terraform.Destroy(t, terraformOptions)
	})

	test_structure.RunTestStage(t, testCaseName+"_create_test_copy", func() {
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
				"account_id":               "terratest-" + run,
				"description":              "created by terratest run: " + run,
				"display_name":             "terratest-" + run,
				"key_aliases":              []string{"primary"},
				"project":                  gcpServiceAccountHostProject,
				"google_credentials":       os.Getenv(googleCredentialsEnvVarName),
				"google_region":            gcpIndonesiaRegion,
				"iam_role_membership_type": "CROSS_PROJECT",
				"cross_project_iam_role_memberships": map[string][]string{
					gcpCrossProjectIamRoleMembershipProjectId: {"roles/storage.objectViewer"},
				},
				"conditions": []map[string]string{
					{
						"title":       "terratest-" + run,
						"expression":  "resource.service == 'storage.googleapis.com'",
						"description": "test condition description" + run,
					},
				},
			},
		})
	})

	test_structure.RunTestStage(t, testCaseName+"_terraform_init", func() {
		terraform.Init(t, terraformOptions)
	})

	test_structure.RunTestStage(t, testCaseName+"_terraform_plan", func() {
		terraform.Plan(t, terraformOptions)
	})

	// Create the service account and credentials
	test_structure.RunTestStage(t, testCaseName+"_terraform_apply", func() {
		terraform.Apply(t, terraformOptions)
	})

	// Set up the credentials locally and "log in"
	test_structure.RunTestStage(t, testCaseName+"_validate_service_account_key", func() {
		jsonOutput := terraform.OutputJson(t, terraformOptions, "service_account_keys")

		//log.Println("private key map: ", privateKeyMap)
		log.Println("JSON output: ", jsonOutput)

		// Unmarshaling the JSON above requires a struct etc etc
		// For now just checking that we got something back.
		assert.NotZero(t, len(jsonOutput), "Output should not be empty.")
	})
}

func TestGoogleServiceAccountCrossProjectMultipleProjectsWithCondition(t *testing.T) {
	run := strings.ToLower(random.UniqueId())

	// [roles/iam.serviceAccountAdmin, roles/iam.serviceAccountKeyAdmin] required in this project
	gcpServiceAccountHostProject := os.Getenv("GOOGLE_PROJECT")

	// [roles/iam.securityAdmin] required in this project
	gcpCrossProjectIamRoleMembershipProjectId := "storage-0994"
	gcpCrossProjectIamRoleMembershipProjectId2 := "tf-shared-vpc-host-78a3"
	gcpIndonesiaRegion := "asia-southeast2"

	// GCP credentials will be sourced from this var. Do not use `GOOGLE_CREDENTIALS`
	// since we will be using that for validating Terraform-created credentials.
	googleCredentialsEnvVarName := "TERRATEST_GOOGLE_CREDENTIALS"
	googleCredentialsFromEnvVar := os.Getenv(googleCredentialsEnvVarName)
	assert.NotZero(t, len(googleCredentialsFromEnvVar), googleCredentialsEnvVarName+" environment variable not set!")
	terraformOptions := &terraform.Options{}

	testCaseName := "google_service_account_cross_project"

	//defer test_structure.RunTestStage(t, testCaseName+"_deferred_destroy", func() {
	//	terraform.Destroy(t, terraformOptions)
	//})

	test_structure.RunTestStage(t, testCaseName+"_create_test_copy", func() {
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
				"account_id":               "terratest-" + run,
				"description":              "created by terratest run: " + run,
				"display_name":             "terratest-" + run,
				"key_aliases":              []string{"primary"},
				"project":                  gcpServiceAccountHostProject,
				"google_credentials":       os.Getenv(googleCredentialsEnvVarName),
				"google_region":            gcpIndonesiaRegion,
				"iam_role_membership_type": "CROSS_PROJECT",
				// Two projects should cause an error
				"cross_project_iam_role_memberships": map[string][]string{
					gcpCrossProjectIamRoleMembershipProjectId:  {"roles/storage.objectViewer"},
					gcpCrossProjectIamRoleMembershipProjectId2: {"roles/storage.objectViewer"},
				},
				"conditions": []map[string]string{
					{
						"title":       "terratest-" + run,
						"expression":  "resource.service == 'storage.googleapis.com'",
						"description": "test condition description" + run,
					},
				},
			},
		})
	})

	test_structure.RunTestStage(t, testCaseName+"_terraform_init", func() {
		terraform.Init(t, terraformOptions)
	})

	test_structure.RunTestStage(t, testCaseName+"_terraform_plan", func() {
		_, planErr := terraform.PlanE(t, terraformOptions)
		assert.Nil(t, planErr, "plan should not error when 2 external projects are specified")
	})
}
