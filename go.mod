module server

go 1.22.4

replace localhost.com/models => ./models

replace localhost.com/db => ./db

replace localhost.com/controllers => ./controllers

replace localhost.com/routes => ./routes

require (
	github.com/lib/pq v1.10.9 // indirect
	localhost.com/controllers v0.0.0-00010101000000-000000000000 // indirect
	localhost.com/db v0.0.0-00010101000000-000000000000 // indirect
	localhost.com/models v0.0.0-00010101000000-000000000000 // indirect
	localhost.com/routes v0.0.0-00010101000000-000000000000
)
