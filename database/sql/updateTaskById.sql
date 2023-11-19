update
    tasks t
set
    t.status = $2,
    t.message = $3
where
    t.id = $1;
