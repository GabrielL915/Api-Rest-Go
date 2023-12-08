### Require
* Docker and Go
* [golang-migrate/migrate](https://github.com/golang-migrate/migrate) 

### Uso
Clone the repository with:
```bash
git clone https://github.com/GabrielL915/Api-Rest-Go.git
```
criar env no bash
```bash
$ cp .env.example .env
```

Build e start do projeto:
```bash
$ docker-compose up --build
```
migrar banco de dados no bash(trocar $PG_USER, $PG_PASS, $PG_DB por valor da .env):
```bash
$ export POSTGRESQL_URL="postgres://$PG_USER:$PG_PASS@localhost:5432/$PG_DB?sslmode=disable"
$ migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
