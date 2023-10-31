#!/bin/bash
# Load environment variables from .env file
if [ -f .env ]; then
  source .env
fi

# Check if required environment variables are set
if [ -z "$POSTGRES_DB" ] || [ -z "$POSTGRES_USER" ] || [ -z "$POSTGRES_PASSWORD" ]; then
  echo "Required environment variables are not set. Please check your .env file."
  exit 1
fi

# Create the database
psql -h "$POSTGRES_HOST" -p 5433 -U "$POSTGRES_USER" -d "$POSTGRES_PASSWORD" -c "CREATE DATABASE $POSTGRES_DB;"

# Check the exit status of the psql command
if [ $? -eq 0 ]; then
  echo "Database $POSTGRES_DB created successfully."
else
  echo "Error creating database $POSTGRES_DB."
fi

