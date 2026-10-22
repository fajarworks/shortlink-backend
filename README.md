# SHORTLINK API
![Go](https://img.shields.io/badge/go-1.27.1-black?style=for-the-badge&logo=go&logoColor=white&labelColor=blue)
![Gin](https://img.shields.io/badge/Gin--gonic-v1.12.0-black?style=for-the-badge&logo=gin&logoColor=black&labelColor=lightblue)
![Postgresql](https://img.shields.io/badge/Postgresql-17.11-black?style=for-the-badge&logo=postgresql&logoColor=black&labelColor=lightyellow)
![redis](https://img.shields.io/badge/Redis-8.10.02-black?style=for-the-badge&logo=redis&logoColor=white&labelColor=FF4438)
 
 
A URL shortener service that lets authenticated users create short links with custom slugs, manage their own links, and redirect visitors to the original URL.
 
 
## Technologies
 
- **Go** 1.27.1 — A statically typed programming language used to build the API and manage its backend services.
- **Gin** 1.12.0 — A lightweight HTTP web framework used for routing and handling API requests.
- **pgx** v5 — A PostgreSQL driver and toolkit used for database access.
- **PostgreSQL** 17.11 — A relational database used to store users and links.
- **Redis** 8.10.2 — An in-memory data store used for session/token management and fast data operations.
- **JWT** (golang-jwt v5) — Token-based authentication for protected endpoints.
- **Swagger** — Generating interactive API documentation and testing endpoints from the API's OpenAPI specification.
## Features
 
### **Authentication** :
- Registration
- Login and Logout using Json Web Token (JWT)
- Password hashing with argon2id
### **Links**:
- Create a short link with a custom slug (alphanumeric characters and hyphens only)
- View all links owned by the authenticated user
- Delete your own links
- Public redirect from a short slug to the original URL
## API Endpoints
 
### Auth
 
| Method | Endpoint         | Auth Required | Description                  |
|--------|------------------|:-------------:|------------------------------|
| POST   | `/auth/register` |       No      | Register a new user          |
| POST   | `/auth/login`    |       No      | Login and receive a JWT      |
| POST   | `/auth/logout`   |      Yes      | Logout and invalidate token  |
 
### Links
 
| Method | Endpoint      | Auth Required | Description                                  |
|--------|---------------|:-------------:|----------------------------------------------|
| POST   | `/links`      |      Yes      | Create a new short link                      |
| GET    | `/links`      |      Yes      | Get all links owned by the current user      |
| DELETE | `/links/:id`  |      Yes      | Delete a link by its ID                      |
| GET    | `/:slug`      |       No      | Redirect to the original URL of the slug     |
 
> Protected endpoints require the header `Authorization: Bearer <token>`.
 
 
# Usage Instruction
 
 
### Database Setup
 
1. Pull Docker Image Postgresql From [dockerhub](https://hub.docker.com/_/postgres)
```sh
docker pull postgres:17.11
```
 
2. Create container
```sh
docker run --name {container_name} -d  -p {hostPort:5432} -v {volume:/var/lib/postgresql/data} --env-file ./. postgres:17.11
```
 
 
### Cache Setup
 
1. Pull Redis Image from [dockerhub](https://hub.docker.com/_/redis)
```zsh
docker pull redis:8.10.2
```
2. setup your redis configuration
```
bind 0.0.0.0
port 6379
 
daemonize no
logfile ""
loglevel notice
 
save 3600 1
save 300 100
save 60 10000
 
appendonly yes
appendfsync everysec
 
maxmemory 256mb
maxmemory-policy allkeys-lru
 
protected-mode yes
 
user default off
user {superuser} on >{password} ~* &* +@all
user {youruser} on >{yourpass} resetchannels ~{prefixkeys}:* +@string +ping +del +@read
```
3. create docker cache container
```sh
docker run --name {container_name} -p {hostPort:6379} -v {redis.conf}:/usr/local/etc/redis/redis.conf -v {volume}:/data redis:8.10.2 redis-server /usr/local/etc/redis/redis.conf
```
 
4. SET UP server environment
```sh
#  CONFIG DB
DB_USER={your_db_user}
DB_PASS={your_db_pass}
DB_HOST={your_db_host}
DB_PORT={your_db_port}
DB_NAME={your_db_name}
 
# PORT BE
PORT={your_server_port}
SERVER_HOST={your_server_host}
 
 
# JWT
 
JWT_SECRET={rand-jwt-secret-key}
JWT_ISSUER={jwt-issuer}
 
#redis
REDIS_NAME={your_redis_username}
REDIS_PASSWORD={your_redis_password}
REDIS_PORT={your_redis_port}
REDIS_HOST={your_redis_host}
```
### Clone Repository
 
```sh
git clone https://github.com/fajarworks/shortlink-backend.git
```
**Download Dependencies**
```sh
go mod download
```
**Run the Server**
```sh
go run .
```
 
### API Documentation
Swagger documentation is available at:
 
 
```bash
http://localhost:{your_server_port}/swagger/index.html
```
 
## License
Distributed under the MIT license. See [License](https://opensource.org/license/mit) for more information
 