import os

config = {
    "golang": {
        "src": os.environ["GO_SRC"],
        "dst": os.environ["GO_DST"],
    },
    "celery": {
        "broker": os.environ["CELERY_BROKER"],
        "results": os.environ["CELERY_BACKEND"],
        "routes": {
            'agent.*': {
                'queue': 'agent'
            }
        }
    },

}
