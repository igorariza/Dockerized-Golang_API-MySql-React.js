#!/bin/bash --utf-8
set -e

DB_USER='root';
DB_PASSWORD='admin';
DB_HOST='db';
DB_PORT=3306 ;
DB_NAME='admin';

echo "ALTER DATABASE ${DB_NAME} CHARACTER SET utf8 COLLATE utf8_general_ci;" | mysql --user=$DB_USER --password=$DB_PASSWORD --host=$DB_HOST --port=$DB_PORT $DB_NAME
echo "ALTER TABLE users CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;" | mysql --user=$DB_USER --password=$DB_PASSWORD --host=$DB_HOST --port=$DB_PORT $DB_NAME

TOTAL_PAGES=$(curl -s https://www.easports.com/fifa/ultimate-team/api/fut/item | jq -r .totalPages)
i=0
k=0
while [ "$i" -le ${TOTAL_PAGES} ]; do
        i=$((i+1));
                
        curl --request GET -s "https://www.easports.com/fifa/ultimate-team/api/fut/item?page=${i}" | jq -r '.items' > res.json
        
        jq -c '.[]' res.json | while read j; do                   
                frtname=$(echo ${j}  | jq -r '.firstName')
                lstname=$(echo ${j}  | jq -r '.lastName')
                postn=$(echo ${j}  | jq -r '.position')
                natn=$(echo ${j}  | jq -r '.nation.abbrName')
                clb=$(echo ${j}  | jq -r '.club.abbrName')
                
                echo "INSERT INTO users (firstname,lastname,position,nation,club,page) VALUES ('${frtname//\'/\'}','${lstname//\'/\'}','${postn}','${natn//\'/\'}','${clb//\'/\'}','${i}')"  | mysql --user=$DB_USER --password=$DB_PASSWORD --host=$DB_HOST --port=$DB_PORT $DB_NAME
                
        done    
done

exec "$@"