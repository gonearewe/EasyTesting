import http.client
import logging

import requests

_SERVER_ADDR = ""
_AUTHENTICATION_TOKEN = ""
_LOGGER = logging.getLogger(__name__)


def setServerAddr(ip_port: str) -> bool:
    try:
        response = requests.get(f"http://{ip_port}/ping")
    except Exception as e:
        _LOGGER.error(e)
        return False
    else:
        _LOGGER.info(f"server addr set to `{ip_port}`: ping success")
        global _SERVER_ADDR
        _SERVER_ADDR = ip_port
        return True


def getAuth(auth_path: str, params) -> bool:
    if _SERVER_ADDR == "":
        _LOGGER.critical(f"broken state to send request: _SERVER_ADDR is not set")
        return False
    try:
        response = requests.get("http://" + _SERVER_ADDR + auth_path, params)
    except Exception as e:
        _LOGGER.error(e)
    else:
        if response.status_code == http.HTTPStatus.OK:
            _LOGGER.info("auth success")
            global _AUTHENTICATION_TOKEN
            _AUTHENTICATION_TOKEN = response.json()["token"]
            return True
        else:
            _LOGGER.warning(f"auth failure: http status code {response.status_code}")
    return False


def request(method: str, path: str, params) -> bool:
    if _SERVER_ADDR == "" or _AUTHENTICATION_TOKEN == "":
        _LOGGER.critical(f"broken state to send request: _SERVER_ADDR({_SERVER_ADDR}) and "
                         f"_AUTHENTICATION_TOKEN({_AUTHENTICATION_TOKEN})")
        return False
    try:
        response = requests.request(method, url="http://" + _SERVER_ADDR + path,
                                    params=params, headers={"Authorization": "Bearer " + _AUTHENTICATION_TOKEN})
    except Exception as e:
        _LOGGER.error(e)
        return False
    else:
        if response.status_code == http.HTTPStatus.OK:
            _LOGGER.info("request success")
            return True
        else:
            _LOGGER.warning(f"request failure: http status code {response.status_code}")
            return False


def get(path: str, params):
    if _SERVER_ADDR == "" or _AUTHENTICATION_TOKEN == "":
        _LOGGER.critical(f"broken state to send request: _SERVER_ADDR({_SERVER_ADDR}) and "
                         f"_AUTHENTICATION_TOKEN({_AUTHENTICATION_TOKEN})")
        return None
    try:
        response = requests.request("GET", url="http://" + _SERVER_ADDR + path,
                                    params=params, headers={"Authorization": "Bearer " + _AUTHENTICATION_TOKEN})
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

def post(path: str, params):
    return request("POST", path, params)


def put(path: str, params):
    return request("PUT", path, params)


def delete(path: str, params):
    return request("DELETE", path, params)
