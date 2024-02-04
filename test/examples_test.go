package test

import (
	"context"
	"github.com/stretchr/testify/assert"
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

func TestServiceAccountWithInvalidAccountId(t *testing.T) {
	t.Parallel()

	runId := strings.ToLower(random.UniqueId())
	exampleModuleName := "service_account_with_keys"
	terraformOptions := &terraform.Options{}

	nameEndingInHyphen := runId + "-"
	nameStartingWithHyphen := "-" + runId

	shortName := "short" + runId
	if len(shortName) > 5 {
		shortName = shortName[len(shortName)-5:]
	}

	longName := "thishasmorethanthirtycharacters" + runId
	if len(longName) > 31 {
		longName = longName[len(longName)-31:]
	}

	invalidAccountIds := []string{nameEndingInHyphen, nameStartingWithHyphen, shortName, longName}

	for _, invalidAccountId := range invalidAccountIds {

		test_structure.RunTestStage(t, exampleModuleName+"_create_test_copy", func() {
			tempTestDir := test_structure.CopyTerraformFolderToTemp(t, "..", "examples/"+exampleModuleName)

			copyErr := copySupportingFiles(
				[]string{
					"providers.tf",
				}, tempTestDir)

			if copyErr != nil {
				t.Fatal("Failed to copy supporting files: " + copyErr.Error())
			}

			dummyGoogleCredentials := `{
									"type": "service_account",
									"project_id": "project-id",
									"private_key_id": "key-id",
									"private_key": "-----BEGIN PRIVATE KEY-----\n(private-key)\n-----END PRIVATE KEY-----\n",
									"client_email": "service-account-email",
									"client_id": "client-id",
									"auth_uri": "https://accounts.google.com/o/oauth2/auth",
									"token_uri": "https://oauth2.googleapis.com/token",
									"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
									"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/service-account-email"
								}`

			terraformOptions = &terraform.Options{
				TerraformDir: tempTestDir,
				Vars: map[string]interface{}{
					"account_id":         invalidAccountId,
					"project_id":         "dummy",
					"google_credentials": dummyGoogleCredentials,
					"google_region":      "asia-southeast2",
				},
			}
		})

		test_structure.RunTestStage(t, exampleModuleName+"_validate", func() {
			//err := terraform.InitAndPlanE(t, terraformOptions)
			_, err := terraform.InitAndPlanE(t, terraformOptions)
			assert.NotNil(t, err, "All given account IDs should be invalid")
		})
	}
}
