## minucube

+ Exiting due to DRV_AS_ROOT: The "docker" driver should not be used with root privileges.

```bash
minikube start --force --driver=docker
```

```bash
sudo usermod -aG docker $USER && newgrp docker
minikube start
```

## # 如何使Docker在具有CGroups V2的Fedora OS 上运行？

```bash
# docker this version of runc doesn't work on cgroups v2: unknown

sudo dnf install -y grubby 

sudo grubby --update-kernel = ALL 
  --args =“ systemd.unified_cgroup_hierarchy = 0”
```

# fedora32运行docker

+ 启用旧的CGroups

```bash
sudo grubby --update-kernel=ALL --args="systemd.unified_cgroup_hierarchy=0"
```

+ 在防火墙中将 Docker 列入白名单

```bash
sudo firewall-cmd --permanent --zone=trusted --add-interface=docker0
sudo firewall-cmd --permanent --zone=FedoraWorkstation --add-masquerade
```

+ 步骤 2：安装 Moby

```bash
sudo dnf install moby-engine docker-compose
sudo systemctl enable docker
```

+ 步骤 3：重新启动并测试

```bash
sudo systemctl reboot
sudo docker run hello-world
```

+ 以管理员身份运行

```bash
sudo groupadd docker
sudo usermod -aG docker $USER
```
