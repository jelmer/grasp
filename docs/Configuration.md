# Configuring Grasp

All configuration in Grasp is optional. If you supply no configuration values then Grasp will default to using a SQLite database in the current working directory.

If you're already running MySQL or PostgreSQL on the server you're installing Grasp on, you'll most likely want to use one of those as your database driver.

To do so, set the appropriate environment variables when starting grasp.

`
source /home/john/grasp.env
grasp server
`

The default configuration looks like this:

```
GRASP_GZIP=true
GRASP_DEBUG=true
GRASP_DATABASE_DRIVER="sqlite3"
GRASP_DATABASE_NAME="./grasp.db"
GRASP_DATABASE_USER=""
GRASP_DATABASE_PASSWORD=""
GRASP_DATABASE_HOST=""
GRASP_DATABASE_SSLMODE=""
GRASP_SECRET="random-secret-string"
```

### Accepted values & defaults

| Name | Default | Description
| :---- | :---| :---
| GRASP_DEBUG | `false` | If `true` will write more log messages.
| GRASP_SERVER_ADDR | `:8080` | The server address to listen on
| GRASP_GZIP | `false` | if `true` will HTTP content gzipped
| GRASP_DATABASE_DRIVER | `sqlite3` | The database driver to use: `mysql`, `postgres` or `sqlite3`
| GRASP_DATABASE_NAME |  | The name of the database to connect to (or path to database file if using sqlite3)
| GRASP_DATABASE_USER |  | Database connection user
| GRASP_DATABASE_PASSWORD | | Database connection password
| GRASP_DATABASE_HOST |  | Database connection host
| GRASP_DATABASE_SSLMODE | | For a list of valid values, look [here for Postgres](https://www.postgresql.org/docs/9.1/static/libpq-ssl.html#LIBPQ-SSL-PROTECTION) and [here for MySQL](https://github.com/Go-SQL-Driver/MySQL/#tls)
| GRASP_DATABASE_URL | | Can be used to specify the connection string for your database, as an alternative to the previous 5 settings. 
| GRASP_SECRET |  | Random string, used for signing session cookies

### Common issues

##### Grasp panics when trying to connect to Postgres: `pq: SSL is not enabled on the server`

This usually means that you're running Postgres without SSL enabled. Set the `GRASP_DATABASE_SSLMODE` config option to remedy this.

```
GRASP_DATABASE_SSLMODE=disable
```

##### Using `GRASP_DATABASE_URL`

When using `GRASP_DATABASE_URL` to manually specify your database connection string, there are a few important things to consider.

- When using MySQL, include `?parseTime=true&loc=Local` in your DSN.
- When using SQLite, include `?_loc=auto` in your DSN.

Examples of valid values:

```
GRASP_DATABASE_DRIVER=mysql
GRASP_DATABASE_URL=root:@tcp/grasp1?loc=Local&parseTime=true
```
