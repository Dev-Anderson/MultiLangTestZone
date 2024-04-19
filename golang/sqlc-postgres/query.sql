-- name: GetAllUsers :many 
SELECT * FROM public.users
ORDER BY id;

-- name: GetUser :one 
select * from public.users 
where public.users.id = $1 limit 1;

-- name: InsertUser :one 
insert into public.users (
  name, email, password
) values (
  $1, $2, $3
) 
returning *; 

-- name: UpdateUser :exec 
update public.users
  set name = $2,
  email = $3,
  password = $4 
where id = $1;

-- name: DeleteUser :exec
delete from public.users 
where id = $1; 

-- name: GetAllUFs :many
select * from public.ufs
order by id; 

-- name: GetCidadePorUF :many
select cidades.nome, cidades.codigo_ibge, ufs.sigla from cidades 
inner join ufs on cidades.uf_id = ufs.id
where ufs.sigla = $1; 