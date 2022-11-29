package test

import (
	"os"
	"testing"

	"github.com/ha-zu/learn-clean-dev/src/configs"
)

func TestNew(t *testing.T) {
	// Env setting struct testing
	wantURL := "http:localhost:1323"
	t.Setenv("DEV_URL", wantURL)

	wantPort := "1323"
	t.Setenv("DEV_PORT", wantPort)

	cnf, err := configs.New()
	if err != nil {
		t.Fatalf("Can not configs Load: %v", err)
	}

	if cnf.DevUrl != wantURL {
		t.Fatal("Do not match development domain")
	}

	if cnf.DevPort != wantPort {
		t.Fatal("Do not match development port")
	}

}

func TestEnvSetting(t *testing.T) {
	// Setenv check testing
	compere := map[string]string{
		"DEV_URL":  "http:localhost:1323",
		"DEV_PORT": "1323",
	}

	t.Run("SetEnv list testing", func(t *testing.T) {
		configs.EnvSetting()
		for k, v := range compere {
			if val, ok := os.LookupEnv(k); !ok || val != v {
				t.Fatalf("Can not setting environment variable. key:%s, value:%s", k, v)
			}
		}
	})
}

func TestEnvLists(t *testing.T) {
	// Get EnvLists testing
	compere := map[string]string{
		"DEV_URL":  "http:localhost:1323",
		"DEV_PORT": "1323",
	}

	t.Run("Get env lists testing", func(t *testing.T) {
		envlists := configs.EnvLists()
		for k, v := range envlists {
			if val, ok := compere[k]; !ok || val != v {
				t.Fatalf("Different get env key:%s, value:%s", k, v)
			}
		}
	})
}
