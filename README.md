# Booking App Written in Go

This is a **bookings and reservations app** built in **Go** (`v1.21.4`) that uses the following packages:

- [Chi Router](https://pkg.go.dev/github.com/go-chi/chi/v5)
- [SCS Session Management](https://pkg.go.dev/github.com/alexedwards/scs/v2)
- [NoSurf CSRF Protection](https://pkg.go.dev/github.com/justinas/nosurf)
- [GoValidator](https://pkg.go.dev/github.com/asaskevich/govalidator)
- [Soda CLI](https://gobuffalo.io/documentation/database/soda/):
  - Install CLI: `$ brew install gobuffalo/tap/pop`
  - [Config](https://gobuffalo.io/documentation/database/configuration/): 
    - Create DB config file: `$ soda g config`
  - [Migrations](https://gobuffalo.io/documentation/database/migrations/): 
    - Generate migration files: `$ soda generate fizz create_users_table`
    - Apply it: `$ soda migrate`
    - Rollback: `$ soda migrate down`
    - Reset(rollback + apply): `$ soda migrate reset` 
  > **NOTE**: Disable auto-formatting in VSCode for Soda CLI to work properly (`database.yml`).
- [PGX](https://pkg.go.dev/github.com/jackc/pgx/v5): PostgreSQL driver and toolkit:
  - Install: `$ go get github.com/jackc/pgx/v5`
  > **NOTE**: The latest version requires also another package: `$ go get github.com/jackc/pgx/v5/pgxpool@v5.5.0`
