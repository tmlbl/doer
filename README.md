doer
====

Doer is a super-simple service for triggering remote actions. It accepts a JSON
configuration file that looks like this:

```json
{
  "tasks": [
    {
      "name": "test-task",
      "commands": [
        "touch main.go"
      ],
      "secret": "abc123"
    }
  ]
}
```

It can be started using with TLS enabled like this:

```bash
doer -cert mycert.pem -key privkey.pem
```

You can then trigger your remote action like this:

```bash
curl https://myserver:8778/test-task?secret=abc123
```
