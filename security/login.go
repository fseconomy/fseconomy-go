package security

func (sec *Security) Login(username string, password string) bool {
	return username == password
}
