// Code generated by go-github-cli/generator; DO NOT EDIT

package services

type TeamsCmd struct {
	List                     TeamsListCmd                     `cmd:"" help:"List teams"`
	Get                      TeamsGetCmd                      `cmd:"" help:"Get team"`
	Create                   TeamsCreateCmd                   `cmd:"" help:"Create team"`
	Edit                     TeamsEditCmd                     `cmd:"" help:"Edit team"`
	Delete                   TeamsDeleteCmd                   `cmd:"" help:"Delete team"`
	ListChild                TeamsListChildCmd                `cmd:"" help:"List child teams"`
	ListRepos                TeamsListReposCmd                `cmd:"" help:"List team repos"`
	CheckManagesRepo         TeamsCheckManagesRepoCmd         `cmd:"" help:"Check if a team manages a repository"`
	AddOrUpdateRepo          TeamsAddOrUpdateRepoCmd          `cmd:"" help:"Add or update team repository"`
	RemoveRepo               TeamsRemoveRepoCmd               `cmd:"" help:"Remove team repository"`
	ListForAuthenticatedUser TeamsListForAuthenticatedUserCmd `cmd:"" help:"List user teams"`
	ListProjects             TeamsListProjectsCmd             `cmd:"" help:"List team projects"`
	ReviewProject            TeamsReviewProjectCmd            `cmd:"" help:"Review a team project"`
	AddOrUpdateProject       TeamsAddOrUpdateProjectCmd       `cmd:"" help:"Add or update team project"`
	RemoveProject            TeamsRemoveProjectCmd            `cmd:"" help:"Remove team project"`
	ListDiscussions          TeamsListDiscussionsCmd          `cmd:"" help:"List discussions"`
	GetDiscussion            TeamsGetDiscussionCmd            `cmd:"" help:"Get a single discussion"`
	CreateDiscussion         TeamsCreateDiscussionCmd         `cmd:"" help:"Create a discussion"`
	EditDiscussion           TeamsEditDiscussionCmd           `cmd:"" help:"Edit a discussion"`
	DeleteDiscussion         TeamsDeleteDiscussionCmd         `cmd:"" help:"Delete a discussion"`
	ListDiscussionComments   TeamsListDiscussionCommentsCmd   `cmd:"" help:"List comments"`
	GetDiscussionComment     TeamsGetDiscussionCommentCmd     `cmd:"" help:"Get a single comment"`
	CreateDiscussionComment  TeamsCreateDiscussionCommentCmd  `cmd:"" help:"Create a comment"`
	EditDiscussionComment    TeamsEditDiscussionCommentCmd    `cmd:"" help:"Edit a comment"`
	DeleteDiscussionComment  TeamsDeleteDiscussionCommentCmd  `cmd:"" help:"Delete a comment"`
	ListMembers              TeamsListMembersCmd              `cmd:"" help:"List team members"`
	GetMember                TeamsGetMemberCmd                `cmd:"" help:"Get team member"`
	AddMember                TeamsAddMemberCmd                `cmd:"" help:"Add team member"`
	RemoveMember             TeamsRemoveMemberCmd             `cmd:"" help:"Remove team member"`
	GetMembership            TeamsGetMembershipCmd            `cmd:"" help:"Get team membership"`
	AddOrUpdateMembership    TeamsAddOrUpdateMembershipCmd    `cmd:"" help:"Add or update team membership"`
	RemoveMembership         TeamsRemoveMembershipCmd         `cmd:"" help:"Remove team membership"`
	ListPendingInvitations   TeamsListPendingInvitationsCmd   `cmd:"" help:"List pending team invitations"`
}

