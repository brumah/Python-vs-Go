from aiohttp import web
import queries.queries as queries

async def index(request):
    with open("index.html", "rb") as file:
        content = file.read()
        return web.Response(body=content, content_type='text/html')

async def get_query(request):
    id = request.rel_url.query.get("id", "")
    
    if id == "1":
        text = queries.blue_button()
    elif id == "2":
        text = queries.green_button()
    elif id == "3":
        text = queries.red_button()
    elif id == "4":
        text = queries.purple_button()
    else:
        text = queries.clear_button()
    
    return web.Response(text=text, content_type='text/html')

app = web.Application()
app.router.add_get('/', index)
app.router.add_get('/getquery', get_query)

web.run_app(app, host='localhost', port=8080)