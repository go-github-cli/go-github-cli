// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import internal "github.com/octo-cli/octo-cli/internal"

type ActivityCmd struct {
	CheckRepoIsStarredByAuthenticatedUser     ActivityCheckRepoIsStarredByAuthenticatedUserCmd     `cmd:""`
	DeleteRepoSubscription                    ActivityDeleteRepoSubscriptionCmd                    `cmd:""`
	DeleteThreadSubscription                  ActivityDeleteThreadSubscriptionCmd                  `cmd:""`
	GetFeeds                                  ActivityGetFeedsCmd                                  `cmd:""`
	GetRepoSubscription                       ActivityGetRepoSubscriptionCmd                       `cmd:""`
	GetThread                                 ActivityGetThreadCmd                                 `cmd:""`
	GetThreadSubscriptionForAuthenticatedUser ActivityGetThreadSubscriptionForAuthenticatedUserCmd `cmd:""`
	ListEventsForAuthenticatedUser            ActivityListEventsForAuthenticatedUserCmd            `cmd:""`
	ListNotificationsForAuthenticatedUser     ActivityListNotificationsForAuthenticatedUserCmd     `cmd:""`
	ListOrgEventsForAuthenticatedUser         ActivityListOrgEventsForAuthenticatedUserCmd         `cmd:""`
	ListPublicEvents                          ActivityListPublicEventsCmd                          `cmd:""`
	ListPublicEventsForRepoNetwork            ActivityListPublicEventsForRepoNetworkCmd            `cmd:""`
	ListPublicEventsForUser                   ActivityListPublicEventsForUserCmd                   `cmd:""`
	ListPublicOrgEvents                       ActivityListPublicOrgEventsCmd                       `cmd:""`
	ListReceivedEventsForUser                 ActivityListReceivedEventsForUserCmd                 `cmd:""`
	ListReceivedPublicEventsForUser           ActivityListReceivedPublicEventsForUserCmd           `cmd:""`
	ListRepoEvents                            ActivityListRepoEventsCmd                            `cmd:""`
	ListRepoNotificationsForAuthenticatedUser ActivityListRepoNotificationsForAuthenticatedUserCmd `cmd:""`
	ListReposStarredByAuthenticatedUser       ActivityListReposStarredByAuthenticatedUserCmd       `cmd:""`
	ListReposStarredByUser                    ActivityListReposStarredByUserCmd                    `cmd:""`
	ListReposWatchedByUser                    ActivityListReposWatchedByUserCmd                    `cmd:""`
	ListStargazersForRepo                     ActivityListStargazersForRepoCmd                     `cmd:""`
	ListWatchedReposForAuthenticatedUser      ActivityListWatchedReposForAuthenticatedUserCmd      `cmd:""`
	ListWatchersForRepo                       ActivityListWatchersForRepoCmd                       `cmd:""`
	MarkNotificationsAsRead                   ActivityMarkNotificationsAsReadCmd                   `cmd:""`
	MarkRepoNotificationsAsRead               ActivityMarkRepoNotificationsAsReadCmd               `cmd:""`
	MarkThreadAsRead                          ActivityMarkThreadAsReadCmd                          `cmd:""`
	SetRepoSubscription                       ActivitySetRepoSubscriptionCmd                       `cmd:""`
	SetThreadSubscription                     ActivitySetThreadSubscriptionCmd                     `cmd:""`
	StarRepoForAuthenticatedUser              ActivityStarRepoForAuthenticatedUserCmd              `cmd:""`
	UnstarRepoForAuthenticatedUser            ActivityUnstarRepoForAuthenticatedUserCmd            `cmd:""`
}

type ActivityCheckRepoIsStarredByAuthenticatedUserCmd struct {
	Repo string `name:"repo" required:"true"`
	internal.BaseCmd
}

func (c *ActivityCheckRepoIsStarredByAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/starred/{repo}")
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("GET")
}

type ActivityDeleteRepoSubscriptionCmd struct {
	Repo string `name:"repo" required:"true"`
	internal.BaseCmd
}

