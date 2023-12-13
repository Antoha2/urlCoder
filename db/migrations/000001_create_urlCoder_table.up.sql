create table urlList (
	id serial unique,
	-- token int not null unique,
	long_url varchar(255) not null,
	created_at TIMESTAMP  not null,
    -- expiry_at TIME,
	primary key(id)
	-- foreign key(user_id) references userlist (user_id) on delete cascade on update cascade
)