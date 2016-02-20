########################
# Run it like below:
# ./bundle.sh <controller host name / IP> <password for admin user>
########################

rm bundle.zip
AUTHTOKEN=$(curl --silent --insecure --data '{"username":"admin","password":"'$2'"}' \
https://$1/auth/login | jq --raw-output .auth_token)
curl --insecure --header "Authorization: Bearer $AUTHTOKEN" \
https://$1/api/clientbundle --output bundle.zip
unzip bundle.zip
. ./env.sh
docker info
