npx tailwindcss -i input.css -o assets/css/styles.css --minify

go run pkg/database/migration/migration.go

go run main.go