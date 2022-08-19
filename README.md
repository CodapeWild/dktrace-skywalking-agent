# Skywalking Agent for Datakit

**Notice:** THIS PROJECT IS STILL IN PROGRESS

This tool used to send standard Skywalking tracing data to Datakit.

The features include:

- build with [go2sky Skywalking Golang lib](https://github.com/SkyAPM/go2sky)
- customized Span data
- configurable multi-thread pressure test

## Config

**Config structure in `config.json`**

```json
{
  "dk_agent": "127.0.0.1:9529",
  "sender": {
    "threads": 1,
    "send_count": 1
  },
  "service": "dktrace-skywalking-agent",
  "dump_size": 1024,
  "random_dump": true,
  "trace": []
}
```

- `dk_agent`: Datakit host address
- `sender.threads`: how many threads will start to send `trace` simultaneously
- `sender.send_count`: how many times `trace` will be send in one `thread`
- `service`: service name
- `dump_size`: the data size in kb used to fillup the trace, 0: no extra data
- `random_dump`: whether to fillup the span with random size extra data
- `trace`: represents a Trace consists of Spans

## Span Structure

**Span structure in `config.json`**

`trace`\[span...\]

```json
{
  "resource": "/get/user/name",
  "operation": "user.getUserName",
  "span_type": "",
  "duration": 1000,
  "error": "access deny, status code 100010",
  "tags": [
    {
      "key": "db.type",
      "value": "sqlite"
    }
  ],
  "children": []
}
```

**Note:** Spans list in `trace` or `children` will generate concurrency Spans, nested in `trace` or `children` will show up as calling chain.

- `resource`: resource name
- `operation`: operation name
- `span_type`: Span type [app cache custom db web]
- `duration`: how long an operation process will last
- `error`: error string
- `tags`: Span meta data, imitate client tags
- `children`: child Spans represent a subsequent function calling from current `operation`
