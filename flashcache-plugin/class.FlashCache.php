<?php
/**
 * FlashCache works with the Go FlashCache server to provide
 * flash message facilities to PHP servers.
 */
class FlashCache {
	/**
	 * Server returns the Server location for the
	 * FlashCache server.
	 */
	public static function Server($path="") {
		$s = getenv("FLASHCACHE_SERVER");
		if (!$s) {
			$s = "http://localhost:16021/";
		}
		return $s.$path;
	}
	/** 
	 * Version returns the version of the FlashCache server.
	 */
    public static function Version() {
    	$v = self::Get("__version");
    	if (!$v) {
    		return " - offline - ";
    	}
    	return $v->version;
    }
	/** 
	 * Set sets the flash value for the given key to the 
	 * received data.
	 */
	public static function Set($key, $data) {
		return self::doRequest($key, $data);
	}
	/** 
	 * Get retrieves any falsh values for the given key.
	 */
	public static function Get($key) {
		return self::doRequest($key, FALSE);
	}
	/**
	 * PutMessage puts a message with the given title and message text (html) and optional css class.
	 * WriteMessages() will write the messages to the browser. This is a convenience
	 * method to work with PHP / Wordpress / etc.
	 */
	public static function PutMessage($title, $msg, $css="", $group="-flashcache-messages") {
		self::Set(session_id() . $group, array("css"=>$css, "title"=>$title, "msg"=>$msg));
	}
	/**
	 * WriteMessages writes any messages saved with PutMessage()
	 */
	public static function WriteMessages($group="-flashcache-messages") {
		$all = self::Get(session_id().$group);
		if (!$all) {
			return;
		}
		if (0<count($all)) {
			echo '<div class="flashcache-messages">';
			foreach ($all as $a) {
				echo <<<EOHTML
<div class="flashcache-message $a->css">
	<div class="flashcache-title">$a->title</div>
	<div class="flashcache-content">$a->msg</div>
</div>
EOHTML;
			}
			echo '</div>';
		}
	}
	// If $data is set, this is a SET request, and the function
	// returns TRUE on success, FALSE otherwise.
	// If $data is NOT set, the function returns the array received
	// from the server, or FALSE if the request fails.
	protected static function doRequest($key, $data=FALSE) {
		$c = curl_init(self::Server($key));
		if ($data) {
			curl_setopt($c, CURLOPT_POST, 1);
			curl_setopt($c, CURLOPT_POSTFIELDS, json_encode($data));
		}
		curl_setopt($c, CURLOPT_RETURNTRANSFER, TRUE);
		$res = curl_exec($c);
		if (FALSE===$res) {
			$err = curl_error($c);
			error_log("ERROR in FLASHCACHE::doRequest: " . $err);
			curl_close($c);
			return FALSE;
		}
		curl_close($c);
		if ($data) {
			return TRUE;
		}
		$json = json_decode($res);
		$err = false;
		if (property_exists($json,"error") && $json->error) {
			error_log("ERROR JSON Decoding FLASHCACHE RESPONSE ($res): " . $json->error);
			return FALSE;
		}
		return $json;
	}
}