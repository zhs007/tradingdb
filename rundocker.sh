docker container stop tradingdb
docker container rm tradingdb
docker run -d -p 7788:7788 -p 7789:7789 --name tradingdb -v $PWD/dat:/home/tradingdb/dat -v $PWD/cfg:/home/tradingdb/cfg tradingdb