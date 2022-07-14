# Video Hosting Web-Service

### Available requests:

* POST   /auth/signin
* POST   /auth/change_password
* POST   /auth/signout
* POST   /user
* GET    /user/:id
* GET    /user/all
* POST   /user/change_password
* PATCH  /user/:id
* DELETE /user/:id
* POST   /group
* GET    /group/:id
* GET    /group/all
* PATCH  /group/:id
* DELETE /group/:id
* POST   /group/user/:id
* GET    /group/user/:id
* DELETE /group/user/:id
* GET    /perm/all
* POST   /perm/user/:id
* GET    /perm/user/:id
* DELETE /perm/user/:id
* POST   /perm/group/:id
* GET    /perm/group/:id
* DELETE /perm/group/:id
* POST   /info
* GET    /info/:id
* GET    /info/all
* PATCH  /info/:id
* DELETE /info/:id
* POST   /video
* GET    /video/:id
* GET    /video/all
* PATCH  /video/:id
* DELETE /video/:id
* GET    /stream/get/:id
* GET    /stream/get/all

### To watch available streams:

Post in your web browser 127.0.0.1:8000/stream

### Deploying:

#### 1. Create an .env file in directory ./configs/ and post variables from example .env.example.

##### struct config.yml
```
db:
  connTimeoutSeconds: 5                             // how many seconds it takes to try to connect to the database
  connLatencyMilliseconds: 50                       // how many milliseconds will be the delay in sending each request to the database
  host: "host"                                      // what ip is the database server
  logConnStatus: false                              // whether to display the status of connecting / disconnecting from the database in the logs
  name: "dbName"                                    // database name
  port: "port"                                      // database port
  sslMode: "disable"                                // SSL encryption for sending/receiving queries to the database
  username: "username"                              // username database

pagination:
  getLimitDefault: 100                              // number of output elements in get methods by default

server:
  debugMode: true                                   // put server state into debug mode
  maxHeaderBytes: 1048576                           // reserve 1 megabyte to store request headers per server
  port: "8000"                                      // web server port
  readTimeoutSeconds: 15                            // maximum time in seconds to receive responses from the server
  writeTimeoutSeconds: 15                           // maximum time in seconds to send requests to the server

session:
  ttlHours: 168                                     // cookie lifetime in browser and session in hours (which is equal to 7 days)
  
```
##### struct .env

```
DB_DRIVER = "oracle_db"                             // name to adjust to your database
DB_PASSWORD = "password"                            // password database
HASHING_PASSWORD_SALT = "aBw#ar\123_=@#$%1"         // password hashing salt
HASHING_TOKEN_SIGNING_KEY = "cD@#$%2"               // signing key for generating and reading tokens
```


2. Create database named "vhosting" in your DBMS and create tables by executing
SQL query file vhosting_database.sql.
3. Build a binary with this command:

```go build ./cmd/app```

4. You have to install or update several AV-libraries.

On Xubuntu 20.04 or higher post it to install/update all libraries:
```
sudo apt-get install libavformat-dev
sudo apt-get install libavresample-dev
sudo apt-get install libavcodec-dev
```
On Debian 11.3 or higher - install/update only two of those:
```
apt install libavformat-dev
apt install libavresample-dev
```
5. Make sure you have nginx installed by executing command:

systemctl status nginx

If it does not installed you have to install it by this command:
```
apt install libpq-dev postgresql postgresql-contrib nginx curl
```
