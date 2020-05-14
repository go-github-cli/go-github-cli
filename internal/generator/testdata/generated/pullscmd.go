// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import (
	"github.com/octo-cli/octo-cli/internal"
)

type PullsCmd struct {
	CheckIfMerged            PullsCheckIfMergedCmd            `cmd:""`
	Create                   PullsCreateCmd                   `cmd:""`
	CreateComment            PullsCreateCommentCmd            `cmd:""`
	CreateReview             PullsCreateReviewCmd             `cmd:""`
	CreateReviewCommentReply PullsCreateReviewCommentReplyCmd `cmd:""`
	CreateReviewRequest      PullsCreateReviewRequestCmd      `cmd:""`
	DeleteComment            PullsDeleteCommentCmd            `cmd:""`
	DeletePendingReview      PullsDeletePendingReviewCmd      `cmd:""`
	DeleteReviewRequest      PullsDeleteReviewRequestCmd      `cmd:""`
	DismissReview            PullsDismissReviewCmd            `cmd:""`
	Get                      PullsGetCmd                      `cmd:""`
	GetComment               PullsGetCommentCmd               `cmd:""`
	GetCommentsForReview     PullsGetCommentsForReviewCmd     `cmd:""`
	GetReview                PullsGetReviewCmd                `cmd:""`
	List                     PullsListCmd                     `cmd:""`
	ListComments             PullsListCommentsCmd             `cmd:""`
	ListCommentsForRepo      PullsListCommentsForRepoCmd      `cmd:""`
	ListCommits              PullsListCommitsCmd              `cmd:""`
	ListFiles                PullsListFilesCmd                `cmd:""`
	ListReviewRequests       PullsListReviewRequestsCmd       `cmd:""`
	ListReviews              PullsListReviewsCmd              `cmd:""`
	Merge                    PullsMergeCmd                    `cmd:""`
	SubmitReview             PullsSubmitReviewCmd             `cmd:""`
	Update                   PullsUpdateCmd                   `cmd:""`
	UpdateBranch             PullsUpdateBranchCmd             `cmd:""`
	UpdateComment            PullsUpdateCommentCmd            `cmd:""`
	UpdateReview             PullsUpdateReviewCmd             `cmd:""`
}

