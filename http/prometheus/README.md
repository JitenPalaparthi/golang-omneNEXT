 podman run --name prometheus --network demo-network -p 9090:9090 -d bitnami/prometheus:latest


 podman exec -it prometheus bash 

  podman run -d --name prometheus --network demo-network -v C:\Users\PalaparthiJitendrana\workspace\go-demos\http\prometheus\prometheus.yml:/opt/bitnami/prometheus/conf/prometheus.yml:ro -p 9090:9090 bitnami/prometheus:latest

```bash

  while ($true) {
>>     Write-Host "=== $(Get-Date) ===" -ForegroundColor Cyan
>>     podman ps
>>     Start-Sleep 2
>> }
```
