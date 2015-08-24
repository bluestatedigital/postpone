# postpone

A framework for distributed job execution.  A hybrid of [Resque](https://github.com/resque/resque), [Sidekiq](http://sidekiq.org), and `cron(8)`, based on [go-workers](https://github.com/jrallison/go-workers).

## scheduler

Enqueues tasks periodically based on a given configuration file (see [sample](scheduler/sample_schedule.yaml)).

    postpone-scheduler --config=schedule.json

But, really, you should wrap that in `consul lock` and run multiple instances.

    consul lock apps/postpone/scheduler postpone-scheduler --config=schedule.json

Of course, the process should be supervised, too.  Furthermore, you probably want to put the schedule config into Consul and render it with `consul-template`.

Both JSON and YAML syntaxes are supported.

## shell-worker

Run up to 3 concurrent jobs submitted to the `shell.high` queue:

    postpone-shell-worker --queue=shell.high --concurrency=3

Obviously the commands specified in the enqueued jobs must be available to the shell-worker.

## common options

      --debug       enable debug [$DEBUG]
      --log-file=   path to JSON log file [$LOG_FILE]
      --redis-host= redis hostname (localhost) [$REDIS_HOST]
      --redis-port= redis port (6379) [$REDIS_PORT]
      --redis-db=   redis database (0) [$REDIS_DB]
