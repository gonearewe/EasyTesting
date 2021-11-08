import http.client

import requests

import api


def login(server_addr: str, teacher_id: str, password: str) -> (bool, str):
    api.SERVER_ADDR = server_addr
    try:
        response = requests.get(server_addr, params={'teacher_id': teacher_id, 'password': password})
    except requests.exceptions.RequestException as e:
        return False, f'网络异常:\n{e.strerror}'
    if response.status_code == http.HTTPStatus.OK:
        api.AUTHENTICATION_TOKEN = response.text
        return True, ''
    else:
        return False, '认证失败'
