<?php

function so_json_serialize( $value ) {
	if ( empty( $value ) ) {
		return;
	}

	$out = array();

	if ( is_array( $value ) ) {
		foreach( $value as $index => $element ) {
			if ( method_exists( $element, 'jsonSerialize' ) ) {
				array_push( $out, $element->jsonSerialize() );
			}
		}

		unset($out['@context']);
		return $out;
	}

	if ( is_object( $value ) ) {
		if ( method_exists( $value, 'jsonSerialize' ) ) {
			$out = $value->jsonSerialize();
		}

		unset($out['@context']);
		return $out;
	}

	$out = $value;
	return $out;
}
