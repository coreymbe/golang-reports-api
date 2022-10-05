# Golang Reports API

This is an example custom API for processing Puppet report data.

## Setup

```
git clone https://github.com/coreymbe/golang-reports-api
```

```
cd golang-reports-api
```

```
go build
```

## Usage

**Note**: The server utilizes PostgreSQL as the backend. Before starting the server you will need to create the database and table used by the server. This can be accomplished by running the `sql` file at `database/puppet_reports.sql`.

Additionally, you will need to set the `dbName`, `dbUser`, and `dbPass` environment variables for the server to successfully connect to the database.

```
export dbName=puppet
export dbUser=postgres
export dbPass=<Database user password>
```

To successfully generate and verify auth tokens for the API, you will need to set the `JWTSecretKey` environment variable with a random secure character string.

```
export JWTSecretKey=<Random character string>
```

---

### Start the server:

```
./golang-reports-api
```

---

### Generate an Auth Token

```
/bin/curl -X POST \
-H "Content-Type: application/json" \
-d '{"username":"admin","password":"ch@ngem3"}' \
http://localhost:2754/auth
```

**Note**: By default the token is valid for **12** hours.

---

### Create

```
export JWToken=<Generated auth token>
```

```
/bin/curl -X POST \
-H "Authorization:Bearer ${JWToken}" \
-H "Content-Type: application/json" \
-d '{"certname":"example-hostname.puppet.com","environment":"production","status":"unchanged","time":"2022-09-25 19:30:00 UTC","transaction_uuid":"c8f1d280-21c6-4fd3-a28a-c150da1b0ecf"}' \
http://localhost:2754/reports/add
```

### READ

All Reports:

```
/bin/curl http://localhost:2754/reports
```

Report by ID:

```
/bin/curl http://localhost:2754/reports/1
```

### DELETE

```
export JWToken=<Generated auth token>
```

```
/bin/curl -X POST \
-H "Authorization:Bearer ${JWToken}" \
http://localhost:2754/reports/remove/1
```
