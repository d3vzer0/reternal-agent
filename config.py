import os

config = {
    "golang": {
        "src": os.getenv("GO_SRC", "./agent/src"),
        "dst": os.getenv("GO_DST", "./agent/dist"),
        "pubkey": os.getenv("PUB_KEY", "/opt/reternal/rsa/pub.pem")
    },
    "celery": {
        "broker": os.getenv("CELERY_BROKER", "redis://localhost:6379"),
        "results": os.getenv("CELERY_BACKEND", "redis://localhost:6379"),
        "routes": {
            'agent.*': {
                'queue': 'agent'
            }
        }
    },
}
