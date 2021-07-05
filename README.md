# Bookings and Reservations

There are tons of hotel's web applications in the world, so I've decided to create one more. In this repository you can find something useful for yourself. Specifically, I've implemented reservation system for hotels (hostels), authentication for each user (client), and admin authorization as well, UI that looks not bad and much more functions that you can find helpful. I would appreciate your feedback.

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

...and you are all set. Let's GOOO!!!
