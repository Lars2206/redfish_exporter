# redfish_exporter

Docker Image based on https://github.com/jenningsloy318/redfish_exporter.
For more Information please look at her Git Repository


##Configurtion

Published Port: 9610

The Config has to be mounted under /etc/redfish_exporter/redfish_exporter.yml

An example configure given as an example:
```yaml
hosts:
  10.36.48.24:
    username: admin
    password: pass
  default:
    username: admin
    password: pass
```
Note that the ```default`` entry is useful as it avoids an error
condition.


