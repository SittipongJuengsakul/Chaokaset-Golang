CREATE KEYSPACE chaokaset WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;

CREATE TABLE chaokaset.users_by_chaokaset (
    userid uuid,
    username text,
    lname text,
    name text,
    password text,
    pic text,
    prefix text,
    role_user int,
    tel text,
    PRIMARY KEY (userid, username)
)
