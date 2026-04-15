-- ============================================================
-- ADMIN SCHEMA: Document Verification
-- Database: public | Prefix: admin_
-- ============================================================

CREATE TYPE admin_owner_type      AS ENUM ('driver', 'vehicle');
CREATE TYPE admin_doc_type        AS ENUM ('license', 'rc', 'insurance', 'permits', 'aadhar', 'pan');
CREATE TYPE admin_document_status AS ENUM ('pending', 'approved', 'rejected', 'expired');

CREATE TABLE IF NOT EXISTS public.admin_documents (
    id          VARCHAR(36)           PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
    owner_id    VARCHAR(36)           NOT NULL,
    owner_type  admin_owner_type      NOT NULL,
    doc_type    admin_doc_type        NOT NULL,
    file_url    TEXT                  NOT NULL,
    status      admin_document_status NOT NULL DEFAULT 'pending',
    reviewer_id VARCHAR(36),
    review_note TEXT,
    reviewed_at TIMESTAMPTZ,
    expires_at  TIMESTAMPTZ,
    created_at  TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ           NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ,
    CONSTRAINT chk_admin_documents_file_url CHECK (length(file_url) >= 1)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.admin_documents
    FOR EACH ROW EXECUTE FUNCTION trigger_set_updated_at();

CREATE INDEX idx_admin_documents_owner     ON public.admin_documents (owner_id, owner_type);
CREATE INDEX idx_admin_documents_composite ON public.admin_documents (owner_id, owner_type, doc_type);
CREATE INDEX idx_admin_documents_status    ON public.admin_documents (status);
CREATE INDEX idx_admin_documents_expiry    ON public.admin_documents (expires_at) WHERE expires_at IS NOT NULL;
CREATE INDEX idx_admin_documents_deleted   ON public.admin_documents (deleted_at) WHERE deleted_at IS NULL;
