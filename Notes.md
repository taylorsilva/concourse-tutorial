# Rough Notes

pre-req's for this is a docker account and a github account.

Doing docker 101 would help but is not necessary.

Can docker-compose run concourse on windows? YES IT DOES

To start with the user should fork the repo

The tutorial should be part of CI. Every release of concourse should be tested against the tutorial to ensure it stays up-to-date.

The test-suite should execute every step of the tutorial as a user would.

How do you have concourse pull images from places other than docker hub?

- The basics
  - start with a pipeline that's a single job that does everything
    - introduce **Gets** that download you code
      - Intro resources later?
    - Introduce **Tasks** that run your unit tests
    - Introduce **Puts** 