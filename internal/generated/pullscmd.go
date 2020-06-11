// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import internal "github.com/octo-cli/octo-cli/internal"

type PullsCmd struct {
	CheckIfMerged               PullsCheckIfMergedCmd               `cmd:""`
	Create                      PullsCreateCmd                      `cmd:""`
	CreateReplyForReviewComment PullsCreateReplyForReviewCommentCmd `cmd:""`
	CreateReview                PullsCreateReviewCmd                `cmd:""`
	CreateReviewComment         PullsCreateReviewCommentCmd         `cmd:""`
	DeletePendingReview         PullsDeletePendingReviewCmd         `cmd:""`
	DeleteReviewComment         PullsDeleteReviewCommentCmd         `cmd:""`
	DismissReview               PullsDismissReviewCmd               `cmd:""`
	Get                         PullsGetCmd                         `cmd:""`
	GetReview                   PullsGetReviewCmd                   `cmd:""`
	GetReviewComment            PullsGetReviewCommentCmd            `cmd:""`
	List                        PullsListCmd                        `cmd:""`
	ListCommentsForReview       PullsListCommentsForReviewCmd       `cmd:""`
	ListCommits                 PullsListCommitsCmd                 `cmd:""`
	ListFiles                   PullsListFilesCmd                   `cmd:""`
	ListRequestedReviewers      PullsListRequestedReviewersCmd      `cmd:""`
	ListReviewComments          PullsListReviewCommentsCmd          `cmd:""`
	ListReviewCommentsForRepo   PullsListReviewCommentsForRepoCmd   `cmd:""`
	ListReviews                 PullsListReviewsCmd                 `cmd:""`
	Merge                       PullsMergeCmd                       `cmd:""`
	RemoveRequestedReviewers    PullsRemoveRequestedReviewersCmd    `cmd:""`
	RequestReviewers            PullsRequestReviewersCmd            `cmd:""`
	SubmitReview                PullsSubmitReviewCmd                `cmd:""`
	Update                      PullsUpdateCmd                      `cmd:""`
	UpdateBranch                PullsUpdateBranchCmd                `cmd:""`
	UpdateReview                PullsUpdateReviewCmd                `cmd:""`
	UpdateReviewComment         PullsUpdateReviewCommentCmd         `cmd:""`
}

type PullsCheckIfMergedCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	internal.BaseCmd
}

func (c *PullsCheckIfMergedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/merge")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	return c.DoRequest("GET")
}

type PullsCreateCmd struct {
	SailorV             bool   `name:"sailor-v-preview"`
	Repo                string `name:"repo" required:"true"`
	Body                string `name:"body"`
	Draft               bool   `name:"draft"`
	MaintainerCanModify bool   `name:"maintainer_can_modify"`
	Base                string `name:"base" required:"true"`
	Head                string `name:"head" required:"true"`
	Title               string `name:"title" required:"true"`
	internal.BaseCmd
}

func (c *PullsCreateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdatePreview("sailor-v", c.SailorV)
	c.UpdateBody("base", c.Base)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("draft", c.Draft)
	c.UpdateBody("head", c.Head)
	c.UpdateBody("maintainer_can_modify", c.MaintainerCanModify)
	c.UpdateBody("title", c.Title)
	return c.DoRequest("POST")
}

type PullsCreateReplyForReviewCommentCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	CommentId  int64  `name:"comment_id" required:"true"`
	Body       string `name:"body" required:"true"`
	internal.BaseCmd
}

func (c *PullsCreateReplyForReviewCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/comments/{comment_id}/replies")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("POST")
}

type PullsCreateReviewCmd struct {
	Repo       string                `name:"repo" required:"true"`
	PullNumber int64                 `name:"pull_number" required:"true"`
	Body       string                `name:"body"`
	Comments   []internal.JSONObject `name:"comments"`
	CommitId   string                `name:"commit_id"`
	Event      string                `name:"event"`
	internal.BaseCmd
}

func (c *PullsCreateReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/reviews")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("comments", c.Comments)
	c.UpdateBody("commit_id", c.CommitId)
	c.UpdateBody("event", c.Event)
	return c.DoRequest("POST")
}

type PullsCreateReviewCommentCmd struct {
	ComfortFade bool   `name:"comfort-fade-preview"`
	Repo        string `name:"repo" required:"true"`
	PullNumber  int64  `name:"pull_number" required:"true"`
	Line        int64  `name:"line"`
	Position    int64  `name:"position"`
	Side        string `name:"side"`
	StartLine   int64  `name:"start_line"`
	StartSide   string `name:"start_side"`
	Body        string `name:"body" required:"true"`
	CommitId    string `name:"commit_id" required:"true"`
	Path        string `name:"path" required:"true"`
	internal.BaseCmd
}

