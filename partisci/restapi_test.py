import json
import sys
import urlparse

sys.path.insert(0, "../clients")
import pypartisci

import requests

endpoint = "http://localhost:7777/api/v1/"

def test_get_server_info():
    url = urlparse.urljoin(endpoint, '_partisci/')
    print url
    response = requests.get(url)
    print response
    print response.content
    info = json.loads(response.content)
    print info
    assert 'version' in info

def test_get_app():
    url = urlparse.urljoin(endpoint, 'app/')
    print url
    response = requests.get(url)
    print response
    print response.content
    info = json.loads(response.content)
    print info
    assert 'data' in info
    # todo: send update, assert update in result
    assert False
