package server

import (
	// built in Golang packages
	"net/http"
	// External packages/dependencies
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	// My own packages
	"github.com/darkside1809/bookings/internal/config"
	"github.com/darkside1809/bookings/internal/handlers"
)

// Init set routes with its handlers, multiplexer(middleware, sessionFuncs, secure requests),
// and HTTP methods
func Init(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)
	mux.Get("/download-summary", handlers.Repo.DownloadReservationSummary)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostShowLogin)
	mux.Get("/user/logout", handlers.Repo.Logout)
	mux.Get("/user/signup", handlers.Repo.ShowSignUp)
	mux.Post("/user/signup", handlers.Repo.PostShowSignUp)
	
	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/dashboard", handlers.Repo.AdminDashboard)
		mux.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		mux.Get("/reservations-all", handlers.Repo.AdminAllReservations)
		mux.Get("/users-all", handlers.Repo.AdminAllUsers)

		mux.Get("/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
		mux.Get("/reservations/{src}/{id}", handlers.Repo.AdminShowReservation)
		mux.Post("/reservations/{src}/{id}", handlers.Repo.AdminPostShowReservation)
		mux.Get("/process-reservation/{src}/{id}", handlers.Repo.AdminProcessReservation)
		mux.Get("/delete-reservation/{src}/{id}", handlers.Repo.AdminDeleteReservation)

		mux.Get("/users/{src}/{id}", handlers.Repo.AdminShowUsers)
		mux.Post("/users/{src}/{id}", handlers.Repo.AdminPostShowUsers)
		mux.Get("/delete-user/{src}/{id}", handlers.Repo.AdminDeleteUser)
	})
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	
	return mux
}
