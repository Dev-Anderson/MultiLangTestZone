CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL primary key, 
    first_name TEXT not null, 
    last_name TEXT, 
    create_at TIMESTAMP default now()
); 