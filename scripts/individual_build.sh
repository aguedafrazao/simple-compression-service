service=$1

tag=$2

REPO_PREFIX=animal505

image="${REPO_PREFIX}/$service:$tag"

cd ../src/apps/$service 

docker build -t "${image}" --build-arg ACCESS_TOKEN_USR="${GITHUB_USER}" --build-arg ACCESS_TOKEN_PWD="${GITHUB_TOKEN}" . 

docker push "${image}"
