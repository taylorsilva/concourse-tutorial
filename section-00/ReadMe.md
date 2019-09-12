# Section 00 - Run Concourse

The goal of this chapter is to accomplish the following:

- Run Concourse locally on your computer
- Install the `fly` cli
- Log into your local Concourse instance

Let's begin!

### Running Concourse

We're going to use the `docker-compose` file from the [Concourse repo](https://github.com/concourse/concourse/) to run Concourse. Please [install docker-compose](https://docs.docker.com/compose/install/) before proceeding.

With `docker-compose` now installed, download the `docker-compose.yml` file from the [Concourse website](https://concourse-ci.org/).

```bash
# this command will download docker-compose.yml into the current directory
curl -O https://concourse-ci.org/docker-compose.yml
```

Start Concourse by running:

```bash
# the -d flag backgrounds docker-compose after the containers start
docker-compose up -d
```

After the `concourse/concourse` and `postgres` images are downloaded from Docker Hub, you can visit [http://localhost:8080](http://localhost:8080) and you should see an empty Concourse dashboard that says "welcome to concourse!".

Congratulations, you have Concourse running locally!

### Download Fly

Fly is the CLI tool used to interact with Concourse. From the homepage you'll see icons/links to the Fly CLI based on your OS. Click the link appropriate for your OS or run:

```bash
# MacOS/Darwin
curl http://localhost:8080/api/v1/cli?arch=amd64&platform=darwin -o fly
# Linux
curl http://localhost:8080/api/v1/cli?arch=amd64&platform=linux -o fly
# Windows
curl http://localhost:8080/api/v1/cli?arch=amd64&platform=windows -o fly.exe
```

Then install `fly` into your path:

```bash
# MacOS/Darwin/Linux
install ./fly /usr/local/bin
# or
mv ./fly /usr/local/bin/ && chmod +x /usr/local/bin/fly
```

For Windows users, you may need to add a new path to your PATH environment variable and place the `fly.exe` CLI into that folder for it to be accessible from your terminal.

You should now be able to type in `fly` and get a the help screen which lists all fly commands.

Congratulations, you've installed fly!