type TeamsListCmd struct {
	baseCmd
	Org     string `required:"" name:"org"`
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *TeamsListCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/orgs/:org/teams"
	c.updateURLPath("org", c.Org)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type TeamsGetCmd struct {
	baseCmd
	TeamId int64 `required:"" name:"team_id"`
}

func (c *TeamsGetCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id"
	c.updateURLPath("team_id", c.TeamId)
	return c.doRequest("GET")
}

type TeamsCreateCmd struct {
	baseCmd
	Org          string   `required:"" name:"org"`
	Name         string   `required:"" name:"name" help:"The name of the team."`
	Description  string   `name:"description" help:"The description of the team."`
	Maintainers  []string `name:"maintainers" help:"The logins of organization members to add as maintainers of the team."`
	RepoNames    []string `name:"repo_names" help:"The full name (e.g., 'organization-name/repository-name') of repositories to add the team to."`
	Privacy      string   "name:\"privacy\" help:\"The level of privacy this team should have. The options are:  \n**For a non-nested team:**  \n\\* `secret` - only visible to organization owners and members of this team.  \n\\* `closed` - visible to all members of this organization.  \nDefault: `secret`  \n**For a parent or child team:**  \n\\* `closed` - visible to all members of this organization.  \nDefault for child team: `closed`  \n**Note**: You must pass the `hellcat-preview` media type to set privacy default to `closed` for child teams. **For a parent or child team:**  \""
	Permission   string   "name:\"permission\" help:\"**Deprecated**. The permission that new repositories will be added to the team with when none is specified. Can be one of:  \n\\* `pull` - team members can pull, but not push to or administer newly-added repositories.  \n\\* `push` - team members can pull and push, but not administer newly-added repositories.  \n\\* `admin` - team members can pull, push and administer newly-added repositories.\""
	ParentTeamId int64    "name:\"parent_team_id\" help:\"The ID of a team to set as the parent team. **Note**: You must pass the `hellcat-preview` media type to use this parameter.\""
}

func (c *TeamsCreateCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/orgs/:org/teams"
	c.updateURLPath("org", c.Org)
	c.updateBody("name", c.Name)
	c.updateBody("description", c.Description)
	c.updateBody("maintainers", c.Maintainers)
	c.updateBody("repo_names", c.RepoNames)
	c.updateBody("privacy", c.Privacy)
	c.updateBody("permission", c.Permission)
	c.updateBody("parent_team_id", c.ParentTeamId)
	return c.doRequest("POST")
}

type TeamsEditCmd struct {
	baseCmd
	TeamId       int64  `required:"" name:"team_id"`
	Name         string `required:"" name:"name" help:"The name of the team."`
	Description  string `name:"description" help:"The description of the team."`
	Privacy      string "name:\"privacy\" help:\"The level of privacy this team should have. Editing teams without specifying this parameter leaves `privacy` intact. The options are:  \n**For a non-nested team:**  \n\\* `secret` - only visible to organization owners and members of this team.  \n\\* `closed` - visible to all members of this organization.  \n**For a parent or child team:**  \n\\* `closed` - visible to all members of this organization.\""
	Permission   string "name:\"permission\" help:\"**Deprecated**. The permission that new repositories will be added to the team with when none is specified. Can be one of:  \n\\* `pull` - team members can pull, but not push to or administer newly-added repositories.  \n\\* `push` - team members can pull and push, but not administer newly-added repositories.  \n\\* `admin` - team members can pull, push and administer newly-added repositories.\""
	ParentTeamId int64  "name:\"parent_team_id\" help:\"The ID of a team to set as the parent team. **Note**: You must pass the `hellcat-preview` media type to use this parameter.\""
}

func (c *TeamsEditCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id"
	c.updateURLPath("team_id", c.TeamId)
	c.updateBody("name", c.Name)
	c.updateBody("description", c.Description)
	c.updateBody("privacy", c.Privacy)
	c.updateBody("permission", c.Permission)
	c.updateBody("parent_team_id", c.ParentTeamId)
	return c.doRequest("PATCH")
}

type TeamsDeleteCmd struct {
	baseCmd
	TeamId int64 `required:"" name:"team_id"`
}

func (c *TeamsDeleteCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id"
	c.updateURLPath("team_id", c.TeamId)
	return c.doRequest("DELETE")
}

type TeamsListChildCmd struct {
	baseCmd
	TeamId  int64 `required:"" name:"team_id"`
	PerPage int64 `name:"per_page" help:"Results per page (max 100)"`
	Page    int64 `name:"page" help:"Page number of the results to fetch."`
}

func (c *TeamsListChildCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/teams"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type TeamsListReposCmd struct {
	baseCmd
	TeamId  int64 `required:"" name:"team_id"`
	PerPage int64 `name:"per_page" help:"Results per page (max 100)"`
	Page    int64 `name:"page" help:"Page number of the results to fetch."`
}

func (c *TeamsListReposCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/repos"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type TeamsCheckManagesRepoCmd struct {
	baseCmd
	TeamId int64  `required:"" name:"team_id"`
	Owner  string `required:"" name:"owner"`
	Repo   string `required:"" name:"repo"`
}

func (c *TeamsCheckManagesRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/repos/:owner/:repo"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	return c.doRequest("GET")
}

type TeamsAddOrUpdateRepoCmd struct {
	baseCmd
	TeamId     int64  `required:"" name:"team_id"`
	Owner      string `required:"" name:"owner"`
	Repo       string `required:"" name:"repo"`
	Permission string "name:\"permission\" help:\"The permission to grant the team on this repository. Can be one of:  \n\\* `pull` - team members can pull, but not push to or administer this repository.  \n\\* `push` - team members can pull and push, but not administer this repository.  \n\\* `admin` - team members can pull, push and administer this repository.  \n  \nIf no permission is specified, the team's `permission` attribute will be used to determine what permission to grant the team on this repository.  \n**Note**: If you pass the `hellcat-preview` media type, you can promote—but not demote—a `permission` attribute inherited through a parent team.\""
}

func (c *TeamsAddOrUpdateRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/repos/:owner/:repo"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	c.updateBody("permission", c.Permission)
	return c.doRequest("PUT")
}

type TeamsRemoveRepoCmd struct {
	baseCmd
	TeamId int64  `required:"" name:"team_id"`
	Owner  string `required:"" name:"owner"`
	Repo   string `required:"" name:"repo"`
}

func (c *TeamsRemoveRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/repos/:owner/:repo"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repo", c.Repo)
	return c.doRequest("DELETE")
}

