package models

type SecretStructJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Engine   string `json:"engine"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

// register data
type SignUp struct {
	Username string `json:"username"`
	UserUUID string `json:"userUUID"`
}