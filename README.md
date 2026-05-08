# MinMaxMacros

MinMaxMacros is a food product comparison web application. It allows users to add food products with nutritional values, price, weight, category, and store information. Users can filter and sort products to find better options, such as products with the best protein-to-calorie ratio.

## Technologies Used

- Vue.js
- TypeScript
- Go
- PostgreSQL
- REST API

## Prerequisites

- Go
- Node.js and npm
- PostgreSQL
- Goose
- sqlc (only needed if regenerating database code)

## Setup

### Backend

1. Navigate to the backend directory:

```bash
cd backend
```

2. Install Go dependencies:

```bash
go mod download
```

3. Create a `.env` file in the backend directory and add:

```env
DB_URL=your_postgresql_connection_string?sslmode=disable
```

4. From ```backend/sql/schema```, run database migrations:

```bash
goose postgres "your_postgresql_connection_string?sslmode=disable" up
```

5. Start the backend:

```bash
go run .
```

### Frontend

1. Navigate to the frontend directory:

```bash
cd frontend
```

2. Install dependencies:

```bash
npm install
```

3. Create a `.env` file in the frontend directory and add:

```env
VITE_API_URL=your_backend_url
```

4. Start the frontend:

```bash
npm run dev
```
