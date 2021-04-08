package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccBitbucketPipelineVariable_Basic(t *testing.T) {
	value := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	key := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	keyUpdate := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck:  PreCheck(t),
		Providers: Providers(),
		Steps: []resource.TestStep{
			{

				Config: testAccPipelineVariableResource(value, key),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bitbucket_pipeline_variable.foo", "key", key),
					resource.TestCheckResourceAttr("bitbucket_pipeline_variable.foo", "workspace", bitbucketWorkspace),
				),
			},
			{
				Config: testAccPipelineVariableResource(value, keyUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bitbucket_pipeline_variable.foo", "key", keyUpdate),
					resource.TestCheckResourceAttr("bitbucket_pipeline_variable.foo", "workspace", bitbucketWorkspace),
				),
			},
		},
	})
}

func testAccPipelineVariableResource(value, key string) string {
	return fmt.Sprintf(`
resource "bitbucket_repository" "foo" {
  workspace   = "%s"
  name        = "unittest"
  description = "unittest"  
}
resource "bitbucket_pipeline_variable" "foo" {
	workspace  = bitbucket_repository.foo.workspace
	repository = bitbucket_repository.foo.id
	key        = "%s"
	value      = "%s"
	secured    = false
  }
`, bitbucketWorkspace, key, value)
}
