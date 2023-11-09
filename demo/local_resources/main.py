import time
import sanic
import psutil

APP = sanic.Sanic(name='Ki_Demo')

@APP.route("/execute", methods=["POST"])
async def handle_execute(req: sanic.request.Request):
    cmd = req.json
    res = {
        'ok': True,
        'pl': {
            'value': await process(cmd)
        },
    }
    
    return sanic.response.json(res)

async def process(cmd):
    print('command -> [%s]' % (cmd))

    match cmd['pm']['key']:
        case 'RESOURCE_CPU':
            return psutil.cpu_percent()
        case 'RESOURCE_TOP':
            return await list_process()

async def list_process():
    result = {}
    data = []
    
    for process in psutil.process_iter(['pid', 'name', 'username', 'cpu_percent', 'memory_percent']):
        process_info = process.info
        pid = process_info['pid']
        name = process_info['name']
        user = process_info['username']
        cpu_percent = process_info['cpu_percent'] or 0.0
        mem_percent = process_info['memory_percent'] or 0.0
    
        value = {
            'pid':         pid,
            'name':        name,
            'user':        user,
            'cpu_percent': cpu_percent,
            'mem_percent': mem_percent,
        }
        
        data.append(value)
    
    data = sorted(data, key=lambda v: v['cpu_percent'], reverse=True)
    while len(data) > 10:
        data = data[:-1]
        
    for value in data:
        result[value['name']] = value['cpu_percent']
        
    return result
    
def main():
    APP.run(
        port=12000)

if __name__ == '__main__':
    main()