
PP=$1
if [ -z "$PP"]
then
    echo package is empty
    exit -1
fi

PACK=${PP##*/}
echo godepgraph $PACK

godepgraph $PP | dot -Tpng -o $PACK.png
if [ ! -e "$PACK.png" ]
then
    echo godepgraph $PACK failed
fi
