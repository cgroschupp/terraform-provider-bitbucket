data "bitbucket_workspace" "workspace" {
  name = "bitbucket-workspace"
}

resource "bitbucket_repository" "repo" {
  workspace   = data.bitbucket_workspace.workspace.id
  name        = "bitbucket-repository"
  description = "bitbucket-repository"
}