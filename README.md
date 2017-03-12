# nethttp_api
nethttp api demo = net/http + httprouter + websocket

* It's fast
  * based on net/http (go 1.8 is already very fast)
* It's simple
  * few main dependences
  ```
  http: net/http
  router: github.com/julienschmidt/httprouter
  mysql: github.com/go-sql-driver/mysql
  redis: github.com/garyburd/redigo/redis
  logs: github.com/astaxie/beego/logs
  websocket: github.com/gorilla/websocket
  ```
  * just a demo scale, very clean
* It's enough
  * support http and websocket
  * support mysql connection pool
  * support redis connection pool