# andydayton.com

YAPW = Yet Another Personal Website


## Deploy

### On Infomaniak/Openstack:
 - https://docs.infomaniak.cloud/documentation/04.object-storage/030.static-website/
 - https://docs.infomaniak.cloud/documentation/04.object-storage/010.s3/
 - https://andydaytondotcom.b28bee57175648379ec940f55adfe842.s3.pub1.infomaniak.cloud/index.html
 - https://b28bee57175648379ec940f55adfe842.s3.pub1.infomaniak.cloud/andydaytondotcom/index.html

### On pico.sh
 - https://andyinabox-andydaytondotcom.pgs.sh/

## Setup

On OSX there's a namespace client between the `swift` language compiler and the OpenStack `swift` client. Also on my system the
python `bin` dir was not in the path. I addressed both by installing the Swift client and then symlinking it as `swiftclient`

```bash
pip install python-swiftclient
# ...
# WARNING: The script swift is installed in '/Library/Frameworks/Python.framework/Versions/2.7/bin' which is not on PATH.
# ...
sudo ln -s /Library/Frameworks/Python.framework/Versions/2.7/bin/swift /usr/local/bin/swiftclient
```

