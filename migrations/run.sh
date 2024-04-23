#!bin/sh

psql -U postgres -d jet-style -f 001_initial_migration_auth.sql
psql -U postgres -d jet-style -f 002_get_users_auth.sql
psql -U postgres -d jet-style -f 003_initial_migration_notes.sql
psql -U postgres -d jet-style -f 004_last_password_change_trigger.sql
psql -U postgres -d jet-style -f 005_get_notes.sql