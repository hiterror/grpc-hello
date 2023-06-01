# CA
openssl genrsa -out server.key 2048
openssl req -new -x509 -key server.key -out server.crt -days 36500

#Country Name (2 letter code) []:cn
#State or Province Name (full name) []:fujian
#Locality Name (eg, city) []:xiamen
#Organization Name (eg, company) []:jiujiayi
#Organizational Unit Name (eg, section) []:cw
#Common Name (eg, fully qualified host name) []:
#Email Address []:4.qq.com

# 用CA来签发证书
openssl genpkey -algorithm RSA -out test.key

openssl req -new -nodes -key test.key -out test.csr -days 3650 \
  -subj "/C=cn/OU=myorg/O=mycomp/CN=myname" \
  -config ./openssl.cnf -extensions v3_req

openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key \
 -CAcreateserial -extfile ./openssl.cnf -extensions v3_req