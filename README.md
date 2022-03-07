# bareska-interview-project

News and Topic management using REST API with domain driven design architecture with some add-on features :
- Caching mechanism using redis
- Filter news by status
- Filter news by topic
- Unit tested API (WIP)

# List of endpoints (apis)

All enppoints including their respective descriptions are listed in /

# How to build

First you need to supply a correct `.env` file matching the sample structure from `.env.example`

You can build this repo manually using `go run .` and provide the postgresql and redis by yourself (also need to edit several configurations in the code for each service endpoints)

Or to make things easier, just use docker compose and run `docker-compose up --build -d` in the root directory of this repository, and you should be able to access this backend service in localhost:backend_port

## Sorry, im lazy

Okay, sure, we all are, so here you go (dont use this for production, in that case dont be lazy instead):
```bash
cat > .env <<EOF
POSTGRES_PASSWORD=hello_there
REDIS_PASSWORD=please_hire_me
MIGRATION_PASSWORD=lol
BACKEND_PORT=1337
DEBUG_MODE=true
LOG_FILE=logs
EOF

docker-compose up --build -d
```

## It didnt work

Make sure you deleted the docker volume of `bareksa_project_db` and `bareksa_project_cache` if you happen to change each of the service password after the first image creation

But, im sure that is not the error: in that case go check `logs.txt` file in the root of the repository folder to find out more
