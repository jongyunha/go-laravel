module app

go 1.21.3

replace github.com/jongyunha/celeritas => ../celeritas

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/jongyunha/celeritas v0.0.0-00010101000000-000000000000
)

require github.com/joho/godotenv v1.5.1 // indirect
