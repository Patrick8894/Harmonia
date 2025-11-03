package auth

type SessionStore interface {
	Create(user string) (token string, err error)
	Get(token string) (user string, ok bool)
	Delete(token string)
}
