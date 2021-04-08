data "bitbucket_workspace" "workspace" {
  name = "bitbucket-workspace"
}

resource "bitbucket_repository" "repo" {
  workspace   = data.bitbucket_workspace.workspace.id
  name        = "bitbucket-repository"
  description = "bitbucket-repository"
}

resource "bitbucket_branch_restriction" "branch_restriction" {
  workspace  = bitbucket_repository.repo.workspace
  repository = bitbucket_repository.repo.id
  kind       = "force"
  pattern    = "master"
}
