# Scaleble - webapp

Scaleble - webapp is a modern full-stack SaaS foundation designed for scalability, security, and long-term maintainability. It combines a Next.js 15 frontend with a Go backend, a PostgreSQL-ready data layer, and a clean architecture that supports growth from MVP to production workloads.

## Overview

This repository provides a production-oriented starter architecture for building authenticated web applications with:

- Next.js App Router and TypeScript on the frontend
- Go for the backend API
- PostgreSQL as the primary relational database
- Redis-ready integration points for caching and sessions
- JWT-based authentication boundaries with room for refresh tokens and OAuth providers
- Docker-friendly project structure for local and production environments

## Tech Stack

Frontend:

- Next.js 15
- React
- TypeScript
- Tailwind CSS
- TanStack Query
- Zod-ready form validation

Backend:

- Go 1.22+
- Chi HTTP router
- PostgreSQL
- Clean Architecture / Hexagonal Architecture boundaries

Infrastructure:

- Docker and Docker Compose ready structure
- Environment-based configuration
- Migration-first database design

## Repository Structure

```text
frontend/   Next.js App Router application
backend/    Go API service and domain layers
docs/       Architecture notes and implementation guidance
docker/     Containerization assets
scripts/    Utility scripts for local development and automation
```

Key frontend folders:

- `app/` for routes, layouts, and providers
- `components/` for reusable UI and layout components
- `context/` for application state providers
- `hooks/` for reusable client hooks
- `services/` for API client logic
- `lib/` for shared helpers
- `types/` for global TypeScript types

Key backend folders:

- `cmd/api/` for the application entrypoint
- `internal/config/` for environment and app configuration
- `internal/handlers/` for HTTP handlers and request/response flow
- `internal/routes/` for route registration
- `internal/services/` for business logic
- `internal/repository/` for data access contracts
- `internal/database/` for PostgreSQL helpers and migrations

## Getting Started

Prerequisites:

- Node.js 18+ for the frontend toolchain in this workspace
- Go 1.22+
- PostgreSQL for local database development

### Frontend

```bash
cd frontend
npm install
npm run dev
```

The application will run on the default Next.js port and includes App Router pages for the home, login, and dashboard flows.

### Backend

```bash
cd backend
go test ./...
go run ./cmd/api
```

The backend listens on `PORT` from the environment and exposes versioned API routes under `/api/v1`.

## Environment Variables

Backend:

- `APP_ENV` - runtime environment, for example `development` or `production`
- `PORT` - HTTP server port
- `DATABASE_URL` - PostgreSQL connection string
- `JWT_ACCESS_SECRET` - signing secret for access tokens
- `JWT_REFRESH_SECRET` - signing secret for refresh tokens
- `ACCESS_TOKEN_TTL_MINUTES` - access token lifetime
- `REFRESH_TOKEN_TTL_DAYS` - refresh token lifetime
- `CORS_ALLOWED_ORIGINS` - comma-separated list of allowed frontend origins

Frontend:

- `NEXT_PUBLIC_API_URL` - backend API base URL used by the frontend service layer

## Architecture

The codebase is structured to keep concerns isolated:

- Presentation lives in the frontend App Router and reusable UI components.
- Business logic belongs in backend services, not handlers.
- Handlers should remain thin and focus on request validation and response mapping.
- Database access should be isolated behind repository interfaces.
- Infrastructure details such as PostgreSQL pooling and configuration live outside the domain layer.

Additional design details are documented in [docs/architecture.md](docs/architecture.md).

## Current Status

This repository currently provides a working foundation with:

- A built frontend shell with layout, dashboard, and login routes
- A compiling Go backend scaffold with route registration and PostgreSQL adapter support
- A starting PostgreSQL schema migration
- Architecture documentation and implementation roadmap

The next recommended step is to implement the authentication workflow end to end, including user persistence, password hashing, JWT issuance, and refresh token rotation.

## Development Notes

- The frontend uses a shared `Providers` wrapper for React Query and auth state.
- The backend is prepared for secure cookie handling, middleware expansion, and repository-backed services.
- Node and Go dependency versions have been aligned with the current workspace toolchain so local builds succeed consistently.