const server = Bun.serve({
  port: 3000,
  async fetch(req) {
    const url = new URL(req.url);

    if ( url.pathname === '/' ) {
      return new Response(
        await Bun.file('index.html').bytes(),
        {
          headers: {"Content-Type": "text/html"}
        }
      );
    } else if (url.pathname === '/app.js') {
      return new Response(
        await Bun.file('app.js').bytes(),
        {
          headers: {"Content-Type": "text/javascript"}
        }
      )
    }
    console.log(url)
    return new Response('404');
  },
});

console.log(`Listening on http://localhost:${server.port} ...`);