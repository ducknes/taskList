insert into
    tasks(status, message)
values
    ($1, $2)
returning id;