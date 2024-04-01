package models

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Website   string `json:"website"`
	Password  string `json:"password"`
}

type UserCreate struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Website   string `json:"website"`
	Password  string `json:"password"`
}



// message User {
// 	string id = 1;
// 	string name = 2;
// 	string last_name = 3;
// 	string username = 4;
// 	string email = 5;
// 	string bio = 6;
// 	string website = 7;
// 	string password = 8;
// 	string refresh_token = 9;
// 	string created_at = 10;
// 	string updated_at = 11;
// 	repeated Post posts = 12;
//   }

type UserBYtokens struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Website   string `json:"website"`
	Password  string `json:"password"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type UserByAccess struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserName    string `json:"user_name"`
	AccessToken string `json:"access_token"`
}

type RegisterUser struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}
