-- id	UUID	Primary key, default random UUID	Avoids exposing sequential user counts
-- email	TEXT	UNIQUE, NOT NULL	Login identifier
-- password_hash	TEXT	NOT NULL	Never store plaintext
-- failed_login_count	INTEGER	DEFAULT 0	Powers account lockout
-- locked_until	TIMESTAMPTZ	NULLABLE	NULL = not locked
-- created_at	TIMESTAMPTZ	DEFAULT now()	Audit trail
-- updated_at	TIMESTAMPTZ	DEFAULT now()	Audit trail
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    failed_login_count INTEGER DEFAULT 0,
    locked_until TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);