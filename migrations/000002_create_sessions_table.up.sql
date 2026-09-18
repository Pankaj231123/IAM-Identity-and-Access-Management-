-- id	UUID	Primary key	Session identifier
-- user_id	UUID	FOREIGN KEY → users(id), NOT NULL	Links session to owner
-- token_hash	TEXT	UNIQUE, NOT NULL	SHA-256 hash of the opaque token — never store the raw token
-- ip_address	TEXT	NULLABLE	For "active devices" listing
-- user_agent	TEXT	NULLABLE	For "active devices" listing
-- expires_at	TIMESTAMPTZ	NOT NULL	Session TTL
-- revoked_at	TIMESTAMPTZ	NULLABLE	NULL = still active
-- created_at	TIMESTAMPTZ	DEFAULT now()	Audit trail
CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash TEXT UNIQUE NOT NULL,
    ip_address TEXT NULL,
    user_agent TEXT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);