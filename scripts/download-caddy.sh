DLURL=https://caddyserver.com/download/linux/amd64?license=personal
if [[ "$OSTYPE" == "darwin"* ]]; then
  DLURL=https://caddyserver.com/download/darwin/amd64?license=personal
fi
mkdir ~/tmp
mkdir ~/tmp/caddy
wget -O ~/tmp/caddy/caddy.tar.gz $DLURL
tar -xzf ~/tmp/caddy/caddy.tar.gz --directory ~/tmp/caddy
mv ~/tmp/caddy/caddy ./
rm -rf ~/tmp/caddy
rmdir ~/tmp
