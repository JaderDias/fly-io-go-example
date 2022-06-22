# go-example

A minimal Go application for [fly.io Getting Started](https://fly.io/docs/getting-started/golang/) documentation and tutorials.

To get started:

1. clone this repo
2. `flyctl launch`
3. view the deployed app with `flyctl open`

To re-deploy

`flyctl deploy`

To connect to PostgreSQL

```
flyctl postgres connect -a firstapp-postgres
postgres=# \c weathered_sky_1150
```

To proxy PostgreSQL

`flyctl proxy 5432 -a firstapp-postgres`

To run locally

`DATABASE_URL=postgres://weathered_sky_1150:{PASSWORD}@localhost:5432/weathered_sky_1150 ./go-example `