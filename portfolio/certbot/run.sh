while true; do
    certbot certonly --standalone --preferred-challenges http -d map.g1y.io -d g1y.io
    sleep 1d
done
