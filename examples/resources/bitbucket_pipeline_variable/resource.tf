data "bitbucket_workspace" "workspace" {
  name = "bitbucket-workspace"
}

resource "bitbucket_repository" "repo" {
  workspace   = data.bitbucket_workspace.workspace.id
  name        = "bitbucket-repository"
  description = "bitbucket-repository"
}

resource "bitbucket_pipeline_variable" "pipeline_variable" {
  workspace  = bitbucket_repository.repo.workspace
  repository = bitbucket_repository.repo.id
  key        = "hello"
  value      = "world"
  secured    = false
}
