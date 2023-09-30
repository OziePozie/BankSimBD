CREATE table IF NOT EXISTS accounts (
                                        account_id SERIAL primary key,
                                        first_Name varchar,
                                        second_Name varchar,
                                        email varchar unique,
                                        password varchar
);
CREATE table IF NOT EXISTS bills (
                                     bill_id SERIAL primary key,
                                     account_id int references accounts (account_id),
    number bigint,
    sum_limit int
    );
CREATE table IF NOT EXISTS cards (
                                     card_id SERIAL primary key,
                                     bill_id int references bills (bill_id),
    number bigint,
    cvv varchar,
    expiration_date timestamp,
    isCardActive bool
    );
CREATE table IF NOT EXISTS history (
                                       history_id SERIAL primary key,
                                       destination_card_id int references cards (card_id),
    arrival_card_id int references cards (card_id),
    date timestamp,
    operation_type varchar,
    sum int

    );