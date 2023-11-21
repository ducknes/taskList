select
    t.id as Id,
    t.stasus as Status,
    t.message as Message
from
    tasks t
where
    t.id = $1;