func (c *PullsCreateReviewCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/comments")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdatePreview("comfort-fade", c.ComfortFade)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("commit_id", c.CommitId)
	c.UpdateBody("line", c.Line)
	c.UpdateBody("path", c.Path)
	c.UpdateBody("position", c.Position)
	c.UpdateBody("side", c.Side)
	c.UpdateBody("start_line", c.StartLine)
	c.UpdateBody("start_side", c.StartSide)
	return c.DoRequest("POST")
}

type PullsDeletePendingReviewCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	ReviewId   int64  `name:"review_id" required:"true"`
	internal.BaseCmd
}

func (c *PullsDeletePendingReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/reviews/{review_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLPath("review_id", c.ReviewId)
	return c.DoRequest("DELETE")
}

type PullsDeleteReviewCommentCmd struct {
	Repo      string `name:"repo" required:"true"`
	CommentId int64  `name:"comment_id" required:"true"`
	internal.BaseCmd
}

func (c *PullsDeleteReviewCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/comments/{comment_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	return c.DoRequest("DELETE")
}

type PullsDismissReviewCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	ReviewId   int64  `name:"review_id" required:"true"`
	Message    string `name:"message" required:"true"`
	internal.BaseCmd
}

func (c *PullsDismissReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/reviews/{review_id}/dismissals")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLPath("review_id", c.ReviewId)
	c.UpdateBody("message", c.Message)
	return c.DoRequest("PUT")
}

type PullsGetCmd struct {
	SailorV    bool   `name:"sailor-v-preview"`
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	internal.BaseCmd
}

func (c *PullsGetCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdatePreview("sailor-v", c.SailorV)
	return c.DoRequest("GET")
}

type PullsGetReviewCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	ReviewId   int64  `name:"review_id" required:"true"`
	internal.BaseCmd
}

