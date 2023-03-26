# gowgboard

docker run --name gowgboard -p 8080:8080 --privileged \
  --device /dev/net/tun:/dev/net/tun \
  -v /etc/wireguard:/etc/wireguard \
  -d gowgboard
  
  
  curl -G http://localhost:8080/interfaces

