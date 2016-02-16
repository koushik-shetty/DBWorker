DO
$body$
BEGIN
	IF NOT EXISTS (
    	SELECT *
    	FROM   pg_catalog.pg_user
    	WHERE  usename = ':admin') THEN

		--create admin role
    	CREATE ROLE :admin WITH SUPERUSER LOGIN PASSWORD ':admin_pass';
   END IF;


	IF NOT EXISTS (
    	SELECT *
    	FROM   pg_catalog.pg_user
    	WHERE  usename = ':user') THEN

		--create user role
    	CREATE ROLE :user WITH LOGIN PASSWORD ':user_pass';
   END IF;
END
$body$;