func (c *PullsGetReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/reviews/{review_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLPath("review_id", c.ReviewId)
	return c.DoRequest("GET")
}

type PullsGetReviewCommentCmd struct {
	ComfortFade  bool   `name:"comfort-fade-preview"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	Repo         string `name:"repo" required:"true"`
	CommentId    int64  `name:"comment_id" required:"true"`
	internal.BaseCmd
}

func (c *PullsGetReviewCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/comments/{comment_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdatePreview("comfort-fade", c.ComfortFade)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type PullsListCmd struct {
	SailorV   bool   `name:"sailor-v-preview"`
	Repo      string `name:"repo" required:"true"`
	Base      string `name:"base"`
	Direction string `name:"direction"`
	Head      string `name:"head"`
	Page      int64  `name:"page"`
	PerPage   int64  `name:"per_page"`
	Sort      string `name:"sort"`
	State     string `name:"state"`
	internal.BaseCmd
}

func (c *PullsListCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("state", c.State)
	c.UpdateURLQuery("head", c.Head)
	c.UpdateURLQuery("base", c.Base)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("sailor-v", c.SailorV)
	return c.DoRequest("GET")
}

type PullsListCommentsForReviewCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	ReviewId   int64  `name:"review_id" required:"true"`
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *PullsListCommentsForReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/reviews/{review_id}/comments")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLPath("review_id", c.ReviewId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsListCommitsCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *PullsListCommitsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/commits")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsListFilesCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *PullsListFilesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/files")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsListRequestedReviewersCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *PullsListRequestedReviewersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/requested_reviewers")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsListReviewCommentsCmd struct {
	ComfortFade  bool   `name:"comfort-fade-preview"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	Repo         string `name:"repo" required:"true"`
	PullNumber   int64  `name:"pull_number" required:"true"`
	Direction    string `name:"direction"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	Since        string `name:"since"`
	Sort         string `name:"sort"`
	internal.BaseCmd
}

func (c *PullsListReviewCommentsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/comments")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("comfort-fade", c.ComfortFade)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type PullsListReviewCommentsForRepoCmd struct {
	ComfortFade  bool   `name:"comfort-fade-preview"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	Repo         string `name:"repo" required:"true"`
	Direction    string `name:"direction"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	Since        string `name:"since"`
	Sort         string `name:"sort"`
	internal.BaseCmd
}

func (c *PullsListReviewCommentsForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/comments")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("comfort-fade", c.ComfortFade)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type PullsListReviewsCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *PullsListReviewsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/reviews")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsMergeCmd struct {
	Repo          string `name:"repo" required:"true"`
	PullNumber    int64  `name:"pull_number" required:"true"`
	CommitMessage string `name:"commit_message"`
	CommitTitle   string `name:"commit_title"`
	MergeMethod   string `name:"merge_method"`
	Sha           string `name:"sha"`
	internal.BaseCmd
}

func (c *PullsMergeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/merge")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateBody("commit_message", c.CommitMessage)
	c.UpdateBody("commit_title", c.CommitTitle)
	c.UpdateBody("merge_method", c.MergeMethod)
	c.UpdateBody("sha", c.Sha)
	return c.DoRequest("PUT")
}

type PullsRemoveRequestedReviewersCmd struct {
	Repo          string   `name:"repo" required:"true"`
	PullNumber    int64    `name:"pull_number" required:"true"`
	Reviewers     []string `name:"reviewers"`
	TeamReviewers []string `name:"team_reviewers"`
	internal.BaseCmd
}

func (c *PullsRemoveRequestedReviewersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/requested_reviewers")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateBody("reviewers", c.Reviewers)
	c.UpdateBody("team_reviewers", c.TeamReviewers)
	return c.DoRequest("DELETE")
}

type PullsRequestReviewersCmd struct {
	Repo          string   `name:"repo" required:"true"`
	PullNumber    int64    `name:"pull_number" required:"true"`
	Reviewers     []string `name:"reviewers"`
	TeamReviewers []string `name:"team_reviewers"`
	internal.BaseCmd
}

func (c *PullsRequestReviewersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/requested_reviewers")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateBody("reviewers", c.Reviewers)
	c.UpdateBody("team_reviewers", c.TeamReviewers)
	return c.DoRequest("POST")
}

type PullsSubmitReviewCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	ReviewId   int64  `name:"review_id" required:"true"`
	Body       string `name:"body"`
	Event      string `name:"event" required:"true"`
	internal.BaseCmd
}

func (c *PullsSubmitReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/reviews/{review_id}/events")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLPath("review_id", c.ReviewId)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("event", c.Event)
	return c.DoRequest("POST")
}

type PullsUpdateBranchCmd struct {
	Lydian          bool   `name:"lydian-preview" required:"true"`
	Repo            string `name:"repo" required:"true"`
	PullNumber      int64  `name:"pull_number" required:"true"`
	ExpectedHeadSha string `name:"expected_head_sha"`
	internal.BaseCmd
}

func (c *PullsUpdateBranchCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/update-branch")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdatePreview("lydian", c.Lydian)
	c.UpdateBody("expected_head_sha", c.ExpectedHeadSha)
	return c.DoRequest("PUT")
}

type PullsUpdateCmd struct {
	SailorV             bool   `name:"sailor-v-preview"`
	Repo                string `name:"repo" required:"true"`
	PullNumber          int64  `name:"pull_number" required:"true"`
	Base                string `name:"base"`
	Body                string `name:"body"`
	MaintainerCanModify bool   `name:"maintainer_can_modify"`
	State               string `name:"state"`
	Title               string `name:"title"`
	internal.BaseCmd
}

func (c *PullsUpdateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdatePreview("sailor-v", c.SailorV)
	c.UpdateBody("base", c.Base)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("maintainer_can_modify", c.MaintainerCanModify)
	c.UpdateBody("state", c.State)
	c.UpdateBody("title", c.Title)
	return c.DoRequest("PATCH")
}

type PullsUpdateReviewCmd struct {
	Repo       string `name:"repo" required:"true"`
	PullNumber int64  `name:"pull_number" required:"true"`
	ReviewId   int64  `name:"review_id" required:"true"`
	Body       string `name:"body" required:"true"`
	internal.BaseCmd
}

func (c *PullsUpdateReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/reviews/{review_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLPath("review_id", c.ReviewId)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("PUT")
}

type PullsUpdateReviewCommentCmd struct {
	ComfortFade bool   `name:"comfort-fade-preview"`
	Repo        string `name:"repo" required:"true"`
	CommentId   int64  `name:"comment_id" required:"true"`
	Body        string `name:"body" required:"true"`
	internal.BaseCmd
}

func (c *PullsUpdateReviewCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/comments/{comment_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdatePreview("comfort-fade", c.ComfortFade)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("PATCH")
}
