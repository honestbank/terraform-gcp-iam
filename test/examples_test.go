package test

import (
	"context"
	"strings"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	storageApi "google.golang.org/api/storage/v1"
)

func TestServiceAccountWithKeysExample(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	runId := strings.ToLower(random.UniqueId())

	exampleModuleName := "service_account_with_keys"
	terraformOptions := &terraform.Options{}

	defer test_structure.RunTestStage(t, exampleModuleName+"_destroy", func() {
		terraform.Destroy(t, terraformOptions)
	})

	test_structure.RunTestStage(t, exampleModuleName+"_create_test_copy", func() {
		tempTestDir := test_structure.CopyTerraformFolderToTemp(t, "..", "examples/"+exampleModuleName)

		copyErr := copySupportingFiles(
			[]string{
				"providers.tf",
			}, tempTestDir)

		if copyErr != nil {
			t.Fatal("Failed to copy supporting files: " + copyErr.Error())
		}

		terraformOptions = &terraform.Options{
			TerraformDir: tempTestDir,
			Vars: map[string]interface{}{
				"account_id":         "terraform-test-" + runId,
				"project_id":         terratestGoogleProjectCompute,
				"google_credentials": terratestGoogleCredentials,
				"google_region":      "asia-southeast2",
			},
		}
	})

	test_structure.RunTestStage(t, exampleModuleName+"_init_and_apply", func() {
		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, exampleModuleName+"_result_assertions", func() {
		keys := terraform.OutputMap(t, terraformOptions, "keys")

		for keyName, keySecret := range keys {
			creds, err := google.CredentialsFromJSON(ctx, []byte(keySecret), storageApi.CloudPlatformScope)
			if err != nil {
				t.Fatalf("Failed to parse JSON credentials for %s key: %s", keyName, err.Error())
			}

			storageClientWithGeneratedKey, err := storage.NewClient(ctx, option.WithCredentials(creds))
			if err != nil {
				t.Fatalf("Failed to create storage client with %s key: %s", keyName, err.Error())
			}

			_, err = storageClientWithGeneratedKey.ServiceAccount(ctx, terratestGoogleProjectCompute)
			if err != nil {
				t.Fatalf("Failed to get default Service Account for GCS %s key: %s", keyName, err.Error())
			}
		}
	})
}

func TestServiceAccountWithMembershipsExample(t *testing.T) {
	t.Parallel()

	// Write test covering different membership types with and without conditions.
}
