import copy

DEFAULT_CONFIG = {
    'sub_expr': '',
    'secret': '',
    'sub_group': '',
    'auto_commit': True,
}

extra_conf = {
    'extra':"none"
}



if __name__ == '__main__':
    #print(DEFAULT_CONFIG)
    print(set(DEFAULT_CONFIG).difference(extra_conf))
    local = copy.copy(DEFAULT_CONFIG)
    local["extra"] = "local"
    DEFAULT_CONFIG.clear()
    print(set(DEFAULT_CONFIG),"\n",set(local))
