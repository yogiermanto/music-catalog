package configs

type (
	Config struct {
		Service  Service
		Database Database
	}

	Service struct {
		SecretJWT string
	}

	Database struct {
		DSN string
	}
)
