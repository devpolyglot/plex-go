package plex

type User struct {
	ID           int    `json:"id,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	Email        string `json:"email,omitempty"`
	JoinedAt     string `json:"joined_at,omitempty"`
	Username     string `json:"username,omitempty"`
	Thumb        string `json:"thumb,omitempty"`
	AuthToken    string `json:"authToken,omitempty"`
	Subscription struct {
		Active   bool     `json:"active,omitempty"`
		Status   string   `json:"status,omitempty"`
		Plan     string   `json:"plan,omitempty"`
		Features []string `json:"features,omitempty"`
	} `json:"subscription,omitempty"`
	Roles struct {
		Roles []string `json:"roles,omitempty"`
	} `json:"roles,omitempty"`
	Entitlements []string `json:"entitlements,omitempty"`
	ConfirmedAt  string   `json:"confirmedAt,omitempty"`
	ForumID      int      `json:"forumId,omitempty"`
	RememberMe   bool     `json:"rememberMe,omitempty"`
}

type SignInResponse struct {
	User User `json:"user,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}
