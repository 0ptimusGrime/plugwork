# plugwork (working title)

The general flow of this is:

Something satisfying the `Readable` interface accepts user input and emits `message.Instruction` structs onto a read queue.

A go func picks messages off of this read queue, and then dispatches it to any devices registred in the `device.Set`. Each of these devices will then transform the generic `message.Instruction` struct into a device-specific implementation -- since I don't imagine that all devices will have compatible implementations of a given abstract feature.

This transformation is done by mapping constants defined in the `capabilities` package to a specific transform function. The result of this transform function is what's ultimately emitted to the end device.

usage:

in one terminal, run `nc -l -u 127.0.0.1:9999`

in another, simply:

```
$ go run main.go
Starting console reader. ctrl+c to exit
> vibrate
> sending message [vibrate] to 1 devices...
sending messsage vibrate to device UDPVibrator
```

You should the UDP message emitted in the terminal running netcat