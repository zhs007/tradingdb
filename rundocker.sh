docker container stop tradingdb
docker container rm tradingdb
docker run -d -p 7888:7888 -p 7889:7889 --name tradingdb -v $PWD/dat:/home/tradingdb/dat -v $PWD/cfg:/home/tradingdb/cfg tradingdb