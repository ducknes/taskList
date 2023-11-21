delete from
    tasks t
where
    t.id = $1
returning
    t.id;