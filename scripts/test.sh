#!/usr/bin/env bash

cd `dirname $0`
cd ..
cd bin

for((i=1;i<=9;i++));do

httpPort=$(expr $i \+ 8000);
rpcPort=$(expr $i \+ 9000);

rm -rf config/riff.yml && ./riff run -name node$i -rpc :$rpcPort -http :$httpPort -join 192.168.1.8:9001,192.168.1.8:9002,192.168.1.8:9003 &
#sleep 1

done


#rm -rf config && go run *.go start -name node5 -rpc :8634 -http :8614 -join 192.168.1.220:8630,192.168.1.220:8631,192.168.1.220:8632
