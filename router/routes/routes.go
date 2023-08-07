package routes

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

	UserPanel string = "Dashboard"

	Test string = "/test"
)

func Geturlpath() map[string]string {
	return map[string]string{
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
	}
}
