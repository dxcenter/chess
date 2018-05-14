Installation
============

Just a proof of work (for Debian):

```sh
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
apt update
apt-get install -yt sid nodejs

apt-get install -y golang yarn npm sudo

useradd -m site
su -l site <<EOF
go get github.com/dxcenter/chess
go get github.com/xaionaro/reform/reform
go install github.com/dxcenter/chess
go install github.com/xaionaro/reform/reform
cd /home/site/go/src/github.com/dxcenter/chess
~/go/bin/reform models
cd frontend
yarn
yarn build
EOF

apt-get install -y nginx
rm -f /etc/nginx/sites-enabled/default
ln -s /home/site/go/src/github.com/dxcenter/chess/doc/nginx.conf /etc/nginx/sites-enabled/chess

nginx -s reload

sudo -u site /home/go/bin/chess
```

