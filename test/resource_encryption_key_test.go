package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
)

func TestAccEBSKey(t *testing.T) {
	testAccEBSKeyProvider := instaclustr.Provider()
	testAccEBSKeyProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccEBSKeyProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid_encryption_key.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
<<<<<<< HEAD
	kmsAlias := os.Getenv("KMS_ALIAS")
	kmsArn := os.Getenv("KMS_ARN")
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, kmsAlias, kmsArn)
	hostname := instaclustr.ApiHostname
	resource.Test(t, resource.TestCase{
		Providers:    testAccEBSKeyProviders,
		PreCheck:     func() { testEnvPreCheck(t) },
=======
	kmsArn := os.Getenv("KMS_ARN")
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, kmsArn)
	hostname := instaclustr.ApiHostname
	resource.Test(t, resource.TestCase{
		Providers:    testAccEBSKeyProviders,
		PreCheck:     func() { AccTestEnvVarsCheck(t) },
>>>>>>> e2a4bb19800c323c205c06f49c37775b3319210e
		CheckDestroy: testCheckAccEBSResourceDeleted("valid", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAccEBSResourceValid("valid"),
					testCheckAccEBSResourceCreated("valid", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestAccEBSKeyInvalid(t *testing.T) {
	testAccEBSKeyProvider := instaclustr.Provider()
	testAccEBSKeyProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccEBSKeyProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/invalid_encryption_key.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
<<<<<<< HEAD
	kmsAlias := os.Getenv("KMS_ALIAS")
	kmsArn := os.Getenv("KMS_ARN")
	resource.Test(t, resource.TestCase{
		Providers: testAccEBSKeyProviders,
		PreCheck:  func() { testEnvPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, apiKey, kmsAlias, kmsArn),
=======
	kmsArn := os.Getenv("KMS_ARN")
	resource.Test(t, resource.TestCase{
		Providers: testAccEBSKeyProviders,
		PreCheck:  func() { AccTestEnvVarsCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, apiKey, kmsArn),
>>>>>>> e2a4bb19800c323c205c06f49c37775b3319210e
				ExpectError: regexp.MustCompile("Error adding encryption key"),
			},
		},
	})
}

<<<<<<< HEAD
func testEnvPreCheck(t *testing.T) {
	if v := os.Getenv("IC_USERNAME"); v == "" {
		t.Fatal("IC_USERNAME for provisioning API must be set for acceptance tests")
	}
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY for provisioning API must be set for acceptance tests")
	}
	if v := os.Getenv("KMS_ALIAS"); v == "" {
		t.Fatal("KMS_ALIAS for AccEBS encryption must be set for acceptance tests")
	}
	if v := os.Getenv("KMS_ARN"); v == "" {
		t.Fatal("KMS_ARN for AccEBS encryption must be set for acceptance tests")
	}
}

=======
>>>>>>> e2a4bb19800c323c205c06f49c37775b3319210e
func testCheckAccEBSResourceValid(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_encryption_key."+resourceName]
		if resourceState == nil {
			return fmt.Errorf("%s: resource not found in state", resourceName)
		}

		instanceState := resourceState.Primary
		if instanceState == nil {
			return fmt.Errorf("resource has no primary instance")
		}
		return nil
	}
}

func testCheckAccEBSResourceCreated(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_encryption_key."+resourceName]
		id := resourceState.Primary.Attributes["key_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		resource, err := client.ReadEncryptionKey(id)
		if err != nil {
			return fmt.Errorf("Failed to read encryption key %s: %s", id, err)
		}
		if resource.ID != id {
			return fmt.Errorf("Encryption key expected %s but got %s", id, resource.ID)
		}
		return nil
	}
}

func testCheckAccEBSResourceDeleted(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_encryption_key."+resourceName]
		id := resourceState.Primary.Attributes["key_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		err := client.DeleteEncryptionKey(id)
		if err == nil {
			return fmt.Errorf("Encryption key %s still exists", id)
		}
		return nil
	}
}
