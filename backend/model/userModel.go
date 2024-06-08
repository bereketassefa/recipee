package model

type TagSchema struct {
	Insert_Tags struct {
		Returning []struct {
			TagName string
		}
	}
}

//	"insert_users": {
//		"affected_rows": 1
//	  }
type CreatedUser struct {
	Insert_users_one struct {
		id string
	}
}

type LoginSchema struct {
	Users []struct {
		Id       string
		Password string
	}
}
