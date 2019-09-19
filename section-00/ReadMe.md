# Section 00 - Run Concourse

The goal of this chapter is to accomplish the following:

- Run Concourse locally on your computer
- Install the `fly` cli
- Log into your local Concourse instance

Let's begin!

### Running Concourse

We're going to use the `docker-compose` file from the [Concourse website](https://concourse-ci.org/) to run Concourse. Please [install docker-compose](https://docs.docker.com/compose/install/) before proceeding.

With `docker-compose` installed, download the `docker-compose.yml` file from the [Concourse website](https://concourse-ci.org/).

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

Fly is the CLI tool used to interact with Concourse. From the homepage you'll see icons/links to the fly CLI based on your OS. Click the link appropriate for your OS or run:

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

You should now be able to type  `fly` into your terminal and get a the help screen listing all fly commands.

Congratulations, you've installed fly!

### Login Using Fly

Let's log into our local Concourse using the basic auth user that the docker file is setup with by default. When logging into a Concourse installation for the first time you need to pass in the URL for Concourse and specify a `target` name. The `target` name will be how we reference our local concourse installation in later commands. We'll name our target `local` but you can name it anything you want.

```
$ fly -t local login -c http://localhost:8080 -u test -p test -n main
logging in to team 'main'


target saved
```

This will create a `.flyrc` file. This is a YAML file that contains the following for each target:

- URL for the given target
- Last team logged into for the given target
- Bearer token for the given target

The `.flyrc` file is located at the following locations:

**Linux**: `$HOME/.flyrc`

**Windows**: `%USERPROFILE%\.flyrc` or `%HOMEDRIVE%\%HOMEPATH%\.flyrc`

The `.flyrc` file looks like this:

```yaml
targets:
  local:
    api: http://localhost:8080
    team: main
    token:
      type: Bearer
      value: <token>
```

You can also log into the Concourse web UI using the `test` user. Visit [http://localhost:8080](http://localhost:8080) and click _Login_ in the top-right corner.

### Difference Between Web UI and Fly CLI

The Web UI and fly CLI differ greatly. The Web UI is mostly a read-only interface into Concourse. There are some actions you can take in the web UI that you can also do from the fly CLI:

- Pause/Unpause pipelines and individual jobs
- Start/Abort jobs
- Pin/Unpin resources
- Make pipelines public (visible to users not logged in)

The fly CLI is able to do everything the web UI does and much more. Run `fly` to see the list of all possible commands. `fly` is how you'll create, update, and delete your pipelines.



This is the end of Section-00. Proceed to Section-01.