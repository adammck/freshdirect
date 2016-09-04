# Fresh Direct

[![GoDoc](https://godoc.org/github.com/adammck/freshdirect?status.svg)](https://godoc.org/github.com/adammck/freshdirect)
[![Build Status](https://travis-ci.org/adammck/freshdirect.svg?branch=master)](https://travis-ci.org/adammck/freshdirect)

This is a Go interface to [FreshDirect](https://freshdirect.com). It's roughly
zero percent complete, but works well enough for me to add groceries to my cart
from the command line, which is a totally normal thing to want to do.


## Usage

    go get https://github.com/adammck/freshdirect/examples/fd

    # Grab this with the web inspector for now
    export FRESHDIRECT_SESSION_ID=xxxx

    # Add items by SKU. Run the following on any product page to grab the SKU:
    #   document.querySelector("form[data-component=product] input[name=skuCode]").value
    fd MEA3335058 BAK4651005 VEG0030538


## License

MIT
