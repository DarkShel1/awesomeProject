create table golang_api.user_balance
(
    user_id      int          not null,
    user_balance INT UNSIGNED not null,
    constraint user_balance_pk
        primary key (user_id)
)
    charset = utf8;

create table golang_api.reservation
(
    user_id  int          not null,
    task_id  int          not null,
    order_id int          not null,
    amount   int unsigned not null,
    constraint reservation_pk
        primary key (order_id),
    constraint reservation_user_balance_user_id_fk
        foreign key (user_id) references golang_api.user_balance (id)
)
    charset = utf8;