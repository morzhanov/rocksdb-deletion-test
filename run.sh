count=50
for i in $(seq $count); do
    ./rocksdb-test
    ls -la /home/vagrant/test-data/
done
