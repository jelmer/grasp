# Installation instructions for Grasp

To install Grasp on your server: 

1. [Download the latest Grasp release](https://github.com/jelmer/grasp/releases) suitable for your platform.
2. Extract the archive to `/usr/local/bin`

```sh
tar -C /usr/local/bin -xzf grasp_$VERSION_$OS_$ARCH.tar.gz
chmod +x /usr/local/bin/grasp
```

Confirm that Grasp is installed properly by running `grasp --version`

```sh
$ grasp --version
Grasp version 1.0.0
```

## Configuring Grasp

> This step is optional. By default, Grasp will use a SQLite database file in the current working directory.

To run the Grasp web server we will need to [configure Grasp](Configuration.md) so that it can connect with your database of choice. 

Let's create a new directory where we can store our configuration file & SQLite database.

```
mkdir ~/my-grasp-site
cd ~/my-grasp-site
```

Then, set the following environment variables:

```
GRASP_SERVER_ADDR=9000
GRASP_GZIP=true
GRASP_DEBUG=true
GRASP_DATABASE_DRIVER="sqlite3"
GRASP_DATABASE_NAME="grasp.db"
GRASP_SECRET="random-secret-string"
```

If you now run `grasp server` then Grasp will start serving up a website on port 9000 using a SQLite database file named `grasp.db`. If that port is exposed then you should now see your Grasp instance running by browsing to `http://server-ip-address-here:9000`.

Check out the [configuration file documentation](Configuration.md) for all possible configuration values, eg if you want to use MySQL or Postgres instead.

## Register your admin user

> This step is required.

To register a user in the Grasp instance we just created, run the following command:

```
grasp user add --email="john@email.com" --password="strong-password"
```

**Note:** if you're running Grasp v1.0.1 or older, the command is `grasp register --email="john@email.com" --password="strong-password"`

## Using NGINX with Grasp

We recommend using NGINX with Grasp, as it simplifies running multiple sites from the same server and handling SSL certificates with LetsEncrypt.

Create a new file in `/etc/nginx/sites-enabled/my-grasp-site` with the following contents. Replace `my-grasp-site.com` with the domain you would like to use for accessing your Grasp installation.

```sh
server {
	server_name my-grasp-site.com;

	location / {
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $remote_addr;
		proxy_set_header Host $host;
		proxy_pass http://127.0.0.1:9000; 
	}
}
```

Test your NGINX configuration and reload NGINX.

```
nginx -t
service nginx reload
```

If you now run `grasp server` again, you should be able to access your Grasp installation by browsing to `http://my-grasp-site.com`.

## Automatically starting Grasp on boot

To ensure the Grasp web server keeps running whenever the system reboots, we should use a process manager. Ubuntu 16.04 and later ship with Systemd.

Create a new file called `/etc/systemd/system/my-grasp-site.service` with the following contents. Replace `$USER` with your actual username.

```
[Unit]
Description=Starts the grasp server
Requires=network.target
After=network.target

[Service]
Type=simple
User=$USER
Restart=always
RestartSec=3
WorkingDirectory=/home/$USER/my-grasp-site
ExecStart=/usr/local/bin/grasp server

[Install]
WantedBy=multi-user.target
```

Reload the Systemd configuration & enable our service so that Grasp is automatically started whenever the system boots.

```
systemctl daemon-reload
systemctl enable my-grasp-site
```

You should now be able to manually start your Grasp web server by issuing the following command.

```
systemctl start my-grasp-site
```

## Tracking snippet

To start tracking pageviews, copy the tracking snippet shown in your Grasp dashboard to all pages of the website you want to track.


### SSL certificate

With [Certbot](https://certbot.eff.org/docs/) for LetsEncrypt installed, adding an SSL certificate to your Grasp installation is as easy as running the following command.

```
certbot --nginx -d my-grasp-site.com
```


