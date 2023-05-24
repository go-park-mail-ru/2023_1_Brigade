package model

type AuthorizedUser struct {
	Id       uint64 `json:"id"       valid:"type(int)"         db:"id"`
	Avatar   string `json:"avatar"                             db:"avatar"`
	Username string `json:"username"                           db:"username"`
	Nickname string `json:"nickname" valid:"nicknameValidator" db:"nickname"`
	Email    string `json:"email"    valid:"emailValidator"    db:"email"`
	Status   string `json:"status"   valid:"type(string)"      db:"status"`
	Password string `json:"password" valid:"passwordValidator" db:"password"`
}

type Contacts struct {
	Contacts []User `json:"contacts"`
}

type User struct {
	Id       uint64 `json:"id"       valid:"type(int)"         db:"id"`
	Username string `json:"username"                           db:"username"`
	Nickname string `json:"nickname" valid:"nicknameValidator" db:"nickname"`
	Email    string `json:"email"    valid:"emailValidator"    db:"email"`
	Status   string `json:"status"   valid:"type(string)"      db:"status"`
	Avatar   string `json:"avatar"                             db:"avatar"`
}

type LoginUser struct {
	Email    string `json:"email"    valid:"emailValidator"    db:"email"`
	Password string `json:"password" valid:"passwordValidator" db:"password"`
}

type RegistrationUser struct {
	Nickname string `json:"nickname" valid:"nicknameValidator" db:"nickname"`
	Email    string `json:"email"    valid:"emailValidator"    db:"email"`
	Password string `json:"password" valid:"passwordValidator" db:"password"`
}

type UpdateUser struct {
	Email           string `json:"email"            valid:"emailValidator"    db:"email"`
	NewAvatarUrl    string `json:"new_avatar_url"                             db:"new_avatar_url"`
	Nickname        string `json:"nickname"         valid:"nicknameValidator" db:"nickname"`
	Status          string `json:"status"           valid:"type(string)"      db:"status"`
	CurrentPassword string `json:"current_password" valid:"passwordValidator" db:"current_password"`
	NewPassword     string `json:"new_password"     valid:"passwordValidator" db:"new_password"`
}

type UserContact struct {
	IdUser    uint64 `json:"id_user"    db:"id_user"`
	IdContact uint64 `json:"id_contact" db:"id_contact"`
}

//map $http_upgrade $connection_upgrade {
//	default upgrade;
//	''      close;
//}
//
//server {
//	listen 80;
//	server_name technogramm;
//
//	return 301 https://$host$request_uri;
//}
//
//#FastCgi cache
//fastcgi_cache_path /tmp/nginx_cache levels=1:2 keys_zone=microcache:10m max_size=500m;
//fastcgi_cache_key "$scheme$request_method$host$request_uri";
//add_header microcache-status $upstream_cache_status;
//
//server {
//	server_name technogramm.ru www.technogramm.ru;
//
//	#GZIP
//	gzip on;
//	gzip_min_length 1000;
//	gzip_comp_level 3;
//
//	gzip_types text/plain;
//	gzip_types text/css;
//	gzip_types application/json;
//	gzip_types application/javascript;
//	gzip_types text/xml;
//	gzip_types application/xml;
//	gzip_types application/xml+rss;
//	gzip_types text/javascript;
//
//	gzip_disable "msie6";
//
//	set $no_cache 0;
//
//	if ($request_method = POST) { set $no_cache 1; }
//
//	if ($query_string != "") { set $no_cache 1; }
//
//	if ($request_uri ~* "wp/admin") { set $no_cache 1; }
//
//	location / {
//		root /var/www/brigade.com/html/2023_1_Brigade/dist;
//		index index.html index.xml;
//		try_files $uri /index.html;
//	}
//
//	location ^~ /avatars/ {
//		root /home/ubuntu/2023_1_Brigade;
//		try_files $uri =404;
//	}
//
//	location /api/v1 {
//		fastcgi_cache microcache;
//		fastcgi_cache_valid 200 60m;
//		fastcgi_cache_bypass $no_cache;
//		fastcgi_no_cache $no_cache;
//
//		proxy_pass http://127.0.0.1:8081;
//		proxy_set_header Host $host;
//		proxy_set_header X-Real-IP $remote_addr;
//		proxy_http_version 1.1;
//		proxy_set_header Upgrade $http_upgrade;
//		proxy_set_header Connection "upgrade";
//		proxy_connect_timeout 7d;
//		proxy_send_timeout 7d;
//		proxy_read_timeout 7d;
//	}
//
//	listen [::]:443 ssl ipv6only=on; # managed by Certbot
//	listen 443 ssl http2; # managed by Certbot
//	ssl_certificate /etc/letsencrypt/live/technogramm.ru/fullchain.pem; # managed by Certbot
//	ssl_certificate_key /etc/letsencrypt/live/technogramm.ru/privkey.pem; # managed by Certbot
//	include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
//	ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
//}
