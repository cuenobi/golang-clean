BEGIN;

CREATE TABLE IF NOT EXISTS audit_log (
    id BIGSERIAL PRIMARY KEY,
    event_id VARCHAR(64) NOT NULL,
    entity_type VARCHAR(64) NOT NULL,
    entity_id BIGINT NOT NULL,
    entity_name VARCHAR(255),
    action VARCHAR(64) NOT NULL,
    username VARCHAR(255) NOT NULL,
    module VARCHAR(128) NOT NULL,
    ip_address VARCHAR(64),
    user_agent TEXT,
    diff_value JSONB,
    organization_id BIGINT NULL,
    organization_name VARCHAR(255),
    occurred_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_audit_log_occurred_at ON audit_log (occurred_at DESC);
CREATE INDEX IF NOT EXISTS idx_audit_log_module ON audit_log (module);
CREATE INDEX IF NOT EXISTS idx_audit_log_action ON audit_log (action);
CREATE INDEX IF NOT EXISTS idx_audit_log_username ON audit_log (username);
CREATE INDEX IF NOT EXISTS idx_audit_log_organization_id ON audit_log (organization_id);
CREATE INDEX IF NOT EXISTS idx_audit_log_entity ON audit_log (entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_audit_log_event_id ON audit_log (event_id);

COMMIT;
