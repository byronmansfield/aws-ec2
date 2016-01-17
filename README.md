# AWS CLI Tool

This project is mainly used for my own personal practice, but hopefully will grow to something I can use in my personal practice or at work. It is intended to be an AWS Command Line Tool to help me get lists of instances with some details about each. It also offers some filtering such as by region or environment.

## Example Useage

You can build the project once you cloned it like so

```
go build
```

Then once you have the executable you can use it just by itself

```
./aws-tools
```
which will give you a full list of the default region `us-west-1` plus all `environments`.

Or you can pass in `region` or `environment`

```
./aws-tools -region=eu-west-1
```

```
./aws-tools -environment=staging
```

```
./aws-tools -region=us-east-1 -environment=dev
```

## Other notes

It is a work in progress and already needs some refactoring in order to scale. I hope to add some error handling to it, as well as more filtering, possibly even colors to make it easier to read. 
