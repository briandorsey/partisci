import json
import subprocess
import sys
import urlparse

sys.path.insert(0, "../clients")
import pypartisci

import requests

server, port = "localhost", 7788
endpoint = "http://%s:%s/api/v1/" % (server, port)

class TestPartisci:
    def setup_class(self):
        self.server = subprocess.Popen(["partisci", "--port=%s" % port])

    def teardown_class(self):
        self.server.kill()

    def test_get_server_info(self):
        url = urlparse.urljoin(endpoint, "_partisci/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        assert "version" in info

    def test_get_app(self):
        url = urlparse.urljoin(endpoint, "summary/app/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        # empty result should still be a list. 
        assert list() == info["data"]


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
        names = set(v["app"] for v in info["data"])
        for app in apps:
            assert app in names
