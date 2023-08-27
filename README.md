# sysinfo Browser

System info in the browser written in go

## Endpoints:

### CPU:

| Endpoint                    | Info                                                                     |
| --------------------------- | ------------------------------------------------------------------------ |
| /info/cpu/name              | Returns CPU model name                                                   |
| /info/cpu/usage             | Returns CPU usage (Realtime)                                             |
| /info/cpu/usage/X           | Returns CPU usages in last X seconds                                     |
| /info/cpu/average           | Returns CPU usage average (Realtime)                                     |
| /info/cpu/usage/average/X   | Returns CPU usage average in last X seconds                              |
| /info/cpu/usage/cinterval/X | Returns the confidence interval according to the usage in last X seconds |

### GPU:

| Endpoint       | Info             |
| -------------- | ---------------- |
| /info/gpu/name | Returns GPU name |

### Memory:

| Endpoint               | Info                           |
| ---------------------- | ------------------------------ |
| /info/mem/available    | Returns available memory       |
| /info/mem/total        | Returns total memory           |
| /info/mem/usage        | Returns memory usage (in MiB)  |
| /info/mem/usagepercent | Returns memory usage (percent) |

### Disks:

| Endpoint                  | Info                     |
| ------------------------- | ------------------------ |
| /info/disks               | Returns Partitions names |
| /info/disks/PARTNAME/size | Returns partition size   |

### OS Related info:

| Endpoint          | Info                        |
| ----------------- | --------------------------- |
| /info/os/desktop  | Returns Desktop Environment |
| /info/os/hostname | Returns hostname            |
| /info/os/kernel   | Returns kernel name         |
| /info/os/name     | Returns distro name         |
| /info/os/uptime   | Returns uptime              |
