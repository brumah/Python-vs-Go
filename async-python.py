from aiohttp import web
import queries.queries as queries
import time

async def index(request):
    with open("index.html", "rb") as file:
        content = file.read()
        return web.Response(body=content, content_type='text/html')

latency_values = set()

async def get_query(request):
    global latency_values
    id = request.rel_url.query.get("id", "")
    start = time.time()
    
    if id == "1":
        text = queries.blue_button()
    elif id == "2":
        text = queries.green_button()
    elif id == "3":
        text = queries.red_button()
    elif id == "4":
        text = queries.purple_button()
    elif id == "5":
        text = queries.clear_button()
    elif id == "6":
        text = queries.export_button(list(latency_values))
        latency_values = set()
    
    print(time.time() - start)
    latency_values.add(round((float(time.time() - start)) * 1000,2))
    return web.Response(text=text, content_type='text/html')

app = web.Application()
app.router.add_get('/', index)
app.router.add_get('/getquery', get_query)

web.run_app(app, host='localhost', port=8080)