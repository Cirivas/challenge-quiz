#!/bin/bash
HOST=$1
PORT=$2

FILES=$(ls ./setup/*.json | sort -n -t _ -k 2)

while : ; do
    echo "[$(date)] Checking connection to REDIS... 😴"

    CHECK=$(redis-cli -h ${HOST} -p ${PORT} PING)
    if [[ $CHECK == *"PONG"* ]]
    then
        break
    fi
    sleep 1
done

echo "[$(date)] Connected! 🤙"

for AFILE in ${FILES[@]}
do
    echo -e "[$(date)] Processing \t$AFILE"
    KEY=`echo $AFILE | sed -n 's/.*\/\(.*\).json/\1/p'`
    echo -e "[$(date)] Using KEY $KEY"
    redis-cli -h ${HOST} -p ${PORT} -x JSON.SET quiz:${KEY} $ < $AFILE
    echo -e "[$(date)] Done \t\t$AFILE"
done

