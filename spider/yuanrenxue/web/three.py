import math
import random,time
import pywasm
import requests

mydict=dict()

def doRequest():
    # Use a breakpoint in the code line below to debug your script.
    count=0
    for page in ["1","2","3","4","5"]:
        url="https://match.yuanrenxue.cn/api/match/3?page="+page
        num=request(url)
        count+=num
    print(count)
    print(mydict)
def request(url):
    headers = {
        'content-length': '0',
        'pragma': 'no-cache',
        'cache-control': 'no-cache',
        'sec-ch-ua': '"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"',
        'sec-ch-ua-mobile': '?0',
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36',
        'sec-ch-ua-platform': '"Windows"',
        'accept': '*/*',
        'origin': 'https://match.yuanrenxue.cn',
        'sec-fetch-site': 'same-origin',
        'sec-fetch-mode': 'cors',
        'sec-fetch-dest': 'empty',
        'referer': 'https://match.yuanrenxue.cn/match/3',
        'accept-encoding': 'gzip, deflate, br',
        'accept-language': 'zh-CN,zh;q=0.9',
        'cookie': 'sessionid=m52z58owhzbsvafj2gje7i6mmsy268zu; Hm_lvt_c99546cf032aaa5a679230de9a95c7db=1684465310,1684482638,1684486875,1684718482; Hm_lvt_9bcbda9cbf86757998a2339a0437208e=1684201182,1684465313,1684487896,1684723383; Hm_lpvt_9bcbda9cbf86757998a2339a0437208e=1684723383; Hm_lpvt_c99546cf032aaa5a679230de9a95c7db=1684723624'}
    session = requests.Session()
    session.headers.clear()
    session.headers.update(headers)
    proxy_dict = {
        "https": "https://127.0.0.1:9994/"
    }
    response = session.post("https://match.yuanrenxue.cn/jssm",verify=False,proxies=proxy_dict)
    dict=response.cookies.get_dict()
    header1 = {
        'content-length': '0',
        'pragma': 'no-cache',
        'cache-control': 'no-cache',
        'sec-ch-ua': '"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"',
        'sec-ch-ua-mobile': '?0',
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36',
        'sec-ch-ua-platform': '"Windows"',
        'accept': '*/*',
        'origin': 'https://match.yuanrenxue.cn',
        'sec-fetch-site': 'same-origin',
        'sec-fetch-mode': 'cors',
        'sec-fetch-dest': 'empty',
        'referer': 'https://match.yuanrenxue.cn/match/3',
        'accept-encoding': 'gzip, deflate, br',
        'accept-language': 'zh-CN,zh;q=0.9',
        'sessionid':dict['sessionid'],
    }
    session = requests.Session()
    session.headers.clear()

    session.headers.update(header1)
    response = session.post(url,verify=False,proxies=proxy_dict,cookies=dict)
    num=0
    jsonData = response.json()
    for data in  jsonData['data']:
        strdata=str(data['value'])
        v=mydict.get(strdata)
        if v==None:
            mydict[strdata]=1
        else:
            mydict[strdata]=v+1
        num+=data['value']
        print(strdata)
    return num

# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    mydict=dict()
    doRequest()

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
