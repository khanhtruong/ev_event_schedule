
-- +migrate Up
CREATE TYPE event_types AS ENUM ('in_person', 'virtual', 'test_drive');
CREATE TYPE confirmation_statuses AS ENUM ('pending', 'confirmed', 'declined');

CREATE TABLE events(
	id INT not null primary key,
    title VARCHAR(200),
    description VARCHAR(200),
	location VARCHAR(200),
	type event_types
);

CREATE TABLE user_accounts(
	id INT not null primary key,
	name varchar(200),
	is_vip BOOLEAN
);

CREATE TABLE event_sessions(
	id INT not null primary key,
	event_id INT,
	slot_amount INT,
	start_time INT,
	end_time INT,
	FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE TABLE subscriptions(
	id UUID not null primary KEY,
	user_id INT,
	event_session_id INT,
	local_event_calendar_id VARCHAR(100),
	status confirmation_statuses,
	on_notify_new_event BOOLEAN,
	FOREIGN KEY (user_id) REFERENCES user_accounts(id),
	FOREIGN KEY (event_session_id) REFERENCES event_sessions(id)
);

-- +migrate Down
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS event_sessions;
DROP TABLE IF EXISTS user_accounts;
DROP TABLE IF EXISTS events;
DROP TYPE IF EXISTS event_types;
DROP TYPE IF EXISTS confirmation_statuses;

