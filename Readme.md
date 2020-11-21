# Adding a db model
- `cd` into the backend directory and run `entc init <model>`
- run `go run main.go -generate=true`


# Adding new fields and/or constraints to a db model

- Modify the model in backend/ent/schemas and cd into backend and run `go run main.go -generate=true` in the backend directory.

# Adding functionality or modifying the backend rpc service

- Modify the service spec in `service/service.proto`.
- Re generate the server and client code by going into the service directory and running `./gen.sh`


# Setup dev DB
- Run `./init_dev_db.sh`

# Building and starting dev server
- Run `./start_dev.sh`
