<?php
/*
Plugin Name: FlashCache
Plugin URI: http://www.lateral.co.za/FlashCache
Description: FlashCache provides a 'flash' caching mechanism for transferring information snippets between browser requests.
Version: 1.0.0
Author: Lateral Alternative CC
Author URI: http://www.lateral.co.za
License: AllRightsReserved to Lateral Alternative CC
*/
include_once("class.FlashCache.php");

add_action('wp_enqueue_scripts', function() {
	wp_register_style('flashcache', plugins_url('flashcache-plugin/css/flashcache.css'),array(), '1.0.0');
	wp_enqueue_style('flashcache');
});
