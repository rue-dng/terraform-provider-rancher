package rancher

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"rancher": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CATTLE_URL"); v == "" {
		t.Fatal("CATTLE_URL must be set for acceptance tests")
	}

	if v := os.Getenv("CATTLE_ACCESS_KEY"); v == "" {
		t.Fatal("CATTLE_ACCESS_KEY must be set for acceptance tests")
	}

	if v := os.Getenv("CATTLE_SECRET_KEY"); v == "" {
		t.Fatal("CATTLE_SECRET_KEY must be set for acceptance tests")
	}
}