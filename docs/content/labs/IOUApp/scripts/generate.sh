if [ ! -z "$1" ]; then
    echo "Usage: ./generate corda|hyperledger yourmodel.json"
    exit 1
fi

if [ ! -z "$2" ]; then
    echo "Usage: ./generate corda|hyperledger yourmodel.json"
    exit 1
fi

dir=$(pwd)
flow=$2

if [ $1 == "corda" ]; then
    dovetail-cli contract generate -b corda -m $flow -v 1.0.0 -t ../artifacts --namespace com.example.iou
else
    export GOPATH=$dir/dovetail:$GOPATH
    dovetail-cli contract generate -b fabric -m $flow -v 1.0.0 -t ../artifacts 
fi