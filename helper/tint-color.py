#!/usr/bin/env python3

# pylama:ignore=E501

import os
import sys

try:
    import colorclass
except ImportError:
    print("tint-color requires colorclass to be installed (you can use pip)", file=sys.stderr)
    sys.exit(1)

# Handle Windows vs everyone else
if os.name == "nt":
    colorclass.Windows.enable(auto_colors=True)
else:
    colorclass.set_dark_background()


def line_process(line):
    """Add any necessary color to the line and return the final, corrected line."""
    # res/small_xist_checks.txt:3:14:warning: 'effete' has issues - rephrase.
    flds = line.strip().split(':')
    if len(flds) != 5:
        return line  # Not in a format we recognize

    clr = colorclass.Color

    # Note that we currently do nothing to msg
    fname, line, col, warn, msg = flds
    fname = clr('{autocyan}%s{/autocyan}' % fname)
    line = clr('{autogreen}%s{/autogreen}' % line)
    col = clr('{autogreen}%s{/autogreen}' % col)
    warn = clr('{autoyellow}%s{/autoyellow}' % warn)

    delim = clr('{b}:{/b}')

    return delim.join([fname, line, col, warn, msg])


def main():
    for line in sys.stdin:
        sys.stdout.write(line_process(line))
        sys.stdout.write('\n')


if __name__ == "__main__":
    main()
