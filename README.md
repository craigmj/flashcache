# flashcache
Flash Cache is an in-memory, short-term JSON cache for transferring 'flash messages' between web pages.

JSON values are set for a key, with one or more HTTP POSTs. When the values are retrieved with an HTTP GET, and array of all set JSON values is returned, and all the set values are cleared in the cache.

Values in the cache expire after 1 minute.

FlashCache was designed as an easy solution for sending 'status' messages to a website user. When tasks succeed, the message is saved to FlashCache. When each page is updated (or on an AJAX timed request), all status updates are retrieved and displayed to the user.

# Usage
FlashCache runs as a Go programme that servers the cache on http://localhost:16021 by default.

To set a value for a given key, POST to /[key]

Keys are held for 1 minute.

To retrieve the set values for a key, GET from /[key]

# Installation and Running (Ubuntu)
Install by cloning the repo, then running `./build.sh` in the root directory. This will install Go dependencies and build flashcache to `bin/flashcache`

To install an upstart job to automatically start flashcache with the system,

	bin/flashcache upstart | sudo tee /etc/init/flashcache.conf

# Environment Values

FlashCache will read an environment variable called `FLASHCACHE_SERVER`, which should be of the form http://localhost:12345/`, and will run on the port defined in that environment variable. You can see a sample script by running
	
	bin/flashcache setenv

You can also get FLASHCACHE to add this to `/etc/environment` with

	sudo bin/flashcache write-env

# PHP & Wordpress

FlashCache comes with a basic PHP class in `flashcache-plugin/class.FlashCache.php`. This will use the `FLASHCACHE_SERVER` environment variable to locate the FlashCache server, or will default to FlashCache's default. Use

	FlashCache::Set($key, $data);
	for (FlashCache::Get($key) as $data) { .. }

to set and retrieve flash cached data. There is also a convenience FlashCache::PutMessage(..) and FlashCache::WriteMessages() method that is easily examined and self-explanatory.

The `flashcache-plugin` directory also functions as a Wordpress plugin. Just drop it into your Wordpress `wp-content/plugins` directory, and activate the FlashCache plugin. Then you will have access to the FlashCache class in your Wordpress code, and it will automatically add some very basic FlashCache css to present messages written with the convenience ::WriteMessages() method.


# Other bits
A GET to /__version returns the current version of FlashCache.
A GET to /__dump presents a text dump of all the data in the cache. This is purely for debugging purposes, and does not clear the cache.
