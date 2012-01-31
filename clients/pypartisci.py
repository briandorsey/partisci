import json
import time
import socket

__version__ = "0.1demo"

def update_template():
    return dict(
        name="",
        version="",
        host="",
        instance=0)


def send_update(update, host, port):
    data = json.dumps(update)
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.connect((host, port))
    s.send(data)
    s.close()
    return

if __name__ == '__main__':
    update = update_template()
    update["name"] = 'python_client_demo'
    update["version"] = __version__
    update["host"] = socket.gethostname()
    while True:
        print "Sending update"
        send_update(update, 'localhost', 7777)
        time.sleep(3)

