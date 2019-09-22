# Shard API

Simple Proof-of-concept to demo sharding as a service

## Getting Started

```bash
go get -u -v .
go build
./shard-api
2019/09/22 11:43:44.082710 main.go:33: Starting shard-api on port 4000...
```

That's it - server's running. You can specify a different port with the `-port` flag.

## Examples

### Create some shards

```bash
$ curl localhost:4000/create -vv -d '{"minimum":3, "shares":10, "secret":"look at my shards"}'
...
{"shards":["M0ozAhU3-Psd3E-eASAJYS1xPIkjgUMyJIf8ofGduJs=WPP_E_sQxyYtDBwqngBSf9YHpv1nJpJ7TDXYm_Kgkkc=","VDn9vOkg054jtBhdUyyK0wweuVetyXfI-nos30KNemQ=lgTVSwe_HUVmQjmOEadflpiY_tPMiTDir5b_m1Ug2F0=","qQ06OniPDEEOdLuUoW9A6RYsQNOYGWgdxsxL-ZDHfGg=tM6ULvKwmGyxrMMh1Pz8G7-iFOsWME7-ApRqeJC9Xm8=","_7eUhIrFIc7MJ70iQRme_VkDlx4wHgEbjFjTFJfyBk4=CVeXiGDX4FmoO_hrWQvhCfpqNYWwgOzzT_lD3V0iDPs=","HV6L1NiuyAcrbfXL6joqcxLZAaVRC3MT-m653XKhDZM=WTPcXGJ3tRytn4ZUa7yV6iobgEF7oi7ojzhcbNh2DdE=","mzZ80tkHz4m0CqnjRgbpjTXNPmhiwzPqb-vusB9B5gs=5fibxMbZ6mKNCGHUgb6s3aV3dm-lxRMHHB9nC2kvyT8=","oyI0Ldjyl8fI1Ly0P7WnymgWGDULpJKjbTCVU-CgGcE=NoqCdnL3cciFebTyl9rRHnHb28MBIRcHumN4H4RD9-A=","p-nJ2v2wFKq7M9US6tqHeoy6kKeeqU8gnuJSr4d8FBg=-dDP9cR5c3N3RAPUEox7dsvxYhpe-bloSWP_GYtunJI=","L_FDCST5yAOkM1mYXkqjynLznKzXEITYMutDz7Prqeg=kqc8WQIv4MQb87mp4mdF0KJI9P1NlM6Ne_fKVkTApdw=","v5CvUXx4fKi99DWP2845dSJyLzqhhnXlgmVQJQFieJk=TPSGxm05AjV-8im8BFtOvwWoi9s-c1izQJBITUbdkvE="]}
```

### Combine some shards

As above, we require a minimum of 3 shards to make it work up to a a maximum of 10

```bash
$ curl localhost:4000/combine -vv -d '{"shards":["M0ozAhU3-Psd3E-eASAJYS1xPIkjgUMyJIf8ofGduJs=WPP_E_sQxyYtDBwqngBSf9YHpv1nJpJ7TDXYm_Kgkkc=","VDn9vOkg054jtBhdUyyK0wweuVetyXfI-nos30KNemQ=lgTVSwe_HUVmQjmOEadflpiY_tPMiTDir5b_m1Ug2F0=","p-nJ2v2wFKq7M9US6tqHeoy6kKeeqU8gnuJSr4d8FBg=-dDP9cR5c3N3RAPUEox7dsvxYhpe-bloSWP_GYtunJI="]}'
...
{"secret":"look at my shards"}
```
