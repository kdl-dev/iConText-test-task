#!/bin/bash
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'icontext_test_task'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE icontext_test_task";