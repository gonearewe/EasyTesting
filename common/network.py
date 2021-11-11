import http.client
import logging

import requests

_SERVER_ADDR = ""
_AUTHENTICATION_TOKEN = ""
_LOGGER = logging.getLogger(__name__)


def setServerAddr(ip_port: str) -> bool:
    try:
        response = requests.get(ip_port + "/ping")
    except Exception as e:
        _LOGGER.error(e)
        return False
    else:
        _LOGGER.info(f"server addr set to `{ip_port}`: ping success")
        global _SERVER_ADDR
        _SERVER_ADDR = ip_port
        return True


def getAuth(auth_path: str, params):
    if _SERVER_ADDR == "":
        _LOGGER.critical(f"broken state to send request: _SERVER_ADDR is not set")
        return
    try:
        response = requests.get(_SERVER_ADDR + auth_path, params)
    except Exception as e:
        _LOGGER.error(e)
    else:
        if response.status_code == http.HTTPStatus.OK:
            _LOGGER.info("auth success")
            global _AUTHENTICATION_TOKEN
            _AUTHENTICATION_TOKEN = response.text
        else:
            _LOGGER.warning(f"auth failure: http status code {response.status_code}")


def request(method: str, path: str, params):
    if _SERVER_ADDR == "" or _AUTHENTICATION_TOKEN == "":
        _LOGGER.critical(f"broken state to send request: _SERVER_ADDR({_SERVER_ADDR}) and "
                         f"_AUTHENTICATION_TOKEN({_AUTHENTICATION_TOKEN})")
        return None
    try:
        response = requests.request(method, _SERVER_ADDR + path, params, auth=("Authorization", _AUTHENTICATION_TOKEN))
    except Exception as e:
        _LOGGER.error(e)
        return None
    else:
        if response.status_code == http.HTTPStatus.OK:
            _LOGGER.info("request success")
            return response.json()
        else:
            _LOGGER.warning(f"request failure: http status code {response.status_code}")
            return None


def get(path: str, params):
    return request("GET", path, params)


def post(path: str, params):
    return request("POST", path, params)
