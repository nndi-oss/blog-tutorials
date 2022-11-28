nginx-cpanel
============

This is a simple Docker Compose setup to show how to configure NGINX to route traffic from `/cpanel` path to a Go service.
When I started out web dev stuff, I always found it magical how sites hosted on CPanel routed to the cpanel login when `/cpanel` was added to the site's URL so I decided to show how that magic potentially works with an NGINX config and a simple (no dependencies) go service.

### How to run

```
$ docker compose up --build
```

Go to the page at [http://localhost/](http://localhost/) and then to [http://localhost/cpanel](http://localhost/cpanel) - the former is served by NGINX directly from it's `html` directory and the latter is served by a Go service, proxied by nginx.

