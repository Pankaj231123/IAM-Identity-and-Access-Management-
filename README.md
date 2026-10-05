# IAM Service

Go IAM service project.

Current work includes Argon2id password hashing and verification, opaque session token generation and SHA-256 hashing, user and session models, PostgreSQL connection setup, and SQL migrations for users and sessions. The user repository supports creating and finding users, tracking failed logins, and setting or clearing a lockout.

PostgreSQL and Redis are defined in `docker-compose.yml`. Database settings are read from environment variables; see `.env.example`.
