echo "Creating release ..."


create_release() {
    AUTH_TOKEN="$1"
    OWNER_NAME="$2"
    REPO_NAME="$3"
    VERSION_NUMBER="$4"
    BRANCH_NAME="$5"
    DESCRIPTION="$6"
    IS_DRAFT="$7"
    IS_PRERELEASE="$8"
    ARTIFACTS="$9"


    curl -H "Authorization: token ${AUTH_TOKEN}" \
    --header "Content-Type: application/json" \
    --request POST \
    --data '{
    "tag_name": "'${VERSION_NUMBER}'",
    "target_commitish": "'${BRANCH_NAME}'",
    "name": "'${VERSION_NUMBER}'",
    "body": "'${DESCRIPTION}'",
    "draft": '${IS_DRAFT}',
    "prerelease": '${IS_PRERELEASE}'
    }' \
    https://api.github.com/repos/${OWNER_NAME}/${REPO_NAME}/releases > create_release_output.txt

    cat create_release_output.txt

    CREATE_MESSAGE=$(cat create_release_output.txt | jq --raw-output '.message')

    if [ "$CREATE_MESSAGE" ];then
        echo "[Error] while creating"
        cat create_release_output.txt
        echo "Cleaning up"
        cleanup_create
        exit 1
    fi

    RELEASE_ENDPOINT=$(cat create_release_output.txt | jq --raw-output '.upload_url')

    echo $RELEASE_ENDPOINT


    set -f                      # avoid globbing (expansion of *).
    array=(${ARTIFACTS//,/ })
    for i in "${!array[@]}"
    do
        echo "Uploading artifact ${array[i]}"
        upload_asset "${array[i]}" "${RELEASE_ENDPOINT}"
    done


    echo "Cleaning up create output"
    cleanup_create
}

upload_asset(){
    FILE_URL="$1"
    RELEASE_ENDPOINT="$2"

    FILE_NAME="${FILE_URL##*/}"

    echo "Uploading assets to ${RELEASE_ENDPOINT}=${FILE_NAME} ..."

    curl -H "Authorization: token ${AUTH_TOKEN}" \
    -F upload=@${FILE_NAME} ${FILE_URL} \
    --request POST \
    ${RELEASE_ENDPOINT}=${FILE_NAME} > upload_assets_output${FILE_NAME}.txt

    cat upload_assets_output${FILE_NAME}.txt

    echo "Cleaning up assets output"
    rm upload_assets_output${FILE_NAME}.txt
}



cleanup_create(){
    rm create_release_output.txt
}

delete_release() {
    AUTH_TOKEN="$1"
    OWNER_NAME="$2"
    REPO_NAME="$3"
    VERSION_NUMBER="$4"
    echo "Deleting release ${VERSION_NUMBER}..."
    echo "Getting releases https://api.github.com/repos/${OWNER_NAME}/${REPO_NAME}/releases ..."

    curl -H "Authorization: token ${AUTH_TOKEN}" \
    https://api.github.com/repos/${OWNER_NAME}/${REPO_NAME}/releases > all_releases_output.txt

    cat all_releases_output.txt

    RELEASE_ID=$(cat all_releases_output.txt | jq --raw-output '.[] , select(.[].tag_name == "${VERSION_NUMBER}") | .id')
    
    echo deleting release with Id: $RELEASE_ID

    if [ ! "$RELEASE_ID" ];then
        echo "[Error] Empty release Id found no release deleted"
        cleanup_delete
        exit 1
    fi

    curl -H "Authorization: token ${AUTH_TOKEN}" \
    --request DELETE \
    https://api.github.com/repos/${OWNER_NAME}/${REPO_NAME}/releases/${RELEASE_ID} > delete_release_output.txt

    DELETE_MESSAGE=$(cat delete_release_output.txt | jq --raw-output '.message')

    if [ "$DELETE_MESSAGE" ];then
        echo "[Error] while deleting"
        cat delete_release_output.txt
        echo "Cleaning up"
        cleanup_delete
        exit 1
    fi
    
    echo "Cleaning up"
    cleanup_delete

}

cleanup_delete(){
    rm all_releases_output.txt
    rm delete_release_output.txt
}


case "$1" in
    "create")
        create_release "$2" "$3" "$4" "$5" "$6" "$7" "$8" "$9" "${10}"
        ;;
    "delete")
        delete_release "$2" "$3" "$4" "$5"
        ;;
    *)
        echo "No option chosen, pick create or delete release"
esac
