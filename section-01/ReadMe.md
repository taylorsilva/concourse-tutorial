# Chapter 1 - Running Tests in CI

Goal: Run your existing **unit tests** in a Concourse task



We're going to run the unit test we have for our `sample-app` in concourse. If you have Go installed already you can run the test locally on your machine from your local copy of this repo.

```
$ cd sample-app
$ go test
PASS
ok  	_/Users/taylor/workspace/concourse-tutorial/sample-app	0.006s
```

We want to do the above in a Concourse task now.

### Creating a Task

A [task](https://concourse-ci.org/tasks.html) is just a container where you execute some script. The task will succeed or fail based on the exit code of the script. A task can have inputs and outputs. Ideally, given the same inputs, your task should always succeed with the same outputs or always fail.

Let's turn "running our unit tests" into a task.

We can start by turning the above steps into a bash script. Let's assume we always run this script from the root of this repo. Call it `run-unit-tests.sh`

```bash
#!/bin/bash

cd sample-app
go test
```

Run it to make sure it works.

```
$ chmod +x run-unit-tests.sh
$ ./run-unit-tests.sh
PASS
ok  	_/Users/taylor/workspace/concourse-tutorial/sample-app	0.006s
```

Now we can start writing the configuration for the task. A tasks configuration can live in it's own YAML file or it can be embedded in a [pipeline's](https://concourse-ci.org/pipelines.html) configuration. We'll follow the common practice of placing the task config in its own file.

Within a task config you can either: point to a script or binary to run, or use this neat trick to embed your script within your task's config! Our first iteration will embed the above script directly in the task's config.

```yaml
---
platform: linux

image_resource:
  type: docker-image
  source:
    repository: golang
    tag: '1.12.7'

inputs:
- name: concourse-tutorial

run:
  path: sh
  args: 
    - -c
    - |
      #!/bin/bash
      
      cd sample-app
      go test
```

Let's go through what this task configuration is telling Concourse.

```yaml
platform: linux
```

This line tells Concourse which type of Concourse Worker to run this task on. There are three types of workers: Linux, Windows, and Darwin (macOS). 

Let's look at the next part of the task config.

```yaml
image_resource:
  type: registry-image
  source:
    repository: golang
    tag: '1.12.7'
```

This part tells Concourse which container image to run your task in. If you have existing docker images that are already loaded with a bunch of your tooling this is where you can leverage them in Concourse. 

The `image_resource` parameter relies on a [Concourse resource](https://concourse-ci.org/resources.html) that returns the [`rootfs`](https://github.com/concourse/registry-image-resource#rootfs) of a container image. Concourse comes with a default resource called [`registry-image`](https://github.com/concourse/registry-image-resource). By default, the `registry-image` resource pulls containers from [Docker Hub](https://hub.docker.com/). In summary, this part of the task config is telling Concourse to use the [Go docker image tagged 1.12.7](https://hub.docker.com/_/golang). By using the Go docker image, we can run our script inside a container where we know Go is installed.

Let's look at the next section.

```yaml
inputs:
- name: concourse-tutorial
```

With this line, we're simply stating that our tasks expects an input called `concourse-tutorial`. Inputs for tasks are mounted as volumes inside a task's container image at the working directory. Each input is placed in a folder with the same name as the input. You can change the name by using the `path` parameter, like this:

```yaml
inputs:
- name: concourse-tutorial
  path: ct
```

To keep things simple, we're going to exclude the `path` parameter. 

With this input in our task, if we were to `ls` inside the container that our task script will be running in, we would see the following:

```
$ ls
concourse-tutorial
```

We see a single directory called `concourse-tutorial`. If our task had more inputs we would see more directories, each matching the name of their respective `input` name.

Time for the last section. The `run` section tells Concourse what script/binary to execute inside the container. You can also pass in as many arguments to your executable as you want. 

```yaml
run:
  path: sh
  args: 
    - -c
    - |
      #!/bin/bash
      
      cd sample-app
      go test
```

In our `run` section we're telling Concourse to execute `sh` and pass in two arguments: 

1. `-c` which tells shell to read and execute commands from the next string
2. our "string" of commands, which is our original bash script.

This is how we can embed our script inside our tasks config. Later on we'll extract the script from our task config. For now this works just fine. 

