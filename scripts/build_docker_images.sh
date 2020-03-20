if [ $# -lt 1 ];
then
    echo ">>> Usage: ./build-docker-images <env prod|dev>"
    exit 1
fi

givenEnv=$1

case "$givenEnv" in
    "prod") 
        TAG="latest"
    ;;
    
    "dev") 
        TAG="dev"
    ;;
    
    *) 
	echo ">>> ERROR: Invalid param '$givenEnv'"
    ;;
esac

echo "Given tag $TAG"

REPO_PREFIX=animal505

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

while IFS= read -d $'\0' -r dir; do
    svcname="$(basename "${dir}")"
    image="${REPO_PREFIX}/$svcname:$TAG"
    (
       cd "${dir}"
       echo "Building: ${image}"
       docker build -t "${image}" --build-arg ACCESS_TOKEN_USR="${GITHUB_USER}" --build-arg ACCESS_TOKEN_PWD="${GITHUB_TOKEN}" . 

       echo "Pushing: ${image}"
       docker push "${image}"
    )
done < <(find "${SCRIPTDIR}/../src/apps" -mindepth 1 -maxdepth 1 -type d -print0)

echo "All images built and pushed!"

