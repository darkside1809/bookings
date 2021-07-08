# Bookings and Reservations

There are tons of hotel's web applications in the world, so I've decided to create one more. In this repository you can find something useful for yourself. Specifically, I've implemented reservation system for hotels (hostels), authentication for each user (client), UI that looks not bad and much more functions that you can find helpful. I would appreciate your feedback.

## Dependencies

- Built in Golang version 1.16
- Uses the [chi router](https://github.com/go-chi/chi) to set routes
- Uses the [alexe dwards SCS](https://github.com/alexedwards/scs/v2) session management
- Uses the [nosurf](https://github.com/justinas/nosurf) for authentication and protection
- Uses the [go validator](https://github.com/asaskevich/govalidator) to validate forms in          server-side
- Uses the [jackc pgx](https://github.com/jackc/pgx)
- Uses the [royal ui admin](https://github.com/BootstrapDash/RoyalUI-Free-Bootstrap-Admin-Template)

## Installation

Make sure that Go is installed on your computer. Type the following command in your terminal

```bash
go get github.com/darkside1809/bookings
```

After it the package is ready to use.

**Import package in your project**

Add following line in your `*.go` file:

```bash
import "github.com/darkside1809/bookings"
```

**Install Soda CLI**

Migration is performed on a database whenever it is necessary to update or revert that database's schema to some newer or older version. To work with database migrations in your project, you need to install Soda CLI.

Get pop package to make easier CRUD operations, run migrations, and build/execute queries.

```bash
go get github.com/gobuffalo/pop/...
```

And install Soda to work with it in the command prompt

```bash
go install github.com/gobuffalo/pop/soda
```

**Soda CLI usage**

Once migrations have been created they can be run with either of the following commands:

```bash
soda migrate
soda migrate up
```

both are identical!

If you want to rollback the last applied migration, use the following command:

```bash
soda migrate down
```

Set up database configuration managed by `database.yml` and locate it at the root of your project.
Here is a sample configuration generated for a new app based on PostgreSQL:

```bash
development:
  dialect: postgres
  database: myapp_development
  user: postgres
  password: postgres
  host: 127.0.0.1
  pool: 5

test:
  url: {{envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"}}

production:
  url: {{envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_production"}}
```

<!-- **Run Application**

To run this application, you need to create `file_name.bat` batch file, UpdateReservation(r models.Reservation) error
	DeleteReservationByID(id int) error
	UpdateProcessedForReservation(id int, processed int) errorand write in names of the files which can be executed. It makes the process of typing in the command line more faster. So write the command down below in the command prompt to start the application or you can just copy `.bat` file from my repository, it's up to you.

```bash
./bookings
``` -->

...and you are all set. Let's Goooo!!!
