package main


import(
	"fmt"
	"os"
	"github.com/joho/godotenv"

)



type apiConfig struct {
	//dbQueries *database.Queries
	dbURL string
	dbToken string
	platform string
}

func main(){
	fmt.Println("Civ API started!")
	godotenv.Load()

	env_platform := os.Getenv("PLATFORM")
	env_dbURL := os.Getenv("TURSO_DATABASE_URL")
	env_dbToken := os.Getenv("TURSO_AUTH_TOKEN")


	apiCfg := apiConfig {
		dbURL: env_dbURL,
		dbToken: env_dbToken,
		platform: env_platform,
	}

	fmt.Println(apiCfg)

}