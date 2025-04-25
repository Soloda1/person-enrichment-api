package main

import (
	"fmt"
	"log"
	"person-enrichment-api/config"
)

func main() {
	cfg := config.MustLoad()

	log.Printf("Environment: %s", cfg.Env)
	log.Printf("Server Address: %s", cfg.HTTPServer.Address)
	log.Printf("Database Connection: postgresql://%s:%s@%s:%s/%s",
		cfg.DATABASE.Username,
		cfg.DATABASE.Password,
		cfg.DATABASE.Host,
		cfg.DATABASE.Port,
		cfg.DATABASE.DbName)

	log.Printf("External APIs:")
	log.Printf("- Agify URL: %s", cfg.EXTERNAL_API.AgifyURL)
	log.Printf("- Genderize URL: %s", cfg.EXTERNAL_API.GenderizeURL)
	log.Printf("- Nationalize URL: %s", cfg.EXTERNAL_API.NationalizeURL)

	fmt.Println("Configuration loaded successfully!")
}
