DB_NAME=$1
DB_BRIGADE_USER=$2
DB_BRIGADE_PASSWORD=$3
DB_BRIGADE_CHATS_USER=$4
DB_BRIGADE_CHATS_PASSWORD=$5
su postgres <<EOF
createdb  $DB_NAME;
psql -U postgres -d  $DB_NAME -f ./db/grant_brigade_user.sql -v user_name='DB_BRIGADE_USER' -v user_password='DB_BRIGADE_PASSWORD'
echo "Brigase user 'DB_BRIGADE_USER' created."
psql -U postgres -d  $DB_NAME -f ./db/grant_brigade_chats_user.sql -v user_name='DB_BRIGADE_CHATS_USER' -v user_password='DB_BRIGADE_CHATS_PASSWORD'
echo "Brigade_chats user 'DB_BRIGADE_CHATS_PASSWORD' created."
EOF
