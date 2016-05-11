########################
# Run it like below:
# ./bundle.sh <controller host name / IP> <password for admin user>
########################

rm -f bundle.zip 2> /dev/null
AUTHTOKEN=$(curl --silent --insecure --data '{"username":"'$2'","password":"'$3'"}' \
  https://$1/auth/login | jq --raw-output .auth_token)
curl --insecure --header "Authorization: Bearer $AUTHTOKEN" \
  https://$1/api/clientbundle --output bundle.zip
unzip bundle.zip
source ./env.sh
docker info