type TeamsListForAuthenticatedUserCmd struct {
	baseCmd
	PerPage int64 `name:"per_page" help:"Results per page (max 100)"`
	Page    int64 `name:"page" help:"Page number of the results to fetch."`
}

func (c *TeamsListForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/user/teams"
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type TeamsListProjectsCmd struct {
	baseCmd
	TeamId  int64 `required:"" name:"team_id"`
	PerPage int64 `name:"per_page" help:"Results per page (max 100)"`
	Page    int64 `name:"page" help:"Page number of the results to fetch."`
}

func (c *TeamsListProjectsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/projects"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type TeamsReviewProjectCmd struct {
	baseCmd
	TeamId    int64 `required:"" name:"team_id"`
	ProjectId int64 `required:"" name:"project_id"`
}

func (c *TeamsReviewProjectCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/projects/:project_id"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("project_id", c.ProjectId)
	return c.doRequest("GET")
}

type TeamsAddOrUpdateProjectCmd struct {
	baseCmd
	TeamId     int64  `required:"" name:"team_id"`
	ProjectId  int64  `required:"" name:"project_id"`
	Permission string "name:\"permission\" help:\"The permission to grant to the team for this project. Can be one of:  \n\\* `read` - team members can read, but not write to or administer this project.  \n\\* `write` - team members can read and write, but not administer this project.  \n\\* `admin` - team members can read, write and administer this project.  \nDefault: the team's `permission` attribute will be used to determine what permission to grant the team on this project. Note that, if you choose not to pass any parameters, you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see '[HTTP verbs](https://developer.github.com/v3/#http-verbs).'  \n**Note**: If you pass the `hellcat-preview` media type, you can promote—but not demote—a `permission` attribute inherited from a parent team.\""
}

func (c *TeamsAddOrUpdateProjectCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/projects/:project_id"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("project_id", c.ProjectId)
	c.updateBody("permission", c.Permission)
	return c.doRequest("PUT")
}

type TeamsRemoveProjectCmd struct {
	baseCmd
	TeamId    int64 `required:"" name:"team_id"`
	ProjectId int64 `required:"" name:"project_id"`
}

func (c *TeamsRemoveProjectCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/projects/:project_id"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("project_id", c.ProjectId)
	return c.doRequest("DELETE")
}

type TeamsListDiscussionsCmd struct {
	baseCmd
	TeamId    int64  `required:"" name:"team_id"`
	Direction string "name:\"direction\" help:\"Sorts the discussion comments by the date they were created. To return the oldest comments first, set to `asc`. Can be one of `asc` or `desc`.\""
	PerPage   int64  `name:"per_page" help:"Results per page (max 100)"`
	Page      int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *TeamsListDiscussionsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLQuery("direction", c.Direction)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type TeamsGetDiscussionCmd struct {
	baseCmd
	TeamId           int64 `required:"" name:"team_id"`
	DiscussionNumber int64 `required:"" name:"discussion_number"`
}

func (c *TeamsGetDiscussionCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions/:discussion_number"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("discussion_number", c.DiscussionNumber)
	return c.doRequest("GET")
}

type TeamsCreateDiscussionCmd struct {
	baseCmd
	TeamId  int64  `required:"" name:"team_id"`
	Title   string `required:"" name:"title" help:"The discussion post's title."`
	Body    string `required:"" name:"body" help:"The discussion post's body text."`
	Private bool   "name:\"private\" help:\"Private posts are only visible to team members, organization owners, and team maintainers. Public posts are visible to all members of the organization. Set to `true` to create a private post.\""
}

func (c *TeamsCreateDiscussionCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions"
	c.updateURLPath("team_id", c.TeamId)
	c.updateBody("title", c.Title)
	c.updateBody("body", c.Body)
	c.updateBody("private", c.Private)
	return c.doRequest("POST")
}

type TeamsEditDiscussionCmd struct {
	baseCmd
	TeamId           int64  `required:"" name:"team_id"`
	DiscussionNumber int64  `required:"" name:"discussion_number"`
	Title            string `name:"title" help:"The discussion post's title."`
	Body             string `name:"body" help:"The discussion post's body text."`
}

func (c *TeamsEditDiscussionCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions/:discussion_number"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("discussion_number", c.DiscussionNumber)
	c.updateBody("title", c.Title)
	c.updateBody("body", c.Body)
	return c.doRequest("PATCH")
}

type TeamsDeleteDiscussionCmd struct {
	baseCmd
	TeamId           int64 `required:"" name:"team_id"`
	DiscussionNumber int64 `required:"" name:"discussion_number"`
}

func (c *TeamsDeleteDiscussionCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions/:discussion_number"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("discussion_number", c.DiscussionNumber)
	return c.doRequest("DELETE")
}

type TeamsListDiscussionCommentsCmd struct {
	baseCmd
	TeamId           int64  `required:"" name:"team_id"`
	DiscussionNumber int64  `required:"" name:"discussion_number"`
	Direction        string "name:\"direction\" help:\"Sorts the discussion comments by the date they were created. To return the oldest comments first, set to `asc`. Can be one of `asc` or `desc`.\""
	PerPage          int64  `name:"per_page" help:"Results per page (max 100)"`
	Page             int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *TeamsListDiscussionCommentsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions/:discussion_number/comments"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("discussion_number", c.DiscussionNumber)
	c.updateURLQuery("direction", c.Direction)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type TeamsGetDiscussionCommentCmd struct {
	baseCmd
	TeamId           int64 `required:"" name:"team_id"`
	DiscussionNumber int64 `required:"" name:"discussion_number"`
	CommentNumber    int64 `required:"" name:"comment_number"`
}

func (c *TeamsGetDiscussionCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions/:discussion_number/comments/:comment_number"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("discussion_number", c.DiscussionNumber)
	c.updateURLPath("comment_number", c.CommentNumber)
	return c.doRequest("GET")
}

type TeamsCreateDiscussionCommentCmd struct {
	baseCmd
	TeamId           int64  `required:"" name:"team_id"`
	DiscussionNumber int64  `required:"" name:"discussion_number"`
	Body             string `required:"" name:"body" help:"The discussion comment's body text."`
}

func (c *TeamsCreateDiscussionCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions/:discussion_number/comments"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("discussion_number", c.DiscussionNumber)
	c.updateBody("body", c.Body)
	return c.doRequest("POST")
}

type TeamsEditDiscussionCommentCmd struct {
	baseCmd
	TeamId           int64  `required:"" name:"team_id"`
	DiscussionNumber int64  `required:"" name:"discussion_number"`
	CommentNumber    int64  `required:"" name:"comment_number"`
	Body             string `required:"" name:"body" help:"The discussion comment's body text."`
}

func (c *TeamsEditDiscussionCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions/:discussion_number/comments/:comment_number"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("discussion_number", c.DiscussionNumber)
	c.updateURLPath("comment_number", c.CommentNumber)
	c.updateBody("body", c.Body)
	return c.doRequest("PATCH")
}

type TeamsDeleteDiscussionCommentCmd struct {
	baseCmd
	TeamId           int64 `required:"" name:"team_id"`
	DiscussionNumber int64 `required:"" name:"discussion_number"`
	CommentNumber    int64 `required:"" name:"comment_number"`
}

func (c *TeamsDeleteDiscussionCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/discussions/:discussion_number/comments/:comment_number"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("discussion_number", c.DiscussionNumber)
	c.updateURLPath("comment_number", c.CommentNumber)
	return c.doRequest("DELETE")
}

type TeamsListMembersCmd struct {
	baseCmd
	TeamId  int64  `required:"" name:"team_id"`
	Role    string "name:\"role\" help:\"Filters members returned by their role in the team. Can be one of:  \n\\* `member` - normal members of the team.  \n\\* `maintainer` - team maintainers.  \n\\* `all` - all members of the team.\""
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *TeamsListMembersCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/members"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLQuery("role", c.Role)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type TeamsGetMemberCmd struct {
	baseCmd
	TeamId   int64  `required:"" name:"team_id"`
	Username string `required:"" name:"username"`
}

func (c *TeamsGetMemberCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/members/:username"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("username", c.Username)
	return c.doRequest("GET")
}

type TeamsAddMemberCmd struct {
	baseCmd
	TeamId   int64  `required:"" name:"team_id"`
	Username string `required:"" name:"username"`
}

func (c *TeamsAddMemberCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/members/:username"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("username", c.Username)
	return c.doRequest("PUT")
}

type TeamsRemoveMemberCmd struct {
	baseCmd
	TeamId   int64  `required:"" name:"team_id"`
	Username string `required:"" name:"username"`
}

func (c *TeamsRemoveMemberCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/members/:username"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("username", c.Username)
	return c.doRequest("DELETE")
}

type TeamsGetMembershipCmd struct {
	baseCmd
	TeamId   int64  `required:"" name:"team_id"`
	Username string `required:"" name:"username"`
}

func (c *TeamsGetMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/memberships/:username"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("username", c.Username)
	return c.doRequest("GET")
}

type TeamsAddOrUpdateMembershipCmd struct {
	baseCmd
	TeamId   int64  `required:"" name:"team_id"`
	Username string `required:"" name:"username"`
	Role     string "name:\"role\" help:\"The role that this user should have in the team. Can be one of:  \n\\* `member` - a normal member of the team.  \n\\* `maintainer` - a team maintainer. Able to add/remove other team members, promote other team members to team maintainer, and edit the team's name and description.\""
}

func (c *TeamsAddOrUpdateMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/memberships/:username"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("username", c.Username)
	c.updateBody("role", c.Role)
	return c.doRequest("PUT")
}

type TeamsRemoveMembershipCmd struct {
	baseCmd
	TeamId   int64  `required:"" name:"team_id"`
	Username string `required:"" name:"username"`
}

func (c *TeamsRemoveMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/memberships/:username"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLPath("username", c.Username)
	return c.doRequest("DELETE")
}

type TeamsListPendingInvitationsCmd struct {
	baseCmd
	TeamId  int64 `required:"" name:"team_id"`
	PerPage int64 `name:"per_page" help:"Results per page (max 100)"`
	Page    int64 `name:"page" help:"Page number of the results to fetch."`
}

func (c *TeamsListPendingInvitationsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/teams/:team_id/invitations"
	c.updateURLPath("team_id", c.TeamId)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}