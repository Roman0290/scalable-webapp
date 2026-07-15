# Go Project 1

This repository is a full-stack SaaS starter built with Next.js, TypeScript, Go, PostgreSQL, and Docker-ready structure.

## Start Here

- Architecture and roadmap: [docs/architecture.md](docs/architecture.md)
- Frontend: [frontend](frontend)
- Backend: [backend](backend)

## Local Development

Frontend:

```bash
cd frontend
npm run dev
```

Backend:

```bash
cd backend
go run ./cmd/api
```

## Notes

- The frontend uses Next.js App Router and React Query.
- The backend is organized around Clean Architecture / Hexagonal Architecture boundaries.
- The database schema is designed for PostgreSQL with room for Redis-backed caching and token/session management.