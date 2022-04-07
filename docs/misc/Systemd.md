# Managing the Grasp process with Systemd

To run Grasp as a service (so it keeps on running in the background and is automatically restarted in case of a server reboot) on Ubuntu 16.04 or later, first ensure you have the `grasp` binary installed and in your `$PATH` so that the command exists.

Then, create a new service config file in the `/etc/systemd/system/` directory.

Example file: `/etc/systemd/system/grasp.service`

The file should have the following contents, with `$USER` substituted with your actual username.

```
[Unit]
Description=Starts the grasp server
Requires=network.target
After=network.target

[Service]
Type=simple
User=$USER
Restart=always
RestartSec=6
WorkingDirectory=/etc/grasp # (or where grasp should store its files)
ExecStart=grasp server

[Install]
WantedBy=multi-user.target
```

Save the file and run `sudo systemctl daemon-reload` to load the changes from disk. 

Then, run `sudo systemctl enable grasp` to start the service whenever the system boots.

### Starting or stopping the Grasp service manually
```
sudo systemctl start grasp
sudo systemctl stop grasp
```

### Using a custom configuration file

If you want to [modify the configuration values for your Grasp service](../Configuration.md), then change the line starting with `ExecStart=...` to include the path to your configuration file.

For example, if you have a configuration file `/home/john/grasp.env` then the line should look like this:

```
ExecStart=grasp --config=/home/john/grasp.env server --addr=:9000
```

#### Start Grasp automatically at boot
```
sudo systemctl enable grasp
```

#### Stop Grasp from starting at boot

```
sudo systemctl disable grasp
```
