# ui_template

0. Copy files from this template to your repositories and remove main.go (example how to expose metrics). Search for `CHANGE_ME` - you will potentially need to adjust your existing app accordingly.
1. [Expose]https://prometheus.io/docs/guides/go-application/ prometheus metrics
   `curl http://localhost:8080/metrics` should return memory, gc, and other default application metrics. You can disable
   your weather metrics or expose them together.
2. Install docker-compose https://docs.docker.com/compose/.
3. Run with `docker-compose up`.
4. Open [grafana UI](localhost:3000) (localhost:3000)
5. You should already have preconfigured [prometheus source](http://localhost:3000/datasources).
6. Create dashboards that show
    - graph for go_memstats_alloc_bytes_total
    - gauge for go_goroutines
7. Have fun
