package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		wantErr bool
	}{
		{
			name: "loads config successfully",
			env: map[string]string{
				"DB_URL":       "postgres://localhost:5432/mydb",
				"LIMEN_SECRET": "super-secret",
				"API_VERSION":  "v1",
				"API_PORT":     "8080",

				"GOOSE_MIGRATION_DIR": "./migrations",
				"GOOSE_DRIVER":        "postgres",
				"GOOSE_DBSTRING":      "postgres://localhost:5432/mydb",
			},
			wantErr: false,
		},
		{
			name: "returns error when required environment variable is missing",
			env: map[string]string{
				"DB_URL":      "postgres://localhost:5432/mydb",
				"API_VERSION": "v1",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				t.Setenv(k, v)
			}

			_, err := NewConfig("/dev/null")

			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
