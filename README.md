# logrus for Golang DataDog Logs

This is an example of using the [logrus](https://pkg.go.dev/github.com/sirupsen/logrus) golang library to ship logs directly to Datadog using the API, no agent install required.

logrus has the concept of [Hooks](https://pkg.go.dev/github.com/sirupsen/logrus#readme-hooks), which you can write your own plugins to write plugins to submit logs to another source.

We're using the [GlobalFreightSolutions/logrus-datadog-hook](https://github.com/GlobalFreightSolutions/logrus-datadog-hook) hook, as this seems the cleanest in terms of documentation and example code.

# Setup

Setup your API key with `DD_API_KEY`:

```
export DD_API_KEY=<Your API key>
```

If you're using DataDog EU, export your `DD_SITE`:

```
export DD_SITE=datadoghq.eu
```

Then, run the code:

```
$ go run main.go
INFO[0000] This is a log sent to datadog - 1
INFO[0001] This is a log sent to datadog - 2
INFO[0002] This is a log sent to datadog - 3
INFO[0003] This is a log sent to datadog - 4
```

Then jump into your log explorer in Datadog, and you should see the logs under: 

![Log Screenshot](https://user-images.githubusercontent.com/1064715/134496592-98e8c784-05cd-4b69-a5db-15e2fb928f73.png)