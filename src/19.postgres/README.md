# Postgres

> **Execute a criação das extensões abaixo em seus postgres antes de rodar as migrations**

```sql
CREATE USER estudos;

ALTER USER estudos WITH ENCRYPTED password 'teste123';

GRANT ALL privileges on database golangestudos to estudos;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT * FROM pg_available_extensions;

SELECT * FROM pg_extension;
```

> **Defina o DSN correto no arquivo `.env`. Tome como exemplo o arquivo `.env.example`**

## Ferramental

Use [golang-migrate cli](https://github.com/golang-migrate/migrate#cli-usage)

```
$ make migrateup
migrate -path /home/lucas/workspaces/go/golang-estudos/src/19.postgres/migrations -database "postgresql://estudos:teste123@localhost:25432/golangestudos?sslmode=disable" -verbose up
2020/09/25 20:26:49 Start buffering 1/u init_schema
2020/09/25 20:26:49 Start buffering 2/u create_account
2020/09/25 20:26:49 Start buffering 3/u create_entries
2020/09/25 20:26:49 Start buffering 4/u create_transfers
2020/09/25 20:26:49 Start buffering 5/u create_indexes
2020/09/25 20:26:49 Read and execute 1/u init_schema
2020/09/25 20:26:49 Finished 1/u init_schema (read 6.932633ms, ran 6.148751ms)
2020/09/25 20:26:49 Read and execute 2/u create_account
2020/09/25 20:26:49 Finished 2/u create_account (read 19.29699ms, ran 19.987021ms)
2020/09/25 20:26:49 Read and execute 3/u create_entries
2020/09/25 20:26:49 Finished 3/u create_entries (read 44.644315ms, ran 11.37786ms)
2020/09/25 20:26:49 Read and execute 4/u create_transfers
2020/09/25 20:26:49 Finished 4/u create_transfers (read 61.820987ms, ran 12.765959ms)
2020/09/25 20:26:49 Read and execute 5/u create_indexes
2020/09/25 20:26:49 Finished 5/u create_indexes (read 83.116029ms, ran 36.592756ms)
2020/09/25 20:26:49 Finished after 125.25886ms
2020/09/25 20:26:49 Closing source and database

```

## Links

-   https://dev.to/techschoolguru/how-to-write-run-database-migration-in-golang-5h6g
