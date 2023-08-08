package url

const (
	Home              string = "/"
	About             string = "/about"
	SignUp            string = "/signup"
	Login             string = "/login"
	VerifyEmail       string = "/verify-email"
	ResendVerifyEmail string = "/resend-verify-email"
	ForgotPassword    string = "/forgot-password"
	ResetPassword            = "/reset-password"
	Logout            string = "/logout"
	Profile           string = "/profile"
	Settings          string = "/settings"
	Static            string = "/public"

	//Blog
	CreatePost string = "/CreatePost"

	UserPanel string = "Dashboard"

	Test string = "/test"
)

func Geturlpath() map[string]string {
	return map[string]string{
		"Static":            Static,
		"Home":              Home,
		"About":             About,
		"Signup":            SignUp,
		"Login":             Login,
		"VerifyEmail":       VerifyEmail,
		"ResendVerifyEmail": ResendVerifyEmail,
		"ForgotPassword":    ForgotPassword,
		"ResetPassword":     ResetPassword,
		"Logout":            "/" + UserPanel + Logout,
		"Profile":           "/" + UserPanel + Profile,
		"Settings":          "/" + UserPanel + Settings,
		"CreatePost":        "/" + UserPanel + CreatePost,
	}
}
