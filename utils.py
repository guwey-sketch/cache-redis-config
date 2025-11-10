from functools import lru_cache
from typing import Any

def load_config(config_file: str) -> dict:
    import toml
    return toml.load(config_file)

def get_redis_url(host: str, port: int, db: int) -> str:
    return f"redis://{host}:{port}/{db}"

@lru_cache(maxsize=None)
def get_config_value(config: dict, key: str) -> Any:
    return config.get(key)