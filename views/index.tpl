<!DOCTYPE html>

<html>
  <head>
    <meta charset="utf-8" />
    <link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon" />
    <title>Komputeko - prikomputila terminokolekto - {{.Lang}}</title>
    <link rel="stylesheet" href="/static/css/komputeko.css" />
  </head>
  <body>
    <div id="logo">
      <a href="/"><h1><img src="/static/img/komputeko{{.Lang}}.svg" alt="Komputeko" /></h1></a>
      <div id="search">
        <form action="/s/{{.Lang}}">
          <input id="vorto" name="vorto" type="text" />
          <input id="submit" type="submit" value="Serĉi" />
        </form>
      </div>
    </div>
    <div id="sercxrezultoj">
      {{.Topbar}} &nbsp;
    </div>
    <div id="pagecontent">
      {{.Pagecontent}}
    </div>
  </body>
</html>
