# iConText-test-task

---
### Task description
<a>https://observant-hero-c7c.notion.site/Go-3cc65a7d7c3e44c19b2e0543a98be2d2</a>

---
### Instruction for starting
0. <p>You need to create an .env file and put it in the root of the application. An example .env file can be seen below.</p>
1. <b>```make test```</b>
2. <b>```make db_init```</b>
3. <b>```make build```</b>
4. <b>```make run```</b>

---

### .env file example
```
# HTTP server
HTTP_ADDR=127.0.0.1
HTTP_PORT=9000

# Postgres
PG_ROLE=kdl
PG_PASSWORD=*your_password*
PG_HOST=127.0.0.1
PG_PORT=5432
PG_DBNAME=icontext_test_task
PG_SSLMODE=disable
PG_INIT_TABLES_FILE_PATH="schema/pg_init_tables.sql"
PG_KEEP_ALIVE_POOL_PERIOD=3

# Redis
RDS_HOST=127.0.0.1
RDS_PORT=6379
RDS_PASSWORD=*your_password*
RDS_DB=10
RDS_KEEP_ALIVE_POOL_PERIOD=3

# 
LOG_FILE=common.log
```

---