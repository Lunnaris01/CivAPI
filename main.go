package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	//"github.com/Lunnaris01/CivAPI/internal/auth"
	"github.com/Lunnaris01/CivAPI/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

//go:embed static/*
var staticFiles embed.FS

var contentTypes = map[string]string{
	".html": "text/html; charset=utf-8",
	".css":  "text/css; charset=utf-8",
	".js":   "text/javascript; charset=utf-8",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".gif":  "image/gif",
	".svg":  "image/svg+xml",
	".json": "application/json",
}

type apiConfig struct {
	db       *database.Queries
	platform string
	port     string
	secretKey string
}

func main() {
	fmt.Println("Civ API started!")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load necessary environment variables with err: %v", err)
	}

	env_platform := os.Getenv("PLATFORM")
	env_dbURL := os.Getenv("TURSO_DATABASE_BASE")
	env_dbToken := os.Getenv("TURSO_AUTH_TOKEN")
	env_port := os.Getenv("PORT")
	evn_secret := os.Getenv("SECRET_KEY")
	dbCombinedURL := env_dbURL + "?authToken=" + env_dbToken
	log.Printf("Connecting to db at %s,", env_dbURL)

	sqlitedb, err := sql.Open("libsql", dbCombinedURL)
	if err != nil {
		log.Fatalf("Failed to connect to database with err: %v\n", err)
	}
	defer sqlitedb.Close()

	dbQueries := database.InitDB(dbCombinedURL)

	log.Println("Database connection successful!")

	apiCfg := apiConfig{
		db:       	dbQueries,
		platform: 	env_platform,
		port:     	env_port,
		secretKey:	evn_secret,
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/*", apiCfg.handlerStatic)
	router.Post("/login", apiCfg.handlerLogin)
	router.Post("/signup", apiCfg.handlerSignup)
	router.Get("/content",apiCfg.handlerDashboard)
	router.Get("/api/games",apiCfg.handlerGetGames)

	log.Printf("Server running and waiting for requests on port %v\n", apiCfg.port)
	http.ListenAndServe(":"+apiCfg.port, router)

	fmt.Println(apiCfg)

}

func (cfg apiConfig) handlerStatic(w http.ResponseWriter, r *http.Request) {
	filepath := r.URL.Path
	cfg.displayFileserverContent(w, filepath)
}

func (cfg apiConfig) displayFileserverContent(w http.ResponseWriter, filepath string) {
	log.Printf("Requested path: %s", filepath)
	var ext string
	if filepath == "/" {
		filepath = "/static/html/login.html"
	} else if !strings.HasPrefix(filepath, "/static/") {
		ext = strings.ToLower(path.Ext(filepath))
		if ext == ".css" {
			filepath = "/static/css" + filepath
		} else if ext == ".html" {
			filepath = "/static/html" + filepath
		} else if ext == ".ico" {
			filepath = "/static/images" + filepath
		} else {
			filepath = "/static/html" + filepath + ".html"
		}
	}
	log.Printf("Filepath to open: %s", strings.TrimPrefix(filepath, "/"))
	f, err := staticFiles.Open(strings.TrimPrefix(filepath, "/"))
	if err != nil {
		log.Printf("Error opening index.html: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	w.Header().Set("Content-Type", contentTypes[ext])

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	if _, err := io.Copy(w, f); err != nil {
		log.Printf("Error copying file to response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
