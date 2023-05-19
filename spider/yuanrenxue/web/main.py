import math
import random,time
import pywasm
import requests

def doRequest():
    # Use a breakpoint in the code line below to debug your script.
    count=0
    for page in ["1","2","3","4","5"]:
        m=getM()
        url="https://match.yuanrenxue.cn/api/match/15?m="+m+"&page="+page
        num=request(url)
        count+=num
    print(count)
def request(url):
    headers={"Cookie":"sessionid=sqf1foifk9w5dtwdm5x2kxyjmafh1nur","User-Agent":"yuanrenxue.project"}
    response = requests.get(url,headers=headers)
    num=0
    jsonData = response.json()
    for data in  jsonData['data']:
        num+=data['value']
    return num
def getM():
    t = int(time.time())
    t1 = int(t / 2)
    t2 = int(t / 2 - math.floor(random.random() * 50 + 1))
    wasm = pywasm.load("main.wasm")
    sign = wasm.exec("encode", [t1, t2])
    return f"{sign}|{t1}|{t2}"

# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    doRequest()

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
