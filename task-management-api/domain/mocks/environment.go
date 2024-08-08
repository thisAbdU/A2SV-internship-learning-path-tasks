package mocks

type Environment struct {
	JwtKey string
	DbURL  string
	DbName string
	Port   string
}

func NewMockEnvironment() *Environment {
	return &Environment{
		JwtKey: "mock-jwt-key",
		DbURL:  "mongodb://localhost:27017",
		DbName: "testdb",
		Port:   "8080",
	}
}
