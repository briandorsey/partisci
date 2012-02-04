import json
import sys
import urlparse

sys.path.insert(0, "../clients")
import pypartisci

import requests

server, port = "localhost", 7777
endpoint = "http://%s:%s/api/v1/" % (server, port)

def test_get_server_info():
    url = urlparse.urljoin(endpoint, "_partisci/")
    print url
    response = requests.get(url)
    print response
    print response.content
    info = json.loads(response.content)
    print info
    assert "version" in info

def test_get_app():
    url = urlparse.urljoin(endpoint, "app/")
    print url
    apps = ["_zz_" + str(i) for i in range(5)]
    print "apps:", apps
    for app in apps:
        pypartisci.send_update(server, port, app, "ver")

    response = requests.get(url)
    print response
    print response.content
    info = json.loads(response.content)
    print info
    assert "data" in info
    names = set(v["name"] for v in info["data"]) 
    for app in apps:
        assert app in names
