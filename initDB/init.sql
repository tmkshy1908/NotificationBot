
-- ユーザー yamadatarou の存在を確認し、存在しなければ作成
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'yamadatarou') THEN
        CREATE USER yamadatarou WITH PASSWORD '1234';
    END IF;
END
$$;

-- データベース notificationdb の存在を確認し、存在しなければ作成
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'notificationdb') THEN
        CREATE DATABASE notificationdb;
    END IF;
END
$$;


CREATE TABLE usertime(
	id text not null,
	hour integer not null,
	minute integer not null
);
