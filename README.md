# NAME

html2words

# SYNOPSIS

	html2words URL

# DESCRIPTION

Extract text from a URL. Ignores text from the following nodes:
* noscript
* script
* style

# INSTALLATION
	curl -L https://github.com/mark-5/go-html2words/releases/download/v0.0.1/html2words-`uname -s`-`uname -m` -o /usr/local/bin/html2words
	chmod +x /usr/local/bin/html2words