func (c *ActivityDeleteRepoSubscriptionCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/subscription")
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("DELETE")
}

type ActivityDeleteThreadSubscriptionCmd struct {
	ThreadId int64 `name:"thread_id" required:"true"`
	internal.BaseCmd
}

func (c *ActivityDeleteThreadSubscriptionCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/notifications/threads/{thread_id}/subscription")
	c.UpdateURLPath("thread_id", c.ThreadId)
	return c.DoRequest("DELETE")
}

type ActivityGetFeedsCmd struct {
	internal.BaseCmd
}

func (c *ActivityGetFeedsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/feeds")
	return c.DoRequest("GET")
}

type ActivityGetRepoSubscriptionCmd struct {
	Repo string `name:"repo" required:"true"`
	internal.BaseCmd
}

func (c *ActivityGetRepoSubscriptionCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/subscription")
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("GET")
}

type ActivityGetThreadCmd struct {
	ThreadId int64 `name:"thread_id" required:"true"`
	internal.BaseCmd
}

func (c *ActivityGetThreadCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/notifications/threads/{thread_id}")
	c.UpdateURLPath("thread_id", c.ThreadId)
	return c.DoRequest("GET")
}

type ActivityGetThreadSubscriptionForAuthenticatedUserCmd struct {
	ThreadId int64 `name:"thread_id" required:"true"`
	internal.BaseCmd
}

func (c *ActivityGetThreadSubscriptionForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/notifications/threads/{thread_id}/subscription")
	c.UpdateURLPath("thread_id", c.ThreadId)
	return c.DoRequest("GET")
}

