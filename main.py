from http.server import SimpleHTTPRequestHandler, HTTPServer
from urllib.parse import urlparse, parse_qs
import queries.queries as queries

class CustomHandler(SimpleHTTPRequestHandler):
    def do_GET(self):
        parsed_url = urlparse(self.path)
        query = parse_qs(parsed_url.query)

        if parsed_url.path == '/':
            self.send_response(200)
            self.send_header('Content-type', 'text/html')
            self.end_headers()
            with open("index.html", "rb") as file:
                content = file.read()
                self.wfile.write(content)
                
        elif parsed_url.path == '/getquery':
            id = query.get("id", [""])[0]

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
            self.send_response(200)
            self.send_header('Content-type', 'text/html')
            self.end_headers()
            self.wfile.write(text.encode())


def run_server(handler_class=CustomHandler, server_class=HTTPServer):
    server_address = ('', 8000)
    httpd = server_class(server_address, handler_class)
    
    print(f"Starting server on port {server_address[1]}...")
    httpd.serve_forever()

if __name__ == '__main__':
    run_server()
