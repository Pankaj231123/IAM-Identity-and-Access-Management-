# IAM Service

Go IAM service project.

Current work includes Argon2id password hashing and verification, opaque session token generation and SHA-256 hashing, user and session models, PostgreSQL connection setup, repository errors, and SQL migrations for users, sessions, the session user index, and `last_seen_at`.

PostgreSQL and Redis are defined in `docker-compose.yml`. Database settings are read from environment variables; see `.env.example`.
