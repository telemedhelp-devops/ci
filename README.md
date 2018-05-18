Installation
============

Just a proof of work (for Debian):

```sh
apt-get update
apt-get install apt-transport-https
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
apt-get update
apt-get install -yt sid nodejs golang npm
apt-get install -y sudo yarn

useradd -m site
su -l site <<EOF
go get gitlab.telemed.help/devops/ci
go get gihub.com/xaionaro/reform/reform
go install gitlab.telemed.help/devops/ci
go install github.com/xaionaro/reform/reform
cd /home/site/go/src/gitlab.telemed.help/devops/ci
~/go/bin/reform models
cd frontend
yarn
yarn build
EOF

apt-get install -y nginx
rm -f /etc/nginx/sites-enabled/default
ln -s /home/site/go/src/gitlab.telemed.help/devops/ci/doc/nginx.conf /etc/nginx/sites-enabled/ci

nginx -s reload

sudo -u site /home/site/go/bin/ci
```
