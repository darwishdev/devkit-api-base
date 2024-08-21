package supaclient

import (
	"log"
	"os"
	"testing"

	"github.com/darwishdev/devkit-api-base/config"
	_ "github.com/lib/pq"
)

var service SupabaseServiceInterface

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../config", "dev")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	service, err = NewSupabaseService(config.SupaUrl, config.SupaKey)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	os.Exit(m.Run())
}
