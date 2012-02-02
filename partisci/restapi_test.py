import json
import urlparse
import requests

endpoint = "http://localhost:7777/api/v1/"

def test_get_server_info():
    url = urlparse.urljoin(endpoint, '_partisci')
    print url
    response = requests.get(url)
    print response
    print response.content
    info = json.loads(response.content)
    print info
    assert 'version' in info
