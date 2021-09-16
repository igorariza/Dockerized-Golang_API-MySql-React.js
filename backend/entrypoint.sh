#!/bin/sh
set -e
# baseName="admin"
# tableName="users"
# FRT_NAME =""
# LST_NAME =""
# PL_POSITION =""
# PL_NATION=""
# myslqquery=""
# TOTAL_PAGES=$(curl -s https://www.easports.com/fifa/ultimate-team/api/fut/item | jq -r .totalPages)
# for i in $TOTAL_PAGES
#     do
#         ITEMS=$(curl -s https://www.easports.com/fifa/ultimate-team/api/fut/item?page= $i | jq -r ' .items ')
#         for j in ITEMS 
#             do 
#                 $FRT_NAME = "${ITEMS}"  | jq -r   ' .[] | .firstName'
#                 $LST_NAME  = "${ITEMS}"  | jq -r   ' .[] | .lastName'
#                 $PL_POSITION  = "${ITEMS}"  | jq -r   ' .[] | .position'
#                 $PL_NATION    = "${ITEMS}"  | jq -r   ' .[] | .nation.abbrName'
                
#                 $myslqquery="USE $baseName;
#                              INSERT INTO $tableName ('firstname', 'lastname', 'position', 'nation' )VALUES(\"$FRT_NAME\",\"$LST_NAME\",\"$PL_POSITION\", ,\"$PL_NATION\");
#                              SELECT * from $tableName;"
#                 mysql -u root -e "$myslqquery"

#         done   
# done

exec "$@"