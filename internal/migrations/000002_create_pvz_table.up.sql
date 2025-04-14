CREATE TABLE IF NOT EXISTS pvz (
                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                     registration_date TIMESTAMPTZ NOT NULL DEFAULT now(),
                     city TEXT NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань'))
);