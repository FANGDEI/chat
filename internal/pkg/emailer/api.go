package emailer

func SendEmail(info Information) {
	defaultEmailerManager.SendEmail(info)
}

func IsEmail(account string) bool {
	return defaultEmailerManager.IsEmail(account)
}
