import logging

# Logging environment
stdlog = logging.getLogger("reternal-agent")
handler = logging.StreamHandler()
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)
stdlog.addHandler(handler)
stdlog.setLevel(logging.INFO)
