CREATE TABLE IF NOT EXISTS receptions (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            date_time TIMESTAMPTZ NOT NULL,
                            pvz_id UUID NOT NULL REFERENCES pvz(id) ON DELETE CASCADE,
                            status TEXT NOT NULL CHECK (status IN ('in_progress', 'close'))
);