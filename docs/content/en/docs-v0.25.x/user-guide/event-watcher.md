---
title: "Event watcher"
linkTitle: "Event watcher"
weight: 14
description: >
  A helper facility to automatically update files when it finds out a new event.
---

![](/images/diff-by-eventwatcher.png)

The only way to upgrade your application with PipeCD is modifying configuration files managed by the Git repositories.
It brings benefits quite a bit, but it can be painful to manually update them every time in some cases (e.g. continuous deployment to your development environment for debugging, the latest prerelease to the staging environment).

If you're experiencing any of the above pains, Event watcher is for you.
Event watcher works as a helper facility to seamlessly link CI and CD. This feature lets you automatically update files managed by your Piped when an arbitrary event has occurred.
While it empowers you to build pretty versatile workflows, the canonical use case is that you trigger a new deployment by image updates, package releases, etc.

This guide walks you through configuring Event watcher and how to push an Event.

## Prerequisites
Before we get into configuring EventWatcher, be sure to configure Piped. See [here](/docs/operator-manual/piped/configuring-event-watcher/) for more details.

## Usage
File updating can be done by registering the latest value corresponding to the Event in the control-plane and comparing it with the current value.

Therefore, you mainly need to:
1. define which values in which files should be updated when a new Event found.
1. integrate a step to push an Event to the control-plane using `pipectl` into your CI workflow.

### 1. Defining Events
Prepare EventWatcher configuration files under the `.pipe/` directory at the root of your Git repository.
In that files, you define which values in which files should be updated when the Piped found out a new Event.

For instance, suppose you want to update the Kubernetes manifest defined in `helloworld/deployment.yaml` when an Event with the name `helloworld-image-update` occurs:

```yaml
apiVersion: pipecd.dev/v1beta1
kind: EventWatcher
spec:
  events:
    - name: helloworld-image-update
      replacements:
        - file: helloworld/deployment.yaml
          yamlField: $.spec.template.spec.containers[0].image
```

The full list of configurable `EventWatcher` fields are [here](/docs/user-guide/configuration-reference/#event-watcher-configuration).

### 2. Pushing an Event with `pipectl`
To register a new value corresponding to Event such as the above in the control-plane, you need to perform `pipectl`.
And we highly recommend integrating a step for that into your CI workflow.

You first need to set-up the `pipectl`:

- Install it on your CI system or where you want to run according to [this guide](/docs/user-guide/command-line-tool/#installation).
- Grab the API key to which the `READ_WRITE` role is attached according to [this guide](/docs/user-guide/command-line-tool/#authentication).

Once you're all set up, pushing a new Event to the control-plane by the following command:

```bash
pipectl event register \
    --address={CONTROL_PLANE_API_ADDRESS} \
    --api-key={API_KEY} \
    --name=helloworld-image-update \
    --data=gcr.io/pipecd/helloworld:v0.2.0
```

You can see the status on the event list page.

![](/images/event-list-page.png)


After a while, Piped will create a commit as shown below:

```diff
     spec:
       containers:
       - name: helloworld
-        image: gcr.io/pipecd/helloworld:v0.1.0
+        image: gcr.io/pipecd/helloworld:v0.2.0
```

NOTE: Keep in mind that it may take a little while because Piped periodically fetches the new events from the control-plane. You can change its interval according to [here](/docs/operator-manual/piped/configuring-event-watcher/#optional-settings-for-watcher).

### [optional] Using labels
Event watcher is a project-wide feature, hence an event name is unique inside a project. That is, you can update multiple repositories at the same time if you use the same event name for different events.

On the contrary, if you want to explicitly distinguish those, we recommend using labels. You can make an event definition unique by using any number of labels with arbitrary keys and values.
Suppose you define an event with the labels `env: dev` and `appName: helloworld`:

```yaml
apiVersion: pipecd.dev/v1beta1
kind: EventWatcher
spec:
  events:
    - name: image-update
      labels:
        env: dev
        appName: helloworld
      replacements:
        - file: helloworld/deployment.yaml
          yamlField: $.spec.template.spec.containers[0].image
```

The file update will be executed only when the labels are explicitly specified with the `--labels` flag.

```bash
pipectl event register \
    --address=CONTROL_PLANE_API_ADDRESS \
    --api-key=API_KEY \
    --name=image-update \
    --labels env=dev,appName=helloworld \
    --data=gcr.io/pipecd/helloworld:v0.2.0
```

Note that it is considered a match only when labels are an exact match.

## Examples
Suppose you want to update your configuration file after releasing a new Helm chart.

You define `.pipe/event-watcher.yaml` like:

```yaml
apiVersion: pipecd.dev/v1beta1
kind: EventWatcher
spec:
  events:
    - name: helm-release
      labels:
        env: dev
        appName: helloworld
      replacements:
        - file: helloworld/app.pipecd.yaml
          yamlField: $.spec.input.helmChart.version
```

Push a new version `0.2.0` as data when the Helm release is completed.

```bash
pipectl event register \
    --address=CONTROL_PLANE_API_ADDRESS \
    --api-key=API_KEY \
    --name=helm-release \
    --labels env=dev,appName=helloworld \
    --data=0.2.0
```

Then you'll see that Piped updates as:

```diff
apiVersion: pipecd.dev/v1beta1
kind: KubernetesApp
spec:
  input:
    helmChart:
      name: helloworld
-     version: 0.1.0
+     version: 0.2.0
```

See [here](https://github.com/pipe-cd/examples/blob/master/.pipe) for more examples.

## Github Actions
If you're using Github Actions in your CI workflow, [actions-event-register](https://github.com/marketplace/actions/pipecd-register-event) is for you!
With it, you can easily register events without any installation.
