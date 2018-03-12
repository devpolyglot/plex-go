package plex

type User struct {
	ID           int    `json:"id"`
	UUID         string `json:"uuid"`
	Email        string `json:"email"`
	JoinedAt     string `json:"joined_at"`
	Username     string `json:"username"`
	Thumb        string `json:"thumb"`
	AuthToken    string `json:"authToken"`
	Subscription struct {
		Active   bool     `json:"active"`
		Status   string   `json:"status"`
		Plan     string   `json:"plan"`
		Features []string `json:"features"`
	} `json:"subscription"`
	Roles struct {
		Roles []string `json:"roles"`
	} `json:"roles"`
	Entitlements []string `json:"entitlements"`
	ConfirmedAt  string   `json:"confirmedAt"`
	ForumID      int      `json:"forumId"`
	RememberMe   bool     `json:"rememberMe"`
}

type SignInResponse struct {
	User User `json:"user"`
}

func (u User) String() string {
	return Stringify(u)
}
