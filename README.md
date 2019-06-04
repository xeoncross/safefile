# Safefile

I wanted a way to carefully save io.Reader streams to files ensuring the contents were flushed to disk and all errors (even Close()) were properly handled.
