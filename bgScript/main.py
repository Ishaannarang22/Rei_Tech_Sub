import time
import glob
import os
import datetime
import requests

API_KEY = "LlWlRvYlJNihsd8SGZZK"
API_ENDPOINT = "https://0.0.0.0/u/time"

alert_time = datetime.timedelta(seconds=60)
path = "../static/user"

def alert(filep):
    data = {'api_key': API_KEY, 'file': filep}
    r = requests.post(url = API_ENDPOINT, data = data)

while True:
    files = glob.glob(path + "/**/*.wav", recursive=True)
    timenow = datetime.datetime.now()
    for i in files:
        filename = os.path.basename(i)
        filename = filename[:len(filename)-4]
        date_time_obj = datetime.datetime.strptime(filename, '%Y-%m-%d,%H:%M')
        time_left = date_time_obj - timenow
        if time_left < alert_time:
            alert(i)
    time.sleep(10)


