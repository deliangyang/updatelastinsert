create table if not exists `test` (
    `id` int(11) not null auto_increment,
    `val` bigint not null default 0 comment 'value',
    primary key (`id`)
);

insert into `test` (`id`, `val`) values (1, 0);
insert into `test` (`id`, `val`) values (2, 0);