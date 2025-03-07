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
- Calculate and Visualize Stats.
- Add friends for your User and to your Games, see their stats or combined stats for yourself and multiple others.
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


# CivAPI Documentation

CivAPI is a RESTful API for managing user accounts and game data. This documentation provides details on how to interact with the API.

---

## Table of Contents
1. [Authentication](#authentication)
2. [Routes](#routes)
   - [Public Routes](#public-routes)
     - [Login](#login)
     - [Signup](#signup)
   - [Protected Routes](#protected-routes)
     - [Get Games](#get-games)
     - [Add Game](#add-game)
     - [Dashboard Content](#dashboard-content)
3. [Error Handling](#error-handling)

---

## Authentication

To access protected routes, you must include a valid JWT (JSON Web Token) in the `Authorization` header of your request. The token can be obtained by logging in via the `/login` endpoint.

**Header Format**:
```
Authorization: Bearer <JWT_TOKEN>
```

---

## Routes

### Public Routes

#### **Login**
- **Endpoint**: `POST /login`
- **Description**: Authenticate a user and retrieve a JWT for accessing protected routes.
- **Request Body**:
  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```
- **Response**:
  ```json
  {
    "User": {
      "ID": "user_id",
      "Username": "your_username"
    },
    "Token": "your_jwt_token"
  }
  ```

#### **Signup**
- **Endpoint**: `POST /signup`
- **Description**: Create a new user account.
- **Request Body**:
  ```json
  {
    "username": "your_username",
    "password": "your_password",
    "displayName": "your_display_name"
  }
  ```
- **Response**:
  ```json
  {
    "Username": "your_username"
  }
  ```

---

### Protected Routes

#### **Get Games**
- **Endpoint**: `GET /api/games`
- **Description**: Retrieve a list of games associated with the authenticated user. The GameCode can be used to share the game with friends who can add themselves to the game.
- **Headers**:
  ```
  Authorization: Bearer <JWT_TOKEN>
  ```
- **Response**:
  ```json
  [
    {
      "Country": "game_country",
      "GameWon": true,
      "WinCondition": "victory_type",
      "GameCode": "123454345321"
    }
  ]
  ```

#### **Add Game**
- **Endpoint**: `POST /api/games`
- **Description**: Add a new game for the authenticated user. There is 2 options to do this, First is by having a country, game_won and win_condition present without sending a GameCode or sending a Gamecode.
- **Headers**:
  ```
  Authorization: Bearer <JWT_TOKEN>
  ```
- **Request Body**:
- Add a new game via form
  ```json
  {
    "country": "game_country",
    "game_won": true,
    "win_condition": "victory_type",
  }
  ```
  
- Add the user to an existing game from another player  
  ```json
  {
    "GameCode": "123454345321"
  }
  ```

- **Response**:
  ```
  Added Successfully
  ```

#### **Dashboard Content**
- **Endpoint**: `GET /content`
- **Description**: Retrieve static content for the user dashboard.
- **Headers**:
  ```
  Authorization: Bearer <JWT_TOKEN>
  ```
- **Response**: Static HTML/CSS/JS content.

---

## Error Handling

If an error occurs, the API will respond with a JSON object containing an error message and status code.

**Error Response Format**:
```json
{
  "error": "Error message",
  "details": "Additional error details (if available)"
}
```

**Common Error Codes**:
- `400 Bad Request`: Invalid request data.
- `401 Unauthorized`: Missing or invalid authentication token.
- `500 Internal Server Error`: Server-side error.

---

## Example Usage

### Login
```bash
curl -X POST http://your-api-url/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "testpassword"}'
```

### Get Games
```bash
curl -X GET http://your-api-url/api/games \
  -H "Authorization: Bearer <JWT_TOKEN>"
```

### Add Game
```bash
curl -X POST http://your-api-url/api/games \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"country": "France", "game_won": true, "win_condition": "Domination"}'
```

---
