import os
import sys

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "../clients"))
import pypartisci

server, port = "localhost", 7777

apps = ["Demo App A",
        "Demo App B", 
        "Demo App C"]

for app in apps:
    pypartisci.send_update(server, port, app, "ver")
