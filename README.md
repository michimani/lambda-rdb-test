lambda-rdb-test
===

In this repository, you can try to see how the number of DB connections changes depending on the implementation method when connecting to RDB from Lambda functions implemented in Go.

# Usage

## Setup containers

```bash
docker-compose up -d
```

## Invoke Lambda functions

### Function does not explicitly close the DB connection

```bash
curl "http://localhost:8001/2015-03-31/functions/function/invocations"
```

This Lambda function will increase the number of DB connections each time it is executed. To release the connections, restart the app with the following command.

```bash
docker-compose restart lambda1
```

This operation will reproduce the state where the AWS Lambda execution environment is replaced.

### Function explicitly closes the DB connection

```bash
curl "http://localhost:8002/2015-03-31/functions/function/invocations"
```

This Lambda function closes the connection after each execution, so no matter how many times it is executed, the number of DB connections will not increase.

### Function eeuses the DB connection

```bash
curl "http://localhost:8003/2015-03-31/functions/function/invocations"
```

This Lambda function creates a DB connection only the first time it is started, and reuses the connection created the first time from the second time on. Therefore, the number of DB connections will increase by one.

## Check the number of DB connections

```bash
docker exec it lambda-rds bash
```

```bash
mysql -proot -e'show status like "Threads_connected";'
```