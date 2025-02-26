# CivAPI
![CI](https://github.com/Lunnaris01/Civ_API/actions/workflows/ci.yml/badge.svg)

A lightweight REST API + Frontend to get Insights into your Civilisation VI data!
Technologies:
- Programming Language: GO
- Database and Database Migration: sqlite (Turso) and goose
- Router: [Chi](https://github.com/go-chi/chi/tree/master)
- CI using Github Actions
- CD using docker and gcloud (currently not in use!)
- Simple Frontend with HTML, CSS and JS

(Upcoming) Features:
- Authentication and Authorization
- Save your game stats
- Calculate and Visualize Stats for yourself, your friendgroup or the whole database.
- "Wiki" like pages for Different Nations and Leaders showcasing their typical win conditions, winrates and others stats.

Project Goals:
Code a stable and secure backend using GO and raw SQL (sqlite) instead of sqlc
Create a simple and light weight frontend for User Interaction.
CI using Github Actions for Testing, Stylechecks/linting, ...
Remote Database using Turso

Optional Long Term Goals:
CD using docker+gcloud.
Improve Frontend to have great looking visualization for the statistics.

## Usage 

### Clone the repo

```bash
git clone https://github.com/Lunnaris01/CivAPI@latest
cd CivAPI
```

### Build the project

```bash
make build
```

### Start the Server

```bash
make run
```
## 🤝 Contributing

Add Code and Tests

### Run the tests

```bash
go test ./...
```

### Ensure your code adheres to style conventions and passes the staticcheck:


```bash
go fmt ./...
```

```bash
staticcheck ./...
```


### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch. 
