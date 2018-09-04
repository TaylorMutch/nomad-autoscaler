# Development Roadmap

My original motivation for developing this tool was to perform simple autoscaling for Nomad jobs across resources.

The MVP for this tool is to have the following functionality:

* Launch a Nomad job that performs autoscaling based on CPU and Memory usage within a static group of resources (pod-level autoscaling).
* Utilize `meta` tags with a Nomad job "task" to specify upper and lower scaling bounds, minimum and maximum pod counts, and time to wait before scaling up or down a deployment.

Stretch goals for the tool would include:

* Enabling host-level autoscaling with cloud providers, starting with Google Cloud.
* Providing predictive scaling based on historic data over a period of time.
* Providing a basic UI for launching Nomad autoscaling jobs.