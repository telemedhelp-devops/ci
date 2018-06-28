package models

//go:generate reform

//reform:user_profiles
type UserProfile struct {
	Id       int    `reform:"id,pk"`
	Username string `reform:"username" sql_size:"255"`
	SMSTo    string `reform:"sms_to"`
}
