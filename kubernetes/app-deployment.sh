#!/bin/sh


echo -e "\033[0;33mPrint env...\033[0m"
env
echo -e "\033[0;33mMoving on the /kubernetes folder...\033[0m"
cd /kubernetes
echo -e "\033[0;33mChecking if folder $1 exist...\033[0m"
if [ ! -d "$1" ]; then
    echo -e "\033[0;33mFolder $1 does not exist, creating...\033[0m"
    mkdir $1
    echo -e "\033[0;33mMoving in the folder...\033[0m"
    cd $1
    echo -e "\033[0;33mCreating the k8s namespace $1...\033[0m"
    kubectl create namespace $1
    echo -e "\033[0;33mChecking if preprod or prod deployment in progress...\033[0m"
    if [ $(printenv HOST_TYPE) = "PREPROD" ]; then
        echo -e "\033[0;33mFind preprod deployment, downloading preprod $1 deployment folder on the develop branch\033[0m"
        echo "yes" | svn export https://github.com/alexandr-io/backend/branches/develop/microservices/$1/deployment/preprod --username=$(printenv SVN_USERNAME) --password=$(printenv SVN_PASSWORD)
    else
        echo -e "\033[0;33mFind prod deployment, downloading preprod $1 deployment file on the master branch\033[0m"
        echo "yes" | svn export https://github.com/alexandr-io/backend/trunk/microservices/$1/deployment/prod --username=$(printenv SVN_USERNAME) --password=$(printenv SVN_PASSWORD)
    fi
    echo -e "\033[0;33mDownload done\033[0m"
else
    echo -e "\033[0;33mFolder $1 exist, moving in the folder...\033[0m"
    cd $1
    echo -e "\033[0;33mDeleting old application...\033[0m"
    kubectl delete -f Deployment.yaml
    echo -e "\033[0;33mCreating new application...\033[0m"
    kubectl apply -f Deployment.yaml
    echo -e "\033[0;33mApplication running.\033[0m"
fi
