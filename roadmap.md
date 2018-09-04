# Development Roadmap

My original motivation for developing this tool was to perform simple autoscaling for Nomad jobs across resources.

The MVP for this tool is to have the following functionality:

* Launch a Nomad job that performs autoscaling based on CPU and Memory usage within a static group of resources (pod-level autoscaling).
* Utilize `meta` tags with a Nomad job "task" to specify upper and lower scaling bounds, as well as minimum and maximum pod counts.

Stretch goals for the tool would include:

* Enabling host-level autoscaling with cloud providers, starting with Google Cloud.
* Providing a basic UI for launching Nomad autoscaling jobs.