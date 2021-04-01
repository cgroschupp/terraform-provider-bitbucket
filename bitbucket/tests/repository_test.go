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
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "repo_slug", repoSlug),
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "workspace", "cgroschupp"),
				),
			},
			{
				Config: testAccBitbucketRepositoryResource(repoSlugUpdated, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "description", description),
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "repo_slug", repoSlugUpdated),
					resource.TestCheckResourceAttr("bitbucket_repository.foo", "workspace", "cgroschupp"),
				),
			},
		},
	})
}

func testAccBitbucketRepositoryResource(name, description string) string {
	return fmt.Sprintf(`
resource "bitbucket_repository" "foo" {
  workspace   = "%s"
  repo_slug   = "%s"
  description = "%s"  
}`, bitbucketWorkspace, name, description)
}
