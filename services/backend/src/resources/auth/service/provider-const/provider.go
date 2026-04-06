package provider

type AuthProvider string

const (
	Authentik AuthProvider = "authentik"
	E2EMock   AuthProvider = "e2e-mock"
)

func (ap AuthProvider) String() string {
	return string(ap)
}
