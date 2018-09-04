# Development Roadmap

My original motivation for developing this tool was to perform simple autoscaling for Nomad jobs across resources.

The MVP for this tool is to have the following functionality:

- [x] Enable manual scaling of a job from a CLI.

- [ ] Enable launching a Nomad job that performs autoscaling based on
    - [ ] CPU usage
    - [ ] Memory Usage
    
- [ ] Utilize `meta` tags with a Nomad job "task" to control:
    - [ ] Upper/lower scaling bounds for the specific metrics (CPU,  Memory)
    - [ ] Minimum/maximum pod counts
    - [ ] Time to wait before scaling up or down a deployment

Stretch goals for the tool would include:

* Enabling host-level autoscaling with cloud providers, starting with Google Cloud.
* Providing predictive scaling based on historic data over a period of time.
* Providing a basic UI for launching Nomad autoscaling jobs.