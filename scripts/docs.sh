#!/usr/bin/env bash

# Description: Build script for Project Dovetail documentation
# Author: torresashjian <https://github.com/torresashjian>
# Mod: mtorres@tibco.com
# Last Updated: 2018-11-11

#--- Variables ---
set -e

export PATH=$PATH:/home/travis/.cargo/bin;

workspaceprep() {
    echo "Preparing workspace..."
}

#--- Download and install prerequisites ---
prerequisites() {
    echo "Getting prerequisites..."
}

#--- Get external docs ---
ext_docs() {
    echo "cloning dovetail-contrib..."
    git clone https://github.com/TIBCOSoftware/dovetail-contrib.git 
    #for i in `find dovetail-contrib/activity -name \*.md` ; do filename=$(basename $(dirname $i)); cp $i docs/content/development/webui/activities/$filename.md; done;
    #for i in `find dovetail-contrib/trigger -name \*.md` ; do filename=$(basename $(dirname $i)); cp $i docs/content/development/webui/triggers/$filename.md; done;
    rm -rf ./dovetail-contrib
}


update_page_cli() {
    echo "Getting the docs for the commandline tools"
    #curl -o docs/content/dovetail-cli/dovetail-cli.md https://raw.githubusercontent.com/TIBCOSoftware/dovetail-cli/master/docs/dovetail-cli.md
}

#--- Update contributions page ---
update_page_contrib() {
    echo "Update contributing page"
    cp CONTRIBUTING.md src/ch05-00-contribute.md
    sed -i '1d' src/ch05-00-contribute.md
    sed -i '1i ---' src/ch05-00-contribute.md
    sed -i '1i weight: 9010' src/ch05-00-contribute.md
    sed -i '1i title: Contributing to Project Dovetail' src/ch05-00-contribute.md
    sed -i '1i ---' src/ch05-00-contribute.md
}

#--- Update introduction page ---
update_page_introduction() {
    cp README.md src/ch00-00-introduction.md
    sed -i '1,4d' src/ch00-00-introduction.md
    sed -i '5,17d' src/ch00-00-introduction.md
    sed -i '1i ---' src/ch00-00-introduction.md
    sed -i '1i pre: "<i class=\\"fas fa-home\\" aria-hidden=\\"true\\"></i> "' src/ch00-00-introduction.md
    sed -i '1i weight: 1000' src/ch00-00-introduction.md
    sed -i '1i title: Introduction' src/ch00-00-introduction.md
    sed -i '1i ---' src/ch00-00-introduction.md
    sed -i "s#images/eventhandlers.png#../images/eventhandlers.png#g" src/ch00-00-introduction.md
}

#--- Update page ---
update_page() {
    echo "Updating all pages"
    ext_docs
    update_page_cli
    update_page_contrib
    update_page_introduction
}

#--- Execute build ---
build() {
    echo "Build docs site..."
    mdbook build
}


dobuild(){
    workspaceprep
    prerequisites
    update_pages
    build
}

dobuild