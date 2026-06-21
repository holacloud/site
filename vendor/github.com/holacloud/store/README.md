# Store

This is a spike project to implement an effective way to actually abstract the
persistence layer of the same entities in different and quite heterogeneous 
persistence backends (postgres, mongodb, memory, disk).

## How to run it

```sh
make test
```

Or run it in docker if MongoDB or Postgres is not available locally:

```sh
make docker-test
```
