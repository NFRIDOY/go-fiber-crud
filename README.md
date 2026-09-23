```
go get github.com/gofiber/fiber/v2
go get -u gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
# Books API

Start the application and PostgreSQL:

```sh
docker compose up --build -d
```

Open Swagger UI at http://localhost:3000/swagger/index.html to operate all five
book endpoints. Expand an endpoint, edit its JSON body or book ID, and click
**Execute**. Create a book first and use the returned `ID` for read, update, and
delete requests.

The OpenAPI specification is available at `/swagger/openapi.json`. Requests use
the same host and port as the UI. Swagger UI loads its pinned JavaScript and CSS
from unpkg, so your browser needs internet access.

Edit `docs/openapi.json` when changing API routes or models. The specification and
UI page are embedded in the Go binary; rebuild the app after editing them.

The current update handler ignores empty strings and a price of zero.

Install dependencies manually:

```sh
