data "bitbucket_workspace" "workspace" {
  name = "bitbucket-workspace"
}

resource "bitbucket_project" "project" {
  workspace   = data.bitbucket_workspace.workspace.id
  name        = "bitbucket-repository"
  key         = "test"
  description = "bitbucket-repository"
}
