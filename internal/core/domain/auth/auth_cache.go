package auth

type AuthCache interface {
	Set(k string, v string)
	Get(k string) string
	Del(k string)
}
