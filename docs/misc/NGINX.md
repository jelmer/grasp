# Using NGINX with Grasp

Let's say you have the Grasp server listening on port 9000 and want to serve it on your domain, `yourgrasp.com`.

We can use NGINX to redirect all traffic for a certain domain to our Grasp application by using the `proxy_pass` directive combined with the port Grasp is listening on. 

Create the following file in `/etc/nginx/sites-enabled/yourgrasp.com`

```
server {
	server_name yourgrasp.com;

	location / {
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $remote_addr;
		proxy_set_header Host $host;
		proxy_pass http://127.0.0.1:9000; 
	}
}
```

If you wish to protect your site using a [Let's Encrypt](https://letsencrypt.org/) HTTPS certificate, you can do so using the [Certbot webroot plugin](https://certbot.eff.org/docs/using.html#webroot). 

```
certbot certonly --webroot --webroot-path /var/www/yourgrasp.com -d yourgrasp.com
```

Your `/etc/nginx/sites-enabled/yourgrasp.com` file should be updated accordingly:

```
server {
	listen 443 ssl http2;
	listen [::]:443 ssl http2;

	server_name yourgrasp.com;

	ssl_certificate /path/to/your/fullchain.pem;
	ssl_certificate_key /path/to/your/privkey.pem;

	location /.well-known {
		alias /var/www/yourgrasp.com/.well-known;
	}

	location / {
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $remote_addr;
		proxy_set_header Host $host;
		proxy_pass http://127.0.0.1:9000; 
	}
}
```

The `alias` directive should point to the location where your `--webroot-path` is specified when generating the certificate (with `/.well-known` appended).

### Test NGINX configuration
```
sudo nginx -t
```

### Reload NGINX configuration

```
sudo service nginx reload
```
