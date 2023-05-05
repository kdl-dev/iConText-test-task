-- via psql
SELECT 'CREATE DATABASE icontext_test_task' 
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'icontext_test_task')\gexec;
