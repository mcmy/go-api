package i18n

type LangStruct struct {
	AuthenticationFailed        string `toml:"authentication-failed"`
	TokenError                  string `toml:"token-error"`
	CachedUserParsingFailed     string `toml:"cached-user-parsing-failed"`
	NotLogin                    string `toml:"not-login"`
	InputParsingFailed          string `toml:"input-parsing-failed"`
	IncorrectUsernameOrPassword string `toml:"incorrect-username-or-password"`
	RedisError                  string `toml:"redis-error"`
	RequestSuccess              string `toml:"request-success"`
	RequestError                string `toml:"request-error"`
	VerificationCodeError       string `toml:"verification-code-error"`
}

var (
	VerificationCodeError       = "verification-code-error"
	RequestSuccess              = "request-success"
	RequestError                = "request-error"
	IncorrectUsernameOrPassword = "incorrect-username-or-password"
	RedisError                  = "redis-error"
	AuthenticationFailed        = "authentication-failed"
	TokenError                  = "token-error"
	CachedUserParsingFailed     = "cached-user-parsing-failed"
	NotLogin                    = "not-login"
	InputParsingFailed          = "input-parsing-failed"
)