type PullsCheckIfMergedCmd struct {
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
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
	Base                string `required:"" name:"base"`
	Body                string `name:"body"`
	Draft               bool   `name:"draft"`
	Head                string `required:"" name:"head"`
	MaintainerCanModify bool   `name:"maintainer_can_modify"`
	Repo                string `required:"" name:"repo"`
	SailorV             bool   `name:"sailor-v-preview"`
	Title               string `required:"" name:"title"`
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

type PullsCreateCommentCmd struct {
	Body        string `required:"" name:"body"`
	ComfortFade bool   `name:"comfort-fade-preview"`
	CommitId    string `required:"" name:"commit_id"`
	Line        int64  `name:"line"`
	Path        string `required:"" name:"path"`
	Position    int64  `name:"position"`
	PullNumber  int64  `required:"" name:"pull_number"`
	Repo        string `required:"" name:"repo"`
	Side        string `name:"side"`
	StartLine   int64  `name:"start_line"`
	StartSide   string `name:"start_side"`
	internal.BaseCmd
}

func (c *PullsCreateCommentCmd) Run(isValueSetMap map[string]bool) error {
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

type PullsCreateReviewCmd struct {
	Body       string                `name:"body"`
	Comments   []internal.JSONObject `name:"comments"`
	CommitId   string                `name:"commit_id"`
	Event      string                `name:"event"`
	PullNumber int64                 `required:"" name:"pull_number"`
	Repo       string                `required:"" name:"repo"`
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

type PullsCreateReviewCommentReplyCmd struct {
	Body       string `required:"" name:"body"`
	CommentId  int64  `required:"" name:"comment_id"`
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *PullsCreateReviewCommentReplyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/comments/{comment_id}/replies")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("POST")
}

type PullsCreateReviewRequestCmd struct {
	PullNumber    int64    `required:"" name:"pull_number"`
	Repo          string   `required:"" name:"repo"`
	Reviewers     []string `name:"reviewers"`
	TeamReviewers []string `name:"team_reviewers"`
	internal.BaseCmd
}

func (c *PullsCreateReviewRequestCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/requested_reviewers")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateBody("reviewers", c.Reviewers)
	c.UpdateBody("team_reviewers", c.TeamReviewers)
	return c.DoRequest("POST")
}

type PullsDeleteCommentCmd struct {
	CommentId int64  `required:"" name:"comment_id"`
	Repo      string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *PullsDeleteCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/comments/{comment_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	return c.DoRequest("DELETE")
}

type PullsDeletePendingReviewCmd struct {
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
	ReviewId   int64  `required:"" name:"review_id"`
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

type PullsDeleteReviewRequestCmd struct {
	PullNumber    int64    `required:"" name:"pull_number"`
	Repo          string   `required:"" name:"repo"`
	Reviewers     []string `name:"reviewers"`
	TeamReviewers []string `name:"team_reviewers"`
	internal.BaseCmd
}

func (c *PullsDeleteReviewRequestCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/requested_reviewers")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateBody("reviewers", c.Reviewers)
	c.UpdateBody("team_reviewers", c.TeamReviewers)
	return c.DoRequest("DELETE")
}

type PullsDismissReviewCmd struct {
	Message    string `required:"" name:"message"`
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
	ReviewId   int64  `required:"" name:"review_id"`
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
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
	SailorV    bool   `name:"sailor-v-preview"`
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

type PullsGetCommentCmd struct {
	ComfortFade  bool   `name:"comfort-fade-preview"`
	CommentId    int64  `required:"" name:"comment_id"`
	Repo         string `required:"" name:"repo"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	internal.BaseCmd
}

func (c *PullsGetCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/comments/{comment_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdatePreview("comfort-fade", c.ComfortFade)
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	return c.DoRequest("GET")
}

type PullsGetCommentsForReviewCmd struct {
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
	ReviewId   int64  `required:"" name:"review_id"`
	internal.BaseCmd
}

func (c *PullsGetCommentsForReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/reviews/{review_id}/comments")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLPath("review_id", c.ReviewId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsGetReviewCmd struct {
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
	ReviewId   int64  `required:"" name:"review_id"`
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

type PullsListCmd struct {
	Base      string `name:"base"`
	Direction string `name:"direction"`
	Head      string `name:"head"`
	Page      int64  `name:"page"`
	PerPage   int64  `name:"per_page"`
	Repo      string `required:"" name:"repo"`
	SailorV   bool   `name:"sailor-v-preview"`
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

type PullsListCommentsCmd struct {
	ComfortFade  bool   `name:"comfort-fade-preview"`
	Direction    string `name:"direction"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	PullNumber   int64  `required:"" name:"pull_number"`
	Repo         string `required:"" name:"repo"`
	Since        string `name:"since"`
	Sort         string `name:"sort"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	internal.BaseCmd
}

func (c *PullsListCommentsCmd) Run(isValueSetMap map[string]bool) error {
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

type PullsListCommentsForRepoCmd struct {
	ComfortFade  bool   `name:"comfort-fade-preview"`
	Direction    string `name:"direction"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	Repo         string `required:"" name:"repo"`
	Since        string `name:"since"`
	Sort         string `name:"sort"`
	SquirrelGirl bool   `name:"squirrel-girl-preview"`
	internal.BaseCmd
}

func (c *PullsListCommentsForRepoCmd) Run(isValueSetMap map[string]bool) error {
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

type PullsListCommitsCmd struct {
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
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
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
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

type PullsListReviewRequestsCmd struct {
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *PullsListReviewRequestsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/{pull_number}/requested_reviewers")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("pull_number", c.PullNumber)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsListReviewsCmd struct {
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
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
	CommitMessage string `name:"commit_message"`
	CommitTitle   string `name:"commit_title"`
	MergeMethod   string `name:"merge_method"`
	PullNumber    int64  `required:"" name:"pull_number"`
	Repo          string `required:"" name:"repo"`
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

type PullsSubmitReviewCmd struct {
	Body       string `name:"body"`
	Event      string `required:"" name:"event"`
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
	ReviewId   int64  `required:"" name:"review_id"`
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
	ExpectedHeadSha string `name:"expected_head_sha"`
	Lydian          bool   `required:"" name:"lydian-preview"`
	PullNumber      int64  `required:"" name:"pull_number"`
	Repo            string `required:"" name:"repo"`
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
	Base                string `name:"base"`
	Body                string `name:"body"`
	MaintainerCanModify bool   `name:"maintainer_can_modify"`
	PullNumber          int64  `required:"" name:"pull_number"`
	Repo                string `required:"" name:"repo"`
	SailorV             bool   `name:"sailor-v-preview"`
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

type PullsUpdateCommentCmd struct {
	Body        string `required:"" name:"body"`
	ComfortFade bool   `name:"comfort-fade-preview"`
	CommentId   int64  `required:"" name:"comment_id"`
	Repo        string `required:"" name:"repo"`
	internal.BaseCmd
}

func (c *PullsUpdateCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/pulls/comments/{comment_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdatePreview("comfort-fade", c.ComfortFade)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("PATCH")
}

type PullsUpdateReviewCmd struct {
	Body       string `required:"" name:"body"`
	PullNumber int64  `required:"" name:"pull_number"`
	Repo       string `required:"" name:"repo"`
	ReviewId   int64  `required:"" name:"review_id"`
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
