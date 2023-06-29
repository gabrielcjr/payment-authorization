# Cred-System Interview Challenge

## Using for this challenge:
- Postgres database
- gRPC
- REST

## Standards
<a href="https://github.com/golang-standards/project-layout">golang standards - project layout</a>
<a href="https://github.com/uber-go/guide">Uber style guide for golang</a>

## go have fun using Go -> 
<a href="https://go.dev/solutions/#case-studies">Why Go</a>

## Using the App
- docker-compose up --build at the root of the app
- in a PostgreSQL client, manually execute the script (copy, paste & run) ./migration/CREATE _TRANSACTION.sql
- in a http client, execute the routes:
  -- HEALTH
  curl --request GET \
    --url http://localhost:3150/health
  -- POST
      curl --request POST \
        --url http://localhost:3150/transactions \
        --header 'Content-Type: application/json' \
        --data '{
        "client_id": "{uuid}",
        "value": {float}
      }'
  -- GET
  curl --request GET \
    --url http://localhost:3150/transactions/{uuid} \
    --header 'Content-Type: application/json'
  -- DELETE
  curl --request DELETE \
    --url http://localhost:3150/transactions/{uuid} \
    --header 'Content-Type: application/json'
  -- UPDATE
  curl --request PUT \
  --url http://localhost:3150/transactions/{uuid} \
  --header 'Content-Type: application/json' \
  --data '{
    "value": {float}
  }'