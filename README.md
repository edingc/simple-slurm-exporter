# simple-slurm-json-exporter
Exposes JSON output of Slurm commands via basic HTTP server for scraping.

There are several good Slurm Prometheus exporters already out there. However, I needed more generic data returned that were not necessarily counters.

This simple Go application exposes the output of select Slurm commands (at this time, only squeue is supported) via JSON. This is done by simply using the "--json" flag and returning the output to the HTTP process for scraping.

I use this in conjunction with Grafana's Infinity Data Source to display statistics on dashboards.
