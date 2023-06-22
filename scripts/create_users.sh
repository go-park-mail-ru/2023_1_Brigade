DB_NAME=$1
DB_BRIGADE_USER=$2
DB_BRIGADE_PASSWORD=$3
DB_BRIGADE_CHATS_USER=$4
DB_BRIGADE_CHATS_PASSWORD=$5
su postgres <<EOF
createdb  $DB_NAME;
psql -U postgres -d  $DB_NAME -f ./sql/PostreSQLCreateMainUser.sql -v user_name='DB_BRIGADE_USER' -v user_password='DB_BRIGADE_PASSWORD'
echo "Brigase user 'DB_BRIGADE_USER' created."
psql -U postgres -d  $DB_NAME -f ./sql/PostreSQLCreateAuthUser.sql -v user_name='DB_BRIGADE_CHATS_USER' -v user_password='DB_BRIGADE_CHATS_PASSWORD'
echo "Brigade_chats user 'DB_BRIGADE_CHATS_PASSWORD' created."
EOF
