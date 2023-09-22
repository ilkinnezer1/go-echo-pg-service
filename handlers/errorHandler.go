package handlers

// ErrorHandler handles with simple errors
func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

// Error Messages based on HTTP and Credentials
func InvalidUserCredentials() map[string]string {
	invalidUserName := map[string]string{
		"error": "Invalid username or password",
	}
	return invalidUserName
}
func FailHashPasswd() map[string]string {
	failHashPassword := map[string]string{
		"error": "Failed to hash password",
	}
	return failHashPassword
}
func FailStoreData() map[string]string {
	failStoreData := map[string]string{
		"error": "Failed to store user in database",
	}
	return failStoreData
}
func FailUpdateLastLoginTime() map[string]string {
	FailUpdateLastLoginTime := map[string]string{
		"error": " Error while updating last login time",
	}
	return FailUpdateLastLoginTime
}
