-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS gcode_metadata (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    machine_name TEXT,
    flavor TEXT,
    time_seconds INTEGER,
    filament_used FLOAT,
    layer_height FLOAT,
    bounding_box JSONB,
    settings JSONB,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_on TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_user_id ON "gcode_metadata" (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "gcode_metadata";
-- +goose StatementEnd