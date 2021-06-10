data "bitbucket_workspace" "workspace" {
  name = "bitbucket-workspace"
}

data "bitbucket_repository" "repository" {
  repository = "bitbucket-repository"
  workspace  = data.bitbucket_workspace.workspace.name
}

resource "bitbucket_group_permission" "permission" {
  workspace  = data.bitbucket_workspace.workspace.id
  repository = data.bitbucket_repository.repository.id
  group      = "bitbucket-group"
  permission = "read"
}
