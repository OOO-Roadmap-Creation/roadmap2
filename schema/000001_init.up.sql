CREATE TABLE published_roadmaps
(
    id            serial        not null,
    version       integer       not null,
    visible       boolean       not null default true,
    title         varchar(255)  not null default '',
    description   varchar(1024) not null default '',
    dateOfPublish timestamp     not null default now()
);
