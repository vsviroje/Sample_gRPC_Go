package helper

import (
	dbAccess "Sample_gRPC/DBAccess"
)

func PrepareUserRec() {
	userList := dbAccess.GetUserList()
	userMap := dbAccess.GetUserMap()

	for i := range userList {
		userMap[userList[i].ID] = userList[i]
	}

	dbAccess.SetUserMap(userMap)
}
