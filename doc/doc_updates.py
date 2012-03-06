import os
import sys

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "../clients/python"))
import pypartisci

server, port = "localhost", 7777

apps = ["Demo App A",
        "Demo App B"]

hosts = ["host1.example.com",
         "host2.example.com"]

versions = ["1.0", "2.0"]

for app in apps:
    for i, host in enumerate(hosts):
        pypartisci.send_update(server, port, app, versions[i], host)
