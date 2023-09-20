-- +goose Up
-- +goose StatementBegin
CREATE TABLE  history
(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    user_id VARCHAR(100) NOT NULL,
    data JSON NOT NULL,
    event_date TIMESTAMP NOT NULL,

    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS history;
-- +goose StatementEnd
