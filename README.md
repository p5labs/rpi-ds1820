# rpi-ds1820
=======

Connecting DS18B20 temperature sensor to Raspberry Pi using GPIO bit-banging method.

Warning: Incorrect wiring can damage your Raspberry Pi or sensor. See connection diagrams online. Specifying wrong GPIO PIN for bit-banging could cause damage if that PIN has connected something else (not DS18B20 data pin). Use at your own risk.

ds1820util is writen in C and uses memory mapped file `/dev/gpiomem` for GPIO manipulation - similar to other tools and libraries written in C. Check that file `/dev/gpiomem` is present in your RPi Linux distribution (`ls /dev` or `ls /dev | grep/gpiomem`).


### Usage

- Download binary `ds1820util`
- `chmod +x ds1820util`
- Run `./ds1820util --pin=4` (specify correct GPIO number!)

Running without parameters produces help message:

```
Usage: ds1820util --pin=4 [--verbose] [--command=read]
  --pin=[number]   logical GPIO PIN number
  --cmd=[cmd]      command to execute. 'convert' - start temperature measurement/conversion. read - read temperature (does not measure, reads last value).
  --verbose        verbose output for more details
```

Running without `--verbose` produces minimal output - temperature in Celsius. This is very convenient if you want to capture output by calling it from shell script or higher level language (golang, python, php).
In default mode ds1820util commands sensor to measure (convert) temperature, waits for it to complete (750ms) and then reads data. This can be split to separate steps by specifying `--cmd=convert` or `--cmd=read`. This allows to optimize time by avoiding to wait 750ms. E.g several sensors on different pins can be commanded to convert temperature value, and after 750ms all of them can be polled one by one. Communication takes just 2-3ms compared to measurement/conversion time of 750ms.
After powering up, DS18B20 loads default value of 85C (degrees celsius). You would get this when sensor has never completed measurement after powering up (or restart).

### Example output

```
 $ ./ds1820util --pin=4 --verbose
GPIO PIN 4
Scratchpad: 52014b467fff0e10ff
Temperature: 21.12C
Resolution 12 bits
Completed in 0.771s
```

### Troubleshooting

Communications are not always very stable. Bit-banging method requires precise timing, what is not always possible from user mode in operating system. Communication reliability depends on wiring, processor load, sensor quality, etc.. 
Error output will provide some clues.

```
Error: no presence signal during init procedure.
```
Communication starts by simple initialization procedure. There was no response signal from device. Possibly wrong wiring.

```
Error: CRC mismatch.
```
Received and calculated scratchpad checksum does not match. Could be caused by wiring or timing issues. Also check Scratchpad byte output in --verbose output.