package references

const (
	CreateUserQuery     = "INSERT INTO (name,surname,username,email,university,password,role,created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW()) RETURNING id"
	GetUserByEmailQuery = ""
	GetUserByIDQuery    = ""
)
