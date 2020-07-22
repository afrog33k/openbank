// Code generated by "scopegen"; DO NOT EDIT.
package profiles

type ServiceScope struct{}

var Scopes = map[string]string{
	"https://auth.bnk.to/profile.read": "View profile data",
	"https://auth.bnk.to/profile.write": "Manage profile data",
}

var AuthScopes = map[string][]string{
	"/profiles.ProfileService/GetProfile": []string{"https://auth.bnk.to/profile.read"},
	"/profiles.ProfileService/GetProfileCard": []string{"https://auth.bnk.to/profile.read"},
}

// Any allows a loose challenge, for claims containing any of the method scopes.
//
// Use All method as a default for OAuth2 style scopes.  Any is useful with more complex scope definitions.
func (svcSc *ServiceScope) Any(method string, claims []string) bool {
	ch := AuthScopes[method]
	for _, s := range ch {
		for _, c := range claims {
			if string(s) == c {
				return true
			}
		}
	}
	return len(ch) == 0
}

// All is the default OAuth2 challenge method, ensuring claims contains all method scopes.
func (svcSc *ServiceScope) All(method string, claims []string) bool {
	ch := AuthScopes[method]
	if len(ch) > len(claims) {
		return false
	}
scopes:
	for _, s := range ch {
		for _, c := range claims {
			if string(s) == c {
				continue scopes
			}
		}
		return false
	}
	return true
}
