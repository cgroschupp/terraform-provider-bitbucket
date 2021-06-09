data "bitbucket_workspace" "workspace" {
  name = "bitbucket-workspace"
}

data "bitbucket_repository" "repository" {
  repository = "bitbucket-repository"
  workspace  = data.bitbucket_workspace.workspace.name
}


resource "bitbucket_default_reviewers" "test" {
  workspace  = data.bitbucket_workspace.workspace.id
  repository = data.bitbucket_repository.repository.id
  user       = "bitbucket-user"
}
