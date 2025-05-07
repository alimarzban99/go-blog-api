package admin

type UserCollection struct{}
type UserResource struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt string `json:"created_at"`
}
