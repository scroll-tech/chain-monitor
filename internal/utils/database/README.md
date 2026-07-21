# database

Connection helpers for the chain-monitor Postgres database (GORM + pgx).

## AWS RDS IAM authentication

`InitDB` can authenticate to AWS RDS/Aurora PostgreSQL using short-lived IAM
auth tokens instead of a DSN password. This removes the need to rotate database
passwords: access is granted via IAM and tokens are regenerated automatically
(they expire every 15 minutes) for each new connection.

Enable it in the `db_config` block:

```json
{
  "driver_name": "postgres",
  "dsn": "postgres://svc_user@mydb.abc123.us-east-1.rds.amazonaws.com:5432/scroll?sslmode=require",
  "maxOpenNum": 200,
  "maxIdleNum": 20,
  "useIAMAuth": true,
  "awsRegion": "us-east-1"
}
```

RDS IAM tokens are signed per `(host, port, db-user)`: the DSN's DB user must be
granted `rds_iam`, and the workload's IAM role must have `rds-db:connect` on the
corresponding `dbuser` resource.

Notes:

- Omit the password from the `dsn`; it is supplied by the generated IAM token.
- `awsRegion` is optional — when empty it is resolved from the default AWS
  config chain (e.g. the `AWS_REGION` environment variable). AWS credentials come
  from the default chain too (env vars, EKS IRSA / EC2 instance role), which
  self-refreshes.
- IAM auth requires TLS, so the DSN must set `sslmode=require` or higher.
  `sslmode=disable`/`allow`/`prefer` (including an unset `sslmode`, which
  defaults to `prefer`) are rejected at startup rather than silently sending the
  token over a connection that may fall back to plaintext. `require` encrypts
  but does not verify the server certificate; for that, use `sslmode=verify-full`
  with `sslrootcert` pointing at the
  [RDS CA bundle](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html).
- Leaving `useIAMAuth` unset preserves the previous password-based behavior.

### Metrics

When IAM auth is active the following Prometheus metrics are exported:

- `database_rds_iam_token_total` — tokens generated (one per new connection).
- `database_rds_iam_token_failure_total` — token generation failures.
- `database_rds_iam_token_duration_seconds` — token generation latency.