type ActivityListEventsForAuthenticatedUserCmd struct {
	Username string `name:"username" required:"true"`
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListEventsForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/events")
	c.UpdateURLPath("username", c.Username)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListNotificationsForAuthenticatedUserCmd struct {
	All           bool   `name:"all"`
	Before        string `name:"before"`
	Page          int64  `name:"page"`
	Participating bool   `name:"participating"`
	PerPage       int64  `name:"per_page"`
	Since         string `name:"since"`
	internal.BaseCmd
}

func (c *ActivityListNotificationsForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/notifications")
	c.UpdateURLQuery("all", c.All)
	c.UpdateURLQuery("participating", c.Participating)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("before", c.Before)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListOrgEventsForAuthenticatedUserCmd struct {
	Username string `name:"username" required:"true"`
	Org      string `name:"org" required:"true"`
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListOrgEventsForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/events/orgs/{org}")
	c.UpdateURLPath("username", c.Username)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListPublicEventsCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListPublicEventsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/events")
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListPublicEventsForRepoNetworkCmd struct {
	Repo    string `name:"repo" required:"true"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListPublicEventsForRepoNetworkCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/networks/{repo}/events")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListPublicEventsForUserCmd struct {
	Username string `name:"username" required:"true"`
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListPublicEventsForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/events/public")
	c.UpdateURLPath("username", c.Username)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListPublicOrgEventsCmd struct {
	Org     string `name:"org" required:"true"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListPublicOrgEventsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/events")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListReceivedEventsForUserCmd struct {
	Username string `name:"username" required:"true"`
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListReceivedEventsForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/received_events")
	c.UpdateURLPath("username", c.Username)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListReceivedPublicEventsForUserCmd struct {
	Username string `name:"username" required:"true"`
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListReceivedPublicEventsForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/received_events/public")
	c.UpdateURLPath("username", c.Username)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListRepoEventsCmd struct {
	Repo    string `name:"repo" required:"true"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListRepoEventsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/events")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListRepoNotificationsForAuthenticatedUserCmd struct {
	Repo          string `name:"repo" required:"true"`
	All           bool   `name:"all"`
	Before        string `name:"before"`
	Page          int64  `name:"page"`
	Participating bool   `name:"participating"`
	PerPage       int64  `name:"per_page"`
	Since         string `name:"since"`
	internal.BaseCmd
}

func (c *ActivityListRepoNotificationsForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/notifications")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("all", c.All)
	c.UpdateURLQuery("participating", c.Participating)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("before", c.Before)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListReposStarredByAuthenticatedUserCmd struct {
	Direction string `name:"direction"`
	Page      int64  `name:"page"`
	PerPage   int64  `name:"per_page"`
	Sort      string `name:"sort"`
	internal.BaseCmd
}

func (c *ActivityListReposStarredByAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/starred")
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListReposStarredByUserCmd struct {
	Username  string `name:"username" required:"true"`
	Direction string `name:"direction"`
	Page      int64  `name:"page"`
	PerPage   int64  `name:"per_page"`
	Sort      string `name:"sort"`
	internal.BaseCmd
}

func (c *ActivityListReposStarredByUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/starred")
	c.UpdateURLPath("username", c.Username)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListReposWatchedByUserCmd struct {
	Username string `name:"username" required:"true"`
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListReposWatchedByUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/subscriptions")
	c.UpdateURLPath("username", c.Username)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListStargazersForRepoCmd struct {
	Repo    string `name:"repo" required:"true"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListStargazersForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/stargazers")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListWatchedReposForAuthenticatedUserCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListWatchedReposForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/subscriptions")
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityListWatchersForRepoCmd struct {
	Repo    string `name:"repo" required:"true"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *ActivityListWatchersForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/subscribers")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ActivityMarkNotificationsAsReadCmd struct {
	LastReadAt string `name:"last_read_at"`
	Read       bool   `name:"read"`
	internal.BaseCmd
}

func (c *ActivityMarkNotificationsAsReadCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/notifications")
	c.UpdateBody("last_read_at", c.LastReadAt)
	c.UpdateBody("read", c.Read)
	return c.DoRequest("PUT")
}

type ActivityMarkRepoNotificationsAsReadCmd struct {
	Repo       string `name:"repo" required:"true"`
	LastReadAt string `name:"last_read_at"`
	internal.BaseCmd
}

func (c *ActivityMarkRepoNotificationsAsReadCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/notifications")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("last_read_at", c.LastReadAt)
	return c.DoRequest("PUT")
}

type ActivityMarkThreadAsReadCmd struct {
	ThreadId int64 `name:"thread_id" required:"true"`
	internal.BaseCmd
}

func (c *ActivityMarkThreadAsReadCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/notifications/threads/{thread_id}")
	c.UpdateURLPath("thread_id", c.ThreadId)
	return c.DoRequest("PATCH")
}

type ActivitySetRepoSubscriptionCmd struct {
	Repo       string `name:"repo" required:"true"`
	Ignored    bool   `name:"ignored"`
	Subscribed bool   `name:"subscribed"`
	internal.BaseCmd
}

func (c *ActivitySetRepoSubscriptionCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/subscription")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("ignored", c.Ignored)
	c.UpdateBody("subscribed", c.Subscribed)
	return c.DoRequest("PUT")
}

type ActivitySetThreadSubscriptionCmd struct {
	ThreadId int64 `name:"thread_id" required:"true"`
	Ignored  bool  `name:"ignored"`
	internal.BaseCmd
}

func (c *ActivitySetThreadSubscriptionCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/notifications/threads/{thread_id}/subscription")
	c.UpdateURLPath("thread_id", c.ThreadId)
	c.UpdateBody("ignored", c.Ignored)
	return c.DoRequest("PUT")
}

type ActivityStarRepoForAuthenticatedUserCmd struct {
	Repo string `name:"repo" required:"true"`
	internal.BaseCmd
}

func (c *ActivityStarRepoForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/starred/{repo}")
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("PUT")
}

type ActivityUnstarRepoForAuthenticatedUserCmd struct {
	Repo string `name:"repo" required:"true"`
	internal.BaseCmd
}

func (c *ActivityUnstarRepoForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/starred/{repo}")
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("DELETE")
}
