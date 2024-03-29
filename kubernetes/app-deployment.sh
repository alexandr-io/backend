#!/bin/sh

set -e

# /kubernetes is the folder with every k8s config files for every micro services.
echo "\033[0;33mMoving on the /kubernetes folder...\033[0m"
cd /kubernetes || exit

if [ ! -d "configs" ]; then
    # Downloading the config repository
    echo "\033[0;33mDownloading the config repository...\033[0m"
    echo "yes" | svn checkout https://github.com/alexandr-io/configs.git/trunk/  --username="$(printenv SVN_USERNAME)" --password="$(printenv SVN_PASSWORD)" && mv trunk configs
    echo "\033[0;33mDone\033[0m"
fi

hostType='prod'
if [ "$(printenv HOST_TYPE)" = "PREPROD" ]; then
    hostType='preprod'
fi
# Check if the folder with the name of the micro service exist (e.g. auth, user, ...)
echo "\033[0;33mChecking if folder $1 exist...\033[0m"
if [ -d "$1" ]; then
    kubectl delete namespace "$1"
    rm -rf "$1"
fi

# Creating the folder and moving the current directory to the created folder
echo "\033[0;33mCreating folder $1\033[0m"
mkdir "$1"
echo "\033[0;33mMoving in the folder...\033[0m"
cd "$1" || exit

# Since the folder was created, the k8s namespace is not defined
echo "\033[0;33mCreating the k8s namespace ""$1""...\033[0m"
kubectl create namespace "$1"
kubectl create secret docker-registry regcred --docker-server=ghcr.io --docker-password="$(printenv GITHUB_TOKEN)" --docker-username="$(printenv SVN_USERNAME)" -n "$1"

# Checking the environment variable HOST_TYPE to figure out which environment need to be deployed, preprod or prod
# HOST_TYPE  =   PREPROD: preprod environment
#                PROD: prod environment
echo "\033[0;33mChecking if preprod or prod deployment in progress...\033[0m"
if [ $hostType = "preprod" ]; then
    # Find the preprod environment

    # Downloading the preprod k8s deployment configuration files from Github (https://github.com/alexandr-io/backend)
    echo "\033[0;33mFind preprod deployment, downloading preprod ""$1"" deployment folder on the develop branch\033[0m"
    echo "yes" | svn export https://github.com/alexandr-io/backend/branches/develop/microservices/"$1"/deployment/preprod --username="$(printenv SVN_USERNAME)" --password="$(printenv SVN_PASSWORD)"
    echo "\033[0;33mDownload done\033[0m"

    # Moving to the preprod configuration
    echo "\033[0;33mMoving to the preprod folder...\033[0m"
    cd preprod || exit
else
    # Find the prod environment

    # Downloading the prod k8s deployment configuration files from Github (https://github.com/alexandr-io/backend)
    echo "\033[0;33mFind prod deployment, downloading prod ""$1"" deployment file on the master branch\033[0m"
    echo "yes" | svn export https://github.com/alexandr-io/backend/trunk/microservices/"$1"/deployment/prod --username="$(printenv SVN_USERNAME)" --password="$(printenv SVN_PASSWORD)"
    echo "\033[0;33mDownload done\033[0m"

    # Moving to the prod configuration
    echo "\033[0;33mMoving to the prod folder...\033[0m"
    cd prod || exit
fi

# Defining configs on kubernetes for the microservice
echo "\033[0;33mDefining configs on kubernetes for the microservice\033[0m"
grep -v '^ *#' < configs.txt | while IFS= read -r config
do
    # Installing configs one by one
    echo "\033[0;33mInstalling configuration for $config...\033[0m"
    for file in /kubernetes/configs/"$config"/"$hostType"/*.yaml
    do
        echo "\033[0;33mFind file ""$file"" to apply\033[0m"
        sed "s/<namespace>/$1/g" "$file" | kubectl apply -f -
    done
done

for file in ./**/*.yaml
do
    kubectl apply -f "$file"
done
