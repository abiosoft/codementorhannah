package model

type User struct {
	ID       int    `json:"user_id"` // struct tags
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsers() ([]User, error) {
	rows, err := db.Query("SELECT * FROM user LIMIT 10")
	if err != nil {
		return nil, err
	}
	var users []User
	var user User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// 1000 records satisfying this
// db.user.find({"usernmae" : "bl"}) // you get 1000
// select * from users where ...

// you get just 10 records only
// db.user.find({"usernmae" : "bl"}).sort({username : -1}).limit(10)
// select * from users where .... order by username desc limit 10
