# Updating Grasp to the latest version

To update your existing Grasp installation to the latest version, first rename your existing Grasp installation so that we can move the new version in its place.

```
mv /usr/local/bin/grasp /usr/local/bin/grasp-old
```

Then, [download the latest release archive suitable for your system architecture from the releases page](https://github.com/jelmer/grasp/releases/latest) and place it in `/usr/local/bin`.

```
tar -C /usr/local/bin -xzf grasp_$VERSION_$OS_$ARCH.tar.gz
chmod +x /usr/local/bin/grasp
``` 

If you now run `grasp --version`, you should see that your system is running the latest version. 

```
$ grasp --version
Grasp version 1.0.0
```


### Restarting your Grasp web server

To start serving up the updated Grasp web application, you will have to restart the Grasp process that is running the web server.

If you've followed the [installation instructions](Installation%20instructions.md) then you are using Systemd to manage the Grasp process. Run `systemctl restart <your-grasp-service>` to restart it.

```
systemctl restart my-grasp-site
```

Alternatively, kill all running Grasp process by issuing the following command.

```
pkill grasp
```
