import pypartisci

SERVER = "localhost"
PORT = 7777
APP = "Python UDP Example"
VER = "1.0"

if __name__ == "__main__":
    pypartisci.send_update_http(SERVER, PORT, APP, VER)