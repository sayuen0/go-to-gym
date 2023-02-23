#!/bin/bash

prefix=$(date "+%Y%m%d%H%M")

echo "Enter the name of the SQL file:"
read filename

up_file="$prefix"_"$filename"".up.sql"
down_file="$prefix"_"$filename"".down.sql"

touch $up_file
touch $down_file

echo "Created files $up_file and $down_file."