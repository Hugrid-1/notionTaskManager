package bot

var (
	adminUserID    int
	allowedUserIDs []int
)

func isAdminUser(userID int) bool {
	return userID == adminUserID
}

func isAllowedUser(userID int) bool {
	for _, allowedID := range allowedUserIDs {
		if userID == allowedID {
			return true
		}
	}
	return false
}
