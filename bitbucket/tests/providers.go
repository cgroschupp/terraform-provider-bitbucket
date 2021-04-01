package tests

import (
	"os"
	"testing"

	"github.com/cgroschupp/terraform-provider-bitbucket/bitbucket"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	bitbucketWorkspace = ""
)

func init() {
	bitbucketWorkspace = os.Getenv("BITBUCKET_WORKSPACE")
}

// Providers returns all providers used during acceptance testing.
func Providers() map[string]*schema.Provider {
	return map[string]*schema.Provider{
		"bitbucket": bitbucket.Provider(),
	}
}

// PreCheck checks if all conditions for an acceptance test are
// met.
func PreCheck(t *testing.T) func() {
	return func() {
		if v := os.Getenv("BITBUCKET_USERNAME"); v == "" {
			t.Fatal("BITBUCKET_USERNAME must be set for acceptance tests")
		}
		if v := os.Getenv("BITBUCKET_PASSWORD"); v == "" {
			t.Fatal("BITBUCKET_PASSWORD must be set for acceptance tests")
		}
		if v := os.Getenv("BITBUCKET_WORKSPACE"); v == "" {
			t.Fatal("BITBUCKET_WORKSPACE must be set for acceptance tests")
		}
	}
}
