# nomad-autoscaler

This tool allows you to launch a Nomad job that performs autoscaling for your other Nomad jobs based on CPU and Memory usage.

This tool allows you to:
* Launch an autoscaling job 

# Owner

@taylorm

## Notes / Open Questions

Should this tool execute and register jobs on the nomad server to perform the autoscaling?

Version 1 could simply just be a tool for scaling up or down a job manually

Version 2 could be used to register a daemon that performs autoscaling for a job