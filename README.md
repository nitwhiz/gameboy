# gameboy

A(nother) game boy emulator written in Go.

This is meant to be used as a library, not a batteries-included emulator to play games. It's not focused on accuracy, but on ease to emulate frame-by-frame.

# Features

- Passes Blargg `cpu_instrs` Test ROMs
- LCD/GFX is supported
- Sound is **not** supported
- Game Boy Color is **not** supported

# Credits

- Heavily inspired by [goboy](https://github.com/Humpheh/goboy)
- Rough implementation ideas from [codeslinger.co.uk](http://www.codeslinger.co.uk/pages/projects/gameboy/hardware.html)
- Impossible without
  - [Pan Docs](https://gbdev.io/pandocs/)
  - [Opcode Reference 1](https://gbdev.io/gb-opcodes//optables/)
  - [Opcode Reference 2](https://meganesu.github.io/generate-gb-opcodes/)
  - [Game Boy Doctor](https://github.com/robert/gameboy-doctor)

