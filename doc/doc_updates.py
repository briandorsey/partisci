import os
import sys

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "../clients"))
import pypartisci

server, port = "localhost", 7777

apps = ["Demo App A",
        "Demo App B", 
        "Demo App C"]

hosts = ["host1.example.com",
         "host2.example.com"]

for app in apps:
    for host in hosts:
        pypartisci.send_update(server, port, app, "ver", host)
