package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccBitbucketRepo_Basic(t *testing.T) {
	repoSlug := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	repoSlugUpdated := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	description := "Repository for Terraform e2e tests"

	resource.Test(t, resource.TestCase{
		PreCheck:  PreCheck(t),
		Providers: Providers(),

		Steps: []resource.TestStep{
			{
				Config: testAccBitbucketRepositoryResource(repoSlug, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "description", description),
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "name", repoSlug),
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "workspace", bitbucketWorkspace),
				),
			},
			{
				Config: testAccBitbucketRepositoryResource(repoSlugUpdated, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "description", description),
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "name", repoSlugUpdated),
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "workspace", bitbucketWorkspace),
				),
			},
		},
	})
}

func testAccBitbucketRepositoryResource(name, description string) string {
	return fmt.Sprintf(`
resource "bitbucket_repository" "foo" {
  workspace   = "%s"
  name        = "%s"
  description = "%s"  
}`, bitbucketWorkspace, name, description)
}
