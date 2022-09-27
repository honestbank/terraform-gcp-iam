package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
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
		//keys, err := terraform.OutputMapE(t, terraformOptions, "keys")
		//if err != nil {
		//	log.Println(err.Error())
		//}

		keysString, err := terraform.OutputJsonE(t, terraformOptions, "keys")
		if err != nil {
			log.Printf("OutputJsonE: %v", err.Error())
		}

		outputMap := map[string]interface{}{}
		if err := json.Unmarshal([]byte(keysString), &outputMap); err != nil {
			log.Printf("OutputJsonE-Unmarshal: %v", err.Error())
		}

		keys := make(map[string]string)
		for k, v := range outputMap {
			keys[k] = fmt.Sprintf("%v", v)
		}

		f, _ := os.ReadFile(terraformOptions.TerraformDir + "/terraform.tfstate")
		os.WriteFile("./terraform.tfstate", f, 0777)

		log.Printf("tfstate: %v", string(f))
		log.Printf("len(keys): %v\n", len(keys))

		for keyName, keySecret := range keys {
			log.Printf("keyname: %v keySecret: %v\n", keyName, keySecret)

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
