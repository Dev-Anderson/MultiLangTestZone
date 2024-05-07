-- name: CreateUser :one
INSERT INTO users (
    id, 
    name, 
    username, 
    email, 
    password, 
    created_at, 
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
    ) RETURNING *;   

-- name: FindUserById :one
SELECT id, username, email, password
FROM public.users
WHERE public.users.id = $1
LIMIT 1;

-- name: FindUserByEmail :one
select * 
from users 
where email = $1
limit 1; 