package test

import (
	"errors"
	"fmt"
	"os"

	"github.com/gruntwork-io/terratest/modules/files"
)

// Define credentials, project, etc at package level for reuse throughout tests and helpers
var envVars = map[string]string{
	"terratestGoogleCredentials":         "TERRATEST_GOOGLE_CREDENTIALS",
	"terratestGoogleProjectCompute":      "TERRATEST_GOOGLE_PROJECT_COMPUTE",
	"terratestGoogleProjectStorage":      "TERRATEST_GOOGLE_PROJECT_STORAGE",
	"terratestGoogleExtraFolderId":       "TERRATEST_GOOGLE_EXTRA_FOLDER_ID",
	"terratestGoogleProjectInsideFolder": "TERRATEST_GOOGLE_PROJECT_INSIDE_FOLDER",
}
var terratestGoogleCredentials = os.Getenv(envVars["terratestGoogleCredentials"])
var terratestGoogleProjectCompute = os.Getenv(envVars["terratestGoogleProjectCompute"])
var terratestGoogleProjectStorage = os.Getenv(envVars["terratestGoogleProjectStorage"])
var terratestGoogleExtraFolderId = os.Getenv(envVars["terratestGoogleExtraFolderId"])
var terratestGoogleProjectInsideFolder = os.Getenv(envVars["terratestGoogleProjectInsideFolder"])

// copySupportingFiles copies one or more files from the current working directory into a destination dir.
// destination should be specified by its absolute (not relative) path beginning with "/".
// This is done to configure providers when using modules without them explicitly defined.
func copySupportingFiles(fileNames []string, destination string) error {
	testFileSourceDir, getTestDirErr := os.Getwd()
	if getTestDirErr != nil {
		msg := fmt.Sprint("Calling t.FailNow(): could not execute os.Getwd(): ", getTestDirErr)
		return errors.New(msg)
	}

	fmt.Println("Test working directory is: ", testFileSourceDir)

	fmt.Println("Copying files: ", fileNames, " to temporary test dir: ", destination)
	for _, file := range fileNames {
		src := testFileSourceDir + "/" + file
		dest := destination + "/" + file
		copyErr := files.CopyFile(src, dest)
		if copyErr != nil {
			msg := fmt.Sprint("😩 Calling t.FailNow(): failed copying from: ", src, " to: ", dest, " with error: ", copyErr)
			return errors.New(msg)
		} else {
			fmt.Println("✌️ Success! Copied from: ", src, " to: ", dest)
		}
	}

	return nil
}

// cleanupSupportingFiles deletes one or more files from a directory, intended to be called after copySupportingFiles
func cleanupSupportingFiles(fileNames []string, destination string) error {
	fmt.Println("Deleting files: ", fileNames, "from dir: ", destination)
	for _, file := range fileNames {
		fullPath := destination + "/" + file
		removeErr := os.Remove(fullPath)
		if removeErr != nil {
			fmt.Println("😩 Failed deleting file ", fullPath, " with error: ", removeErr)
			return removeErr
		} else {
			fmt.Println("✌️ Success! Deleted file: ", fullPath)
		}
	}
	return nil
}

// getGoogleCredentials reads a static service account credentials JSON file named gcp-creds.json from the test/ folder
//func getGoogleCredentials() string {
//	envGoogleCredentials, envPresent := os.LookupEnv("GOOGLE_CREDENTIALS")
//	if envPresent {
//		return envGoogleCredentials
//	}
//
//	fileGoogleCredentials, errReadingGCredsFromFile := os.ReadFile("gcp-creds.json")
//	if errReadingGCredsFromFile == nil {
//		return string(fileGoogleCredentials)
//	}
//	panic("No Google credentials available")
//}
