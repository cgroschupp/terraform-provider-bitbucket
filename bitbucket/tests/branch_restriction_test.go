package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccBitbucketBranchRestriction_Basic(t *testing.T) {
	// value := acctest.RandIntRange(1, 5)
	kind := "force"
	kindUpdate := "require_passing_builds_to_merge"

	resource.Test(t, resource.TestCase{
		PreCheck:  PreCheck(t),
		Providers: Providers(),
		Steps: []resource.TestStep{
			{

				Config: testAccPipelineBranchRestrictionResource(kind),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bitbucket_branch_restriction.foo", "kind", kind),
					resource.TestCheckResourceAttr("bitbucket_branch_restriction.foo", "workspace", bitbucketWorkspace),
				),
			},
			{
				Config: testAccPipelineBranchRestrictionValueResource(kindUpdate, 1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bitbucket_branch_restriction.foo", "kind", kindUpdate),
					resource.TestCheckResourceAttr("bitbucket_branch_restriction.foo", "workspace", bitbucketWorkspace),
					resource.TestCheckResourceAttr("bitbucket_branch_restriction.foo", "value", "1"),
				),
			},
		},
	})
}

func testAccPipelineBranchRestrictionResource(kind string) string {
	return fmt.Sprintf(`
resource "bitbucket_repository" "foo" {
  workspace   = "%s"
  name        = "unittest"
  description = "unittest"  
}
resource "bitbucket_branch_restriction" "foo" {
	workspace  = bitbucket_repository.foo.workspace
	repository = bitbucket_repository.foo.id
	kind        = "%s"
	pattern     = "master"
  }
`, bitbucketWorkspace, kind)
}

func testAccPipelineBranchRestrictionValueResource(kind string, value int) string {
	return fmt.Sprintf(`
resource "bitbucket_repository" "foo" {
  workspace   = "%s"
  name        = "unittest"
  description = "unittest"  
}
resource "bitbucket_branch_restriction" "foo" {
	workspace  = bitbucket_repository.foo.workspace
	repository = bitbucket_repository.foo.id
	kind        = "%s"
	pattern     = "master"
	value       = %d
  }
`, bitbucketWorkspace, kind, value)
}
