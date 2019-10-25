# Rough Notes

pre-req's for this is a docker account and a github account.

Doing docker 101 would help but is not necessary.

To start with the user should fork the repo

The tutorial should be part of CI. Every release of concourse should be tested against the tutorial to ensure it stays up-to-date.

The test-suite should execute every step of the tutorial as a user would.

### Tutorial Plan

- Run Concourse using docker-compose and login with the fly cli
- Run a task using `fly execute` (Hello World)
- Put the hello world task in a pipeline
- Add a time-resource to the hello-world pipeline. hello-world pipeline now auto-triggers!
- Start building a real pipeline.
  - Fork this repo
  - Create a new pipeline with a single git resource. Check that it works.



### Other Topics:

- How to add custom resources to my pipeline?
- How to create a task image?
