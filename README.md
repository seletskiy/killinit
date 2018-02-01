# killinit

Ever tried to send kill signal from one container to another in Kubernetes pod?
There is solution.

Just specify `killinit --listen /mnt/kill -- original command` as entrypoint
for container which `original command` should receive kill signal you want to
send.

Mount shared directory in both containers as described there: [1]

Send kill signal from second container by writing kill signal name into `/mnt/kill`:

```
echo HUP > /mnt/kill
```

[1]: https://kubernetes.io/docs/tasks/access-application-cluster/communicate-containers-same-pod-shared-volume/
