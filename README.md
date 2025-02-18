# CivAPI
![CI](https://github.com/Lunnaris01/Civ_API/actions/workflows/ci.yml/badge.svg)

A lightweight REST API + Frontend to get Insights into your Civ VI data!
Technologies:
- Programming Language: GO
- Database: sqlite (Turso)
- Router: [Chi](https://github.com/go-chi/chi/tree/master)
- CI using Github Actions
- CD using docker and gcloud

(Upcoming) Features:
- Usermanagement
- Friendgroups
- Save your game stats
- Calculate individual stats (winrate, rate per leader, average round to win, most common wincondition, ...)
- See accumulated stats for friendgroups
- ...

Project Goals:
Code a stable and secure backend using GO and raw SQL (sqlite) instead of sqlc
Create a simple and light weight frontend for User Interaction.
CI using Github Actions for Testing, Stylechecks/linting, ...
Remote Database using Turso

Optional Long Term Goals:
CD using docker+gcloud.
Improve Frontend to have great looking visualization for the statistics